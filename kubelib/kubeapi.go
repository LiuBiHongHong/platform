package kubelib

import (
	"io"
	"log"
	"net/http"
)

type Deployment struct{}

type Service struct {
	Metadata struct {
		Name              string `json:"name"`
		Namespace         string `json:"namespace"`
		Uid               string `json:"uid"`
		ResourceVersion   string `json:"resourceVersion"`
		CreationTimestamp string `json:"creationTimestamp"`
		Labels            struct {
			Id string `json:"id"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			TargetPort int    `json:"targetPort"`
			Port       int    `json:"port"`
			Protocol   string `json:"protocol"`
			Name       string `json:"name"`
		} `json:"ports"`
		Selector struct {
			Id string `json:"id"`
		} `json:"selector"`
		ClusterIP string `json:"clusterIP"`
		Type      string `json:"type"`
	} `json:"spec"`
}

type PersistentVolumeClaim struct{}

type Namespace struct{}

type PersistentVolume struct{}

// Deployment enables declarative updates for Pods and ReplicaSets.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#deployment-v1beta1-apps

// Create a Deployment
func CreateDeployment(apiserver string, ns string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(
		apiserver+"/apis/apps/v1beta1/namespaces/"+ns+"/deployments",
		"application/json",
		body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Delete a Deployment
func DeleteDeployment(apiserver string, ns string, name string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/apis/apps/v1beta1/namespaces/"+ns+"/deployments/"+name,
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Read the specified Deployment
func ReadDeployment(apiserver string, ns string, name string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/apis/apps/v1beta1/namespaces/" + ns + "/deployments/" + name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind Deployment
func ListDeployment(apiserver string, ns string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/apis/apps/v1beta1/namespaces/" + ns + "/deployments")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Service is a named abstraction of software service (for example, mysql)
// consisting of local port (for example 3306) that the proxy listens on, and
// the selector that determines which pods will answer requests sent through
// the proxy.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#service-v1-core

// Create a Service
func CreateService(apiserver string, ns string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(
		apiserver+"/api/v1/namespaces/"+ns+"/services",
		"application/json",
		body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Delete a Service
func DeleteService(apiserver string, ns string, name string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/namespaces/"+ns+"/services/"+name,
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Read the specified Service
func ReadService(apiserver string, ns string, name string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/services/" + name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind Service
func ListService(apiserver string, ns string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/services")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind Service in all namespaces
func ListAllService(apiserver string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/services")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// PersistentVolumeClaim is a user's request for and claim to a persistent
// volume. A PersistentVolume must be allocated in the cluster to use this.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#persistentvolumeclaim-v1-core

// Create a PersistentVolumeClaim
func CreatePersistentVolumeClaim(apiserver string, ns string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(
		apiserver+"/api/v1/namespaces/"+ns+"/persistentvolumeclaims",
		"application/json",
		body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Delete a PersistentVolumeClaim
func DeletePersistentVolumeClaim(apiserver string, ns string, name string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/namespaces/"+ns+"/persistentvolumeclaims/"+name,
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Read the specified PersistentVolumeClaim
func ReadPersistentVolumeClaim(apiserver string, ns string, name string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/persistentvolumeclaims/" + name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind PersistentVolumeClaim
func ListPersistentVolumeClaim(apiserver string, ns string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/persistentvolumeclaims")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Namespace provides a scope for Names. Use of multiple namespaces is optional.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#namespace-v1-core

// Create a Namespace
func CreateNamespace(apiserver string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(
		apiserver+"/api/v1/namespaces",
		"application/json",
		body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Delete a Namespace
func DeleteNamespace(apiserver string, name string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/namespaces/"+name,
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Read the specified Namespace
func ReadNamespace(apiserver string, name string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind Namespace
func ListNamespace(apiserver string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/namespaces")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// PersistentVolume (PV) is a storage resource provisioned by an administrator.
// It is analogous to a node.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#persistentvolume-v1-core

// Create a PersistentVolume
func CreatePersistentVolume(apiserver string, body io.Reader) (*http.Response, error) {
	resp, err := http.Post(
		apiserver+"/api/v1/persistentvolumes",
		"application/json",
		body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Delete a PersistentVolume
func DeletePersistentVolume(apiserver string, name string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/persistentvolumes/"+name,
		nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// Read the specified PersistentVolume
func ReadPersistentVolume(apiserver string, name string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/persistentvolumes/" + name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

// List or watch objects of kind PersistentVolume
func ListPersistentVolume(apiserver string) (*http.Response, error) {
	resp, err := http.Get(apiserver + "/api/v1/persistentvolumes")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
