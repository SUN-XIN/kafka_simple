* start zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
* start kafka
bin/kafka-server-start.sh config/server.properties
* consumer of topic test
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
* producer of topic test
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test