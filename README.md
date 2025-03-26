# Data Dispatcher Service

The Data Dispatcher is a Golang service responsible for:
* serialize the Data Source CSV file containing the water level readings
* send the time series reading to the message bus service

The Dockerfile and docker-compose is used in the development scope to test its interaction with
the project message bus service.