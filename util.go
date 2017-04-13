package platform

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/structs"
	"github.com/liubihonghong/platform/kubelib"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/xid"
)

// Error returns the config parsing error.
type ErrorParseConfig string

// ErrorParseConfig denotes encountering an error while trying to parsing
// starting configs.
func (e ErrorParseConfig) Error() string {
	return fmt.Sprintf("%s config parsing failed", string(e))
}

// Error returns the config field setting error.
type ErrorSetConfig string

// ErrorSetConfig denotes encountering an error while trying to set starting
// configs.
func (e ErrorSetConfig) Error() string {
	return fmt.Sprintf("%s config setting failed", string(e))
}

func parseConfig(data []interface{}) ([]interface{}, error) {
	var cs []interface{}
	cs = make([]interface{}, 0)
	for _, c := range data {
		switch k := c.(map[string]interface{})["kind"].(string); k {
		case "Deployment":
			var deploy kubelib.Deployment
			err := mapstructure.Decode(c, &deploy)
			if err != nil {
				log.Println(err)
				return nil, ErrorParseConfig(k)
			}
			cs = append(cs, deploy)
		case "Service":
			var svc kubelib.Service
			err := mapstructure.Decode(c, &svc)
			if err != nil {
				log.Println(err)
				return nil, ErrorParseConfig(k)
			}
			cs = append(cs, svc)
		case "PersistentVolumeClaim":
			var pvc kubelib.PersistentVolumeClaim
			err := mapstructure.Decode(c, &pvc)
			if err != nil {
				log.Println(err)
				return nil, ErrorParseConfig(k)
			}
			cs = append(cs, pvc)
		case "Namespace":
			var ns kubelib.Namespace
			err := mapstructure.Decode(c, &ns)
			if err != nil {
				log.Println(err)
				return nil, ErrorParseConfig(k)
			}
			cs = append(cs, ns)
		case "PersistentVolume":
			var pv kubelib.PersistentVolume
			err := mapstructure.Decode(c, &pv)
			if err != nil {
				log.Println(err)
				return nil, ErrorParseConfig(k)
			}
			cs = append(cs, pv)
		default:
			return nil, ErrorParseConfig(k)
		}
	}
	return cs, nil
}

func setConfig(cs []interface{}) ([]interface{}, error) {
	sid := xid.New().String()
	pvcl := make([]string, 0)

	for i, c := range cs {
		switch k := structs.Map(c)["Kind"].(string); k {
		case "Deployment":
			tmp := c.(kubelib.Deployment)
			tmp.Metadata.Name = sid
			tmp.Metadata.Labels.Id = sid
			// Set ClaimName field for the deployment
			for _, v := range tmp.Spec.Template.Spec.Volumes {
				if v.PersistentVolumeClaim != nil {
					if len(pvcl) == 0 {
						return nil, ErrorSetConfig("ClaimName")
					}
					v.PersistentVolumeClaim.ClaimName = pvcl[0]
					pvcl = pvcl[1:]
				}
			}
			cs[i] = tmp
		case "Service":
			tmp := c.(kubelib.Service)
			tmp.Metadata.Name = sid
			tmp.Metadata.Labels.Id = sid
			// Set Port field for the service
			// for _, p := range tmp.Spec.Ports {
			// 	fp, err := getFreePort()
			// 	if fp == 0 || err != nil {
			// 		log.Println(err)
			// 		return nil, ErrorSetConfig("Port")
			// 	}
			// 	*p.Port = fp
			// }
			cs[i] = tmp
		case "PersistentVolumeClaim":
			tmp := c.(kubelib.PersistentVolumeClaim)
			pvcid := xid.New().String()
			tmp.Metadata.Name = pvcid
			tmp.Metadata.Labels.Id = sid
			cs[i] = tmp
			pvcl = append(pvcl, pvcid)
		case "PersistentVolume":
			tmp := c.(kubelib.PersistentVolume)
			pvid := xid.New().String()
			tmp.Metadata.Name = pvid
			tmp.Metadata.Labels.Id = sid
			cs[i] = tmp
		}
	}
	return cs, nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
