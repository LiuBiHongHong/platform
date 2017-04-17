package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	var s *Service
	s = NewService()
	assert.Equal(t, 1, s.Status)
}

func TestGetAllService(t *testing.T) {
	services, err := GetAllService()
	assert.Nil(t, err)
	assert.NotNil(t, services)
}

func TestGetService(t *testing.T) {
	s := new(Service)
	s.Id = "10000000"
	services = append(services, s)
	v, err := GetService("10000000")
	assert.Nil(t, err)
	assert.Equal(t, "10000000", v.Id)
}

func TestDeleteService(t *testing.T) {}
