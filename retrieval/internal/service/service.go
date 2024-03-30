package service

import (
	"fmt"
	"log"
	"encoding/json"
	"service/internal/config"
	"service/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Service interface {
	Echo(number int) string
	EmployeeGet(number int) string
}

type service struct{}

func NewService(cfg config.Service) (Service, error) {
	return &service{}, nil
}

func (s *service) Echo(number int) string {
	return fmt.Sprintf("%d", number)
}

func (s *service) EmployeeGet(id int) string {
	db, err := sqlx.Connect(
		"postgres",
		"user=admin dbname=employees sslmode=disable password=123 host=database")
	if err != nil {
	   log.Panic("%v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
	   log.Fatal(err)
	} else {
	   log.Println("Successfuly connected")
	}

	employee := models.Employee{}
	resp := fmt.Sprintf("SELECT * FROM employee where employeeid = %v;", id)
	rows, _ := db.Queryx(resp)
	for rows.Next() {
	   	err := rows.StructScan(&employee)
	   	if err != nil {
			log.Panic("%v", err)
   		}	
   		log.Printf("%#v\n", employee)
  	}

	jsonData, err := json.Marshal(employee)

  	return string(jsonData)
}
