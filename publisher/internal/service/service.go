package service

import (
	"fmt"
	"flag"
	"log"
	"encoding/json"
	"service/internal/config"
	"service/internal/models"
	"github.com/nats-io/nats.go"
)

type Service interface {
	Echo(number int) string
	EmployeeAdd(request *models.Request) string
}

type service struct{}

func NewService(cfg config.Service) (Service, error) {
	return &service{}, nil
}

func (s *service) Echo(number int) string {
	return fmt.Sprintf("%d", number)
}

func (s *service) EmployeeAdd(request *models.Request) string {
	var urls = flag.String("s", "stan-nats:4222", "The nats server URLs (separated by comma)")
	var reply = flag.String("reply", "", "Sets a specific reply subject")

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		log.Panicf("Connection error: %v", err)
	}
	defer nc.Close()
	
	jsonData, err := json.Marshal(request)
	subj, msg := "test", jsonData

	if reply != nil && *reply != "" {
		nc.PublishRequest(subj, *reply, msg)
	} else {
		nc.Publish(subj, msg)
	}

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Panicf("Last error: %v", err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
	return "OK"
}
