package platform

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/liubihonghong/platform/kubelib"
	"github.com/mitchellh/mapstructure"
)

// Error returns the formatted service deletion error.
type ErrorDeleteService struct {
	id   string
	body string
}

// ErrorDeleteService denotes encountering an error while trying to delete an app.
func (e ErrorDeleteService) Error() string {
	return fmt.Sprintf("%s delete failed with response %s", string(e.id), string(e.body))
}

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
	// 100: ready
	// 98: have not been visited
	// 1: initialized
	Status int `json:"status"`
}

func NewService() *Service {
	s := new(Service)
	s.Status = 1
	return s
}

func GetAllService() ([]*Service, error) {
	if len(services) == 0 {
		resp, err := kubelib.ListAllService(apiserver)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		var data map[string]interface{}
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&data); err != nil {
			log.Println(err)
			return nil, err
		}
		for _, item := range data["items"].([]interface{}) {
			var svc kubelib.Service
			err := mapstructure.Decode(item, &svc)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			var s *Service
			s = NewService()
			s.Id = svc.Metadata.Name
			s.App = svc.Metadata.Namespace
			// TODO: Set service port.
			services = append(services, s)
		}
	}
	return services, nil
}

func GetService(id string) (*Service, error) {
	if len(services) == 0 {
		GetAllService()
	}

	for _, s := range services {
		if s.Id == id {
			s.Status = 100
			return s, nil
		}
	}
	return nil, ErrorNotFound{id, "service"}
}

func DeleteService() error {
	// TODO
	return nil
}
