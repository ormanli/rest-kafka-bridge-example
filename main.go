package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	stdHttp "net/http"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/goccy/go-json"
	"github.com/golang/protobuf/proto"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerHost string `split_words:"true"`
	ServerPort int    `required:"true" split_words:"true"`
	KafkaHost  string `required:"true" split_words:"true"`
	KafkaPort  int    `required:"true" split_words:"true"`
}

func main() {
	logger := watermill.NewStdLogger(true, true)

	var c Config
	err := envconfig.Process("app", &c)
	if err != nil {
		logger.Error("can't initialise config", err, nil)
		os.Exit(1)
	}

	kafkaPublisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{fmt.Sprintf("%s:%d", c.KafkaHost, c.KafkaPort)},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		logger.Error("can't initialise kafka publisher", err, nil)
		os.Exit(1)
	}

	httpSubscriber, err := http.NewSubscriber(
		fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort),
		http.SubscriberConfig{
			UnmarshalMessageFunc: func(topic string, request *stdHttp.Request) (*message.Message, error) {
				b, err := ioutil.ReadAll(request.Body)
				if err != nil {
					return nil, fmt.Errorf("can't read body: %w", err)
				}

				return message.NewMessage(watermill.NewUUID(), b), nil
			},
		},
		logger,
	)
	if err != nil {
		logger.Error("can't initialise http subscriber", err, nil)
		os.Exit(1)
	}

	r, err := message.NewRouter(
		message.RouterConfig{},
		logger,
	)
	if err != nil {
		logger.Error("can't initialise router", err, nil)
		os.Exit(1)
	}

	r.AddMiddleware(
		middleware.Recoverer,
		middleware.CorrelationID,
	)
	r.AddPlugin(plugin.SignalsHandler)
	r.AddHandler(
		"rest_to_kafka",
		"/readings",
		httpSubscriber,
		"readings",
		kafkaPublisher,
		func(msg *message.Message) ([]*message.Message, error) {
			var tr TemperatureReading

			if err := json.Unmarshal(msg.Payload, &tr); err != nil {
				return nil, fmt.Errorf("can't unmarshal json message: %w", err)
			}

			if tr.MachineId == "" {
				return nil, errors.New("empty machine id")
			}

			if !tr.Timestamp.IsValid() {
				return nil, errors.New("empty timestamp")
			}

			b, err := proto.Marshal(&tr)
			if err != nil {
				return nil, fmt.Errorf("can't marshal protobuf message: %w", err)
			}

			return []*message.Message{message.NewMessage(msg.UUID, b)}, nil
		},
	)

	go func() {
		<-r.Running()
		_ = httpSubscriber.StartHTTPServer()
	}()

	_ = r.Run(context.Background())
}
