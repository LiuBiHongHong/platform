package kubelib

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

// Error returns the config parsing error.
type ConfigError string

// ConfigError denotes encountering an error while trying to parsing configs.
func (e ConfigError) Error() string {
	return fmt.Sprintf("%s config parsing failed", string(e))
}

type Deployment struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		// TODO: Change later
		Labels struct {
			App  string `json:"app"`
			Type string `json:"type"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Strategy struct {
			Type string `json:"type"`
		} `json:"strategy"`
		Replicas *int `json:"replicas"`
		Template struct {
			Metadata struct {
				// TODO: Change later
				Labels struct {
					App  string `json:"app"`
					Tier string `json:"tier"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Image string `json:"image"`
					Name  string `json:"name"`
					Env   []struct {
						Name      string  `json:"name"`
						Value     *string `json:"value"`
						ValueFrom *struct {
							FieldRef struct {
								FieldPath string `json:"fieldPath"`
							} `json:"fieldRef"`
						} `json:"valueFrom"`
					} `json:"env"`
					Args  []string `json:"args"`
					Ports []struct {
						ContainerPort *int   `json:"containerPort"`
						Name          string `json:"name"`
					} `json:"ports"`
					VolumeMounts []struct {
						Name      string `json:"name"`
						MountPath string `json:"mountPath"`
					} `json:"volumeMounts"`
				} `json:"containers"`
				Volumes []struct {
					Name                  string `json:"name"`
					PersistentVolumeClaim struct {
						ClaimName string `json:"claimName"`
					} `json:"persistentVolumeClaim"`
				} `json:"volumes"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

type Service struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		// TODO: Change later
		Labels struct {
			App  string `json:"app"`
			Type string `json:"type"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			TargetPort *int   `json:"targetPort"`
			Port       *int   `json:"port"`
			Protocol   string `json:"protocol"`
			Name       string `json:"name"`
		} `json:"ports"`
		// TODO: Change later
		Selector struct {
			App  string `json:"app"`
			Tier string `json:"tier"`
		} `json:"selector"`
		ClusterIP string `json:"clusterIP"`
		Type      string `json:"type"`
	} `json:"spec"`
}

type PersistentVolumeClaim struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		// TODO: Change later
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		AccessModes []string `json:"accessModes"`
		Resources   struct {
			Requests struct {
				Storage string `json:"storage"`
			} `json:"requests"`
		} `json:"resources"`
	} `json:"spec"`
}

type Namespace struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
}

type PersistentVolume struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
		// TODO: Change later
		Labels struct {
			App  string `json:"app"`
			Type string `json:"type"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Capacity struct {
			Storage string `json:"storage"`
		} `json:"capacity"`
		AccessModes []string `json:"accessModes"`
		HostPath    struct {
			Path string `json:"path"`
		} `json:"hostPath"`
	} `json:"spec"`
}

func ParseConfig(data []interface{}) ([]interface{}, error) {
	var configs []interface{}
	configs = make([]interface{}, 0)
	for _, c := range data {
		switch k := c.(map[string]interface{})["kind"].(string); k {
		case "Deployment":
			var deploy Deployment
			err := mapstructure.Decode(c, &deploy)
			if err != nil {
				log.Println(err)
				return nil, ConfigError(k)
			}
			configs = append(configs, deploy)
		case "Service":
			var svc Service
			err := mapstructure.Decode(c, &svc)
			if err != nil {
				log.Println(err)
				return nil, ConfigError(k)
			}
			configs = append(configs, svc)
		case "PersistentVolumeClaim":
			var pvc PersistentVolumeClaim
			err := mapstructure.Decode(c, &pvc)
			if err != nil {
				log.Println(err)
				return nil, ConfigError(k)
			}
			configs = append(configs, pvc)
		case "Namespace":
			var ns Namespace
			err := mapstructure.Decode(c, &ns)
			if err != nil {
				log.Println(err)
				return nil, ConfigError(k)
			}
			configs = append(configs, ns)
		case "PersistentVolume":
			var pv PersistentVolume
			err := mapstructure.Decode(c, &pv)
			if err != nil {
				log.Println(err)
				return nil, ConfigError(k)
			}
			configs = append(configs, pv)
		default:
			return nil, ConfigError(k)
		}
	}
	return configs, nil
}
