package platform

import (
	"fmt"
	"log"
)

var (
	// The slice containing all services on the machine
	services []*Service
)

func init() {
	services = make([]*Service, 0)
}

type Service struct {
	Id   string `json:"id"`
	App  string `json:"app"`
	Port string `json:"port"`

	// Status code of an app
	//
	// 0: initialized
	Status int `json:"status"`
}

func NewService() *Service {
	s := new(Service)
	s.Status = 0
	return s
}

func GetAllService() {
	// TODO
}

func GetService() {
	// TODO
}

func DeleteService() {
	// TODO
}

func Test() {
	a, _ := GetApp("6af06892-369d-4f2a-9a63-668b9f9e2044")
	t, err := setConfig(a.Configs)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t)
}
