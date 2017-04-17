package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	deploy := map[string]interface{}{
		"apiVersion": "apps/v1beta1",
		"kind":       "Deployment",
		"metadata":   nil,
		"spec":       nil}
	svc := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Service",
		"metadata":   nil,
		"spec":       nil}
	pvc := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "PersistentVolumeClaim",
		"metadata":   nil}
	ns := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Namespace",
		"metadata":   nil,
		"spec":       nil}
	pv := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "PersistentVolume",
		"metadata":   nil,
		"spec":       nil}
	cs := []interface{}{deploy, svc, pvc, ns, pv}
	v, err := parseConfig(cs)
	assert.Nil(t, err)
	assert.NotNil(t, v)
}

func TestSetConfig(t *testing.T) {}

func TestGetFreePort(t *testing.T) {
	p, err := getFreePort()
	assert.Nil(t, err)
	assert.NotEqual(t, p, 0)
}
