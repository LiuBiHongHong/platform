package kubelib

type Deployment struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		Labels    struct {
			Id string `json:"id"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Strategy struct {
			Type string `json:"type"`
		} `json:"strategy"`
		Replicas *int `json:"replicas"`
		Template struct {
			Metadata struct {
				Labels struct {
					Id string `json:"id"`
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
						ContainerPort int    `json:"containerPort"`
						Name          string `json:"name"`
					} `json:"ports"`
					VolumeMounts []struct {
						Name      string `json:"name"`
						MountPath string `json:"mountPath"`
					} `json:"volumeMounts"`
				} `json:"containers"`
				Volumes []struct {
					Name     string `json:"name"`
					HostPath *struct {
						Path string `json:"path"`
					} `json:"hostPath"`
					PersistentVolumeClaim *struct {
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
		Labels    struct {
			Id string `json:"id"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			TargetPort *int   `json:"targetPort"`
			Port       *int   `json:"port"`
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

type PersistentVolumeClaim struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		Labels    struct {
			Id string `json:"id"`
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
		Name   string `json:"name"`
		Labels struct {
			Id string `json:"id"`
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
