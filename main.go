package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/fabioar87/data-dispatcher/proto"
)

const (
	csvFilePath = "26027_athlone_s_complete.csv"
)

func main() {
	brokers := []string{"localhost:9092", "kafka:9092", "0.0.0.0:9092"}
	topic := "test-topic"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating Kafka producer: ", err)
	}
	defer producer.Close()

	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1

	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error reading header", err)
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if len(record) > 0 && strings.HasPrefix(record[0], "#") {
			continue
		}

		if err != nil {
			log.Println("bad record:", err)
			continue
		}

		dayMin, err := ParseFloat(record[1])
		if err != nil {
			continue
		}

		dayMean, err := ParseFloat(record[2])
		if err != nil {
			continue
		}

		dayMax, err := ParseFloat(record[3])
		if err != nil {
			continue
		}

		waterLevel := &pb.WaterLevel{
			Date:    record[0],
			DayMin:  dayMin,
			DayMean: dayMean,
			DayMax:  dayMax,
		}

		payload, err := proto.Marshal(waterLevel)
		if err != nil {
			log.Printf("Error serialising proto: %v", err)
			continue
		}

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(payload),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Error sending message to kafka: %v", err)
		} else {
			log.Printf("Message sent to partition %d at offset %d", partition, offset)
		}

		// Simulating the sensor delay
		time.Sleep(500 * time.Millisecond)
	}
}

func ParseFloat(data string) (float32, error) {
	data_ := strings.TrimSpace(data)
	if data_ == "" {
		return 0, errors.New("invalid data")
	}

	value, err := strconv.ParseFloat(data_, 32)
	if err != nil {
		return 0, fmt.Errorf("error parsing value %v", err)
	}

	return float32(value), nil
}
