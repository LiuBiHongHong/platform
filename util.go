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
type ErrorSetConfigField string

// ErrorSetConfigField denotes encountering an error while trying to setting
// fields of starting configs.
func (e ErrorSetConfigField) Error() string {
	return fmt.Sprintf("%s config parsing failed", string(e))
}

func ParseConfig(data []interface{}) ([]interface{}, error) {
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

func SetConfigField(cs []interface{}) ([]interface{}, error) {
	guid := xid.New()
	fmt.Printf("%q\n", guid.String())

	for i, c := range cs {
		switch k := structs.Map(c)["Kind"].(string); k {
		case "Deployment":
			tmp := c.(kubelib.Deployment)
			tmp.Metadata.Name = "@@@"
			cs[i] = tmp
		case "Service":
			tmp := c.(kubelib.Service)
			tmp.Metadata.Name = "^^^"
			cs[i] = tmp
		case "PersistentVolumeClaim":
			tmp := c.(kubelib.PersistentVolumeClaim)
			tmp.Metadata.Name = "!!!"
			cs[i] = tmp
		case "Namespace":
			tmp := c.(kubelib.Namespace)
			tmp.Metadata.Name = "&&&"
			cs[i] = tmp
		case "PersistentVolume":
			tmp := c.(kubelib.PersistentVolume)
			tmp.Metadata.Name = "***"
			cs[i] = tmp
		}
	}
	return cs, nil
}

func GetFreePort() (int, error) {
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
