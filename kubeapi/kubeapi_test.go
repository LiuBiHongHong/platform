package kubeapi

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
	resp := CreateDeployment(apiserver, "default", bytes.NewBuffer(jsonStr))
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadDeployment(t *testing.T) {
	resp := ReadDeployment(apiserver, "default", "deployment-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListDeployment(t *testing.T) {
	resp := ListDeployment(apiserver, "default")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteDeployment(t *testing.T) {
	resp := DeleteDeployment(apiserver, "default", "deployment-example")
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
	resp := CreateService(apiserver, "default", bytes.NewBuffer(jsonStr))
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadService(t *testing.T) {
	resp := ReadService(apiserver, "default", "service-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListService(t *testing.T) {
	resp := ListService(apiserver, "default")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteService(t *testing.T) {
	resp := DeleteService(apiserver, "default", "service-example")
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
	resp := CreatePersistentVolumeClaim(apiserver, "default", bytes.NewBuffer(jsonStr))
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadPersistentVolumeClaim(t *testing.T) {
	resp := ReadPersistentVolumeClaim(apiserver, "default", "pvc-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListPersistentVolumeClaim(t *testing.T) {
	resp := ListPersistentVolumeClaim(apiserver, "default")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeletePersistentVolumeClaim(t *testing.T) {
	resp := DeletePersistentVolumeClaim(apiserver, "default", "pvc-example")
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
	resp := CreateNamespace(apiserver, bytes.NewBuffer(jsonStr))
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadNamespace(t *testing.T) {
	resp := ReadNamespace(apiserver, "namespace-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListNamespace(t *testing.T) {
	resp := ListNamespace(apiserver)
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeleteNamespace(t *testing.T) {
	resp := DeleteNamespace(apiserver, "namespace-example")
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
	resp := CreatePersistentVolume(apiserver, bytes.NewBuffer(jsonStr))
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestReadPersistentVolume(t *testing.T) {
	resp := ReadPersistentVolume(apiserver, "pv-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestListPersistentVolume(t *testing.T) {
	resp := ListPersistentVolume(apiserver)
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}

func TestDeletePersistentVolume(t *testing.T) {
	resp := DeletePersistentVolume(apiserver, "pv-example")
	if resp.StatusCode >= 400 {
		t.Error("Expected", 200, "got", resp.StatusCode)
	}
}
