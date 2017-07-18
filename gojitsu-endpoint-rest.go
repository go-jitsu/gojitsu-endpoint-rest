package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/Shopify/sarama"
	"log"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	tenant := tenant(&request.Header)

	path := strings.Split(request.URL.Path, "/")

	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)

	if (path[1] == "data") {
		key := request.Header.Get("KEY")
		streamName := path[2]
		event := DataEvent{streamName, key, body}

		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				log.Fatalln(err)
			}
		}()

		msg := &sarama.ProducerMessage{Topic: tenant +  ".data." + event.StreamName, Key: sarama.StringEncoder(event.Key), Value: sarama.StringEncoder(event.Body)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("FAILED to send message: %s\n", err)
		} else {
			log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
	} else if (path[1] == "service") {
		service := path[2]
		function := path[3]
		event := ServiceEvent{service, function, body}
		fmt.Println(writer, event)
	} else {
		fmt.Println(writer, "REST endpoint supports only data and service events.")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func tenant(headers *http.Header) string  {
	return "anonymous"
}

// Data types

type ServiceEvent struct {
	Service  string
	Function string
	Body     []byte
}

func (event ServiceEvent) String() string {
	return fmt.Sprintf("service=%s function=%s body=%s", event.Service, event.Function, event.Body)
}

type DataEvent struct {
	StreamName string
	Key        string
	Body       []byte
}