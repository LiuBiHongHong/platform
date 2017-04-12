package kubelib

import (
	"bytes"
	"testing"
)

var (
	apiserver string = "http://127.0.0.1:8001"
)

func TestCreateDeployment(t *testing.T) {
	jsonStr := []byte(`
    {
      "apiVersion": "apps/v1beta1",
      "kind": "Deployment",
      "metadata": {
        "name": "deployment-example"
      },
      "spec": {
        "replicas": 3,
        "revisionHistoryLimit": 10,
        "template": {
          "metadata": {
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:1.10",
                "ports": [
                  {
                    "containerPort": 80
                  }
                ]
              }
            ]
          }
        }
      }
    }`)
	resp, err := CreateDeployment(apiserver, "default", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadDeployment(t *testing.T) {
	resp, err := ReadDeployment(apiserver, "default", "deployment-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListDeployment(t *testing.T) {
	resp, err := ListDeployment(apiserver, "default")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteDeployment(t *testing.T) {
	resp, err := DeleteDeployment(apiserver, "default", "deployment-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestCreateService(t *testing.T) {
	jsonStr := []byte(`
    {
      "kind": "Service",
      "apiVersion": "v1",
      "metadata": {
        "name": "service-example"
      },
      "spec": {
        "ports": [
          {
            "name": "http",
            "port": 80,
            "targetPort": 80
          }
        ],
        "selector": {
          "app": "nginx"
        },
        "type": "LoadBalancer"
      }
    }`)
	resp, err := CreateService(apiserver, "default", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadService(t *testing.T) {
	resp, err := ReadService(apiserver, "default", "service-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListService(t *testing.T) {
	resp, err := ListService(apiserver, "default")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteService(t *testing.T) {
	resp, err := DeleteService(apiserver, "default", "service-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestCreatePersistentVolumeClaim(t *testing.T) {
	jsonStr := []byte(`
    {
      "apiVersion": "v1",
      "kind": "PersistentVolumeClaim",
      "metadata": {
        "name": "pvc-example"
      },
      "spec": {
        "accessModes": [
          "ReadWriteOnce"
        ],
        "resources": {
          "requests": {
            "storage": "10Mi"
          }
        }
      }
    }`)
	resp, err := CreatePersistentVolumeClaim(apiserver, "default", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadPersistentVolumeClaim(t *testing.T) {
	resp, err := ReadPersistentVolumeClaim(apiserver, "default", "pvc-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListPersistentVolumeClaim(t *testing.T) {
	resp, err := ListPersistentVolumeClaim(apiserver, "default")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeletePersistentVolumeClaim(t *testing.T) {
	resp, err := DeletePersistentVolumeClaim(apiserver, "default", "pvc-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestCreateNamespace(t *testing.T) {
	jsonStr := []byte(`
    {
      "apiVersion": "v1",
      "kind": "Namespace",
      "metadata": {
        "name": "namespace-example"
      }
    }`)
	resp, err := CreateNamespace(apiserver, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadNamespace(t *testing.T) {
	resp, err := ReadNamespace(apiserver, "namespace-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListNamespace(t *testing.T) {
	resp, err := ListNamespace(apiserver)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteNamespace(t *testing.T) {
	resp, err := DeleteNamespace(apiserver, "namespace-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestCreatePersistentVolume(t *testing.T) {
	jsonStr := []byte(`
    {
      "apiVersion": "v1",
      "kind": "PersistentVolume",
      "metadata": {
        "name": "pv-example"
      },
      "spec": {
        "capacity": {
          "storage": "10Mi"
        },
        "accessModes": [
          "ReadWriteOnce"
        ],
        "hostPath": {
          "path": "/mnt/platforpv-example"
        }
      }
    }`)
	resp, err := CreatePersistentVolume(apiserver, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadPersistentVolume(t *testing.T) {
	resp, err := ReadPersistentVolume(apiserver, "pv-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListPersistentVolume(t *testing.T) {
	resp, err := ListPersistentVolume(apiserver)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeletePersistentVolume(t *testing.T) {
	resp, err := DeletePersistentVolume(apiserver, "pv-example")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}
