# rest-kafka-bridge-example

This is an example of streaming messages from JSON over HTTP to Protobuf over Kafka using [Watermill](https://github.com/ThreeDotsLabs/watermill).

To run it execute `dcoker-compose up rest-kafka-bridge-example`. It will start application under `localhost:8080`.
Post following body to `localhost:8080/readings`.
```
{
   "timestamp":"2021-02-14T15:46:47.100739+01:00",
   "machine_id":"ac089c14-a033-4ade-a690-2627328a7f67",
   "temperature":10.5
}
```
