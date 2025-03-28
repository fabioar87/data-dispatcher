# Data Dispatcher Service

The Data Dispatcher is a Golang service responsible for:
* serialize the Data Source CSV file containing the water level readings
* send the time series reading to the message bus service

The Dockerfile and docker-compose is used in the development scope to test its interaction with
the project message bus service.

Create a Makefile to automate the commands.

### Kafka Commands - Used during the design process
Reference of some important kafka commands.********

```shell
# Creating a new kafka topic
$ docker-compose exec kafka kafka-topics --create --topic test-topic --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092

# List all topics
$ docker-compose exec kafka kafka-topics --list --bootstrap-server localhost:9092

# Publish messages in the topic
$ docker-compose exec kafka kafka-console-producer --topic test-topic --bootstrap-server localhost:9092

# Consuming messages from the topic
$ docker-compose exec kafka kafka-console-consumer --topic test-topic --from-beginning --bootstrap-server localhost:9092
```