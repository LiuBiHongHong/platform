package kubeapi

import (
	"io"
	"log"
	"net/http"
)

// Deployment enables declarative updates for Pods and ReplicaSets.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#deployment-v1beta1-apps

// Create a Deployment
func CreateDeployment(apiserver string, ns string, body io.Reader) *http.Response {
	resp, err := http.Post(
		apiserver+"/apis/apps/v1beta1/namespaces/"+ns+"/deployments",
		"application/json",
		body)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Delete a Deployment
func DeleteDeployment(apiserver string, ns string, name string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/apis/apps/v1beta1/namespaces/"+ns+"/deployments/"+name,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Read the specified Deployment
func ReadDeployment(apiserver string, ns string, name string) *http.Response {
	resp, err := http.Get(apiserver + "/apis/apps/v1beta1/namespaces/" + ns + "/deployments/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// List or watch objects of kind Deployment
func ListDeployment(apiserver string, ns string) *http.Response {
	resp, err := http.Get(apiserver + "/apis/apps/v1beta1/namespaces/" + ns + "/deployments")
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Service is a named abstraction of software service (for example, mysql)
// consisting of local port (for example 3306) that the proxy listens on, and
// the selector that determines which pods will answer requests sent through
// the proxy.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#service-v1-core

// Create a Service
func CreateService(apiserver string, ns string, body io.Reader) *http.Response {
	resp, err := http.Post(
		apiserver+"/api/v1/namespaces/"+ns+"/services",
		"application/json",
		body)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Delete a Service
func DeleteService(apiserver string, ns string, name string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/namespaces/"+ns+"/services/"+name,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Read the specified Service
func ReadService(apiserver string, ns string, name string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/services/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// List or watch objects of kind Service
func ListService(apiserver string, ns string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/services")
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// PersistentVolumeClaim is a user's request for and claim to a persistent
// volume. A PersistentVolume must be allocated in the cluster to use this.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#persistentvolumeclaim-v1-core

// Create a PersistentVolumeClaim
func CreatePersistentVolumeClaim(apiserver string, ns string, body io.Reader) *http.Response {
	resp, err := http.Post(
		apiserver+"/api/v1/namespaces/"+ns+"/persistentvolumeclaims",
		"application/json",
		body)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Delete a PersistentVolumeClaim
func DeletePersistentVolumeClaim(apiserver string, ns string, name string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/namespaces/"+ns+"/persistentvolumeclaims/"+name,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Read the specified PersistentVolumeClaim
func ReadPersistentVolumeClaim(apiserver string, ns string, name string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/persistentvolumeclaims/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// List or watch objects of kind PersistentVolumeClaim
func ListPersistentVolumeClaim(apiserver string, ns string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/namespaces/" + ns + "/persistentvolumeclaims")
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// PersistentVolume (PV) is a storage resource provisioned by an administrator.
// It is analogous to a node.
//
// For more information, please check the official documentation:
//   https://kubernetes.io/docs/api-reference/v1.6/#persistentvolume-v1-core

// Create a PersistentVolume
func CreatePersistentVolume(apiserver string, body io.Reader) *http.Response {
	resp, err := http.Post(
		apiserver+"/api/v1/persistentvolumes",
		"application/json",
		body)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Delete a PersistentVolume
func DeletePersistentVolume(apiserver string, name string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE",
		apiserver+"/api/v1/persistentvolumes/"+name,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Read the specified PersistentVolume
func ReadPersistentVolume(apiserver string, name string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/persistentvolumes/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// List or watch objects of kind PersistentVolume
func ListPersistentVolume(apiserver string) *http.Response {
	resp, err := http.Get(apiserver + "/api/v1/persistentvolumes")
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
