package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/liubihonghong/platform/kubelib"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// The url of Kubernetes api server
	apiserver string = "http://127.0.0.1:8001"
	// The url of market
	market string = "http://127.0.0.1:15000"
	// The url of mongodb
	mongo string = "127.0.0.1:27000"
	// The slice containing all apps on the machine
	apps []*App
	// The slice containing all apps on the market
	items []*Item
)

func init() {
	apps = make([]*App, 0)
	items = make([]*Item, 0)
}

// Error returns the formatted not found error.
type ErrorNotFound struct {
	id   string
	kind string
}

// ErrorNotFound denotes failing to find a service, an app or an item.
func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("%s %s not found", string(e.id), string(e.kind))
}

// Error returns the formatted docker image download error.
type ErrorDownloadImage struct {
	id    string
	image string
}

// ErrorDownloadImage denotes encountering an error while trying to download
// docker images of an app.
func (e ErrorDownloadImage) Error() string {
	return fmt.Sprintf("%s %s download failed", string(e.id), string(e.image))
}

// Error returns the formatted app start error.
type ErrorStartApp struct {
	id   string
	body string
}

// ErrorStartApp denotes encountering an error while trying to start an app.
func (e ErrorStartApp) Error() string {
	return fmt.Sprintf("%s start failed with response %s", string(e.id), string(e.body))
}

type App struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Author      string        `json:"author"`
	PictureUrl  string        `json:"pictureUrl"`
	Description string        `json:"description"`
	Size        int           `json:"size"`
	Configs     []interface{} `json:"configs"`

	// Status code of an app
	//
	// 100: ready
	// 98: have not been visited
	// 80: starting a service
	// 77: error happened when starting a service
	// 50: on deleting
	// 22: error happened when deleting
	// 0: initialized
	Status int `json:"status"`
}

type Item struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	PictureUrl  string `json:"pictureUrl"`
	Description string `json:"description"`

	// Status code of an item
	//
	// 100: has been already downloaded
	// 80: on downloading
	// 77: error happened when downloading
	// 0: initialized
	Status int `json:"status"`
}

func NewApp() *App {
	a := new(App)
	a.Configs = make([]interface{}, 0)
	a.Status = 0
	return a
}

func NewItem() *Item {
	i := new(Item)
	i.Status = 0
	return i
}

func GetAllApp() ([]*App, error) {
	// Check if apps is empty. If apps is empty, get apps from the database then
	// return it. Otherwise, return apps.
	if len(apps) == 0 {
		session, err := mgo.Dial(mongo)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("dorry").C("app")
		var result []map[string]interface{}
		err = c.Find(bson.M{}).All(&result)

		// Initialize the apps
		for _, data := range result {
			var a *App
			a = NewApp()
			err := mapstructure.Decode(data, &a)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			a.Configs, err = parseConfig(data["configs"].([]interface{}))
			if err != nil {
				log.Println(err)
				return nil, err
			}
			a.Status = 100
			apps = append(apps, a)
		}
	}
	return apps, nil
}

func GetApp(id string) (*App, error) {
	if len(apps) == 0 {
		GetAllApp()
	}

	for _, a := range apps {
		if a.Id == id {
			a.Status = 100
			return a, nil
		}
	}
	return nil, ErrorNotFound{id, "app"}
}

func StartApp(id string) (*App, error) {
	if len(apps) == 0 {
		GetAllApp()
	}

	for _, a := range apps {
		if a.Id == id {
			a.Status = 80
			_, err := setConfig(a.Configs)
			if err != nil {
				log.Println(err)
			}
			for _, c := range a.Configs {
				switch k := structs.Map(c)["Kind"].(string); k {
				case "Deployment":
					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(&c)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					resp, err := kubelib.CreateDeployment(apiserver, c.(kubelib.Deployment).Metadata.Namespace, &b)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					var data map[string]interface{}
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(&data); err != nil {
						log.Println(err)
						return nil, err
					}
					if data["code"] != nil && data["code"].(float64) >= 400 {
						var b bytes.Buffer
						err := json.NewEncoder(&b).Encode(&data)
						if err != nil {
							log.Println(err)
							return nil, err
						}
						return nil, ErrorStartApp{id, strings.Trim(b.String(), "\n")}
					}
				case "Service":
					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(&c)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					resp, err := kubelib.CreateService(apiserver, c.(kubelib.Service).Metadata.Namespace, &b)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					var data map[string]interface{}
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(&data); err != nil {
						log.Println(err)
						return nil, err
					}
					if data["code"] != nil && data["code"].(float64) >= 400 {
						var b bytes.Buffer
						err := json.NewEncoder(&b).Encode(&data)
						if err != nil {
							log.Println(err)
							return nil, err
						}
						return nil, ErrorStartApp{id, strings.Trim(b.String(), "\n")}
					}
				case "PersistentVolumeClaim":
					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(&c)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					resp, err := kubelib.CreatePersistentVolumeClaim(apiserver, c.(kubelib.PersistentVolumeClaim).Metadata.Namespace, &b)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					var data map[string]interface{}
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(&data); err != nil {
						log.Println(err)
						return nil, err
					}
					if data["code"] != nil && data["code"].(float64) >= 400 {
						var b bytes.Buffer
						err := json.NewEncoder(&b).Encode(&data)
						if err != nil {
							log.Println(err)
							return nil, err
						}
						return nil, ErrorStartApp{id, strings.Trim(b.String(), "\n")}
					}
				case "Namespace":
					// TODO: Skip if namespace exists.
					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(&c)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					resp, err := kubelib.CreateNamespace(apiserver, &b)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					var data map[string]interface{}
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(&data); err != nil {
						log.Println(err)
						return nil, err
					}
					if data["code"] != nil && data["code"].(float64) == 409 {
						fmt.Printf("%s already exists\n", k)
					} else if data["code"] != nil && data["code"].(float64) >= 400 {
						var b bytes.Buffer
						err := json.NewEncoder(&b).Encode(&data)
						if err != nil {
							log.Println(err)
							return nil, err
						}
						return nil, ErrorStartApp{id, strings.Trim(b.String(), "\n")}
					}
				case "PersistentVolume":
					var b bytes.Buffer
					err := json.NewEncoder(&b).Encode(&c)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					resp, err := kubelib.CreatePersistentVolume(apiserver, &b)
					if err != nil {
						log.Println(err)
						return nil, err
					}
					var data map[string]interface{}
					dec := json.NewDecoder(resp.Body)
					if err := dec.Decode(&data); err != nil {
						log.Println(err)
						return nil, err
					}
					if data["code"] != nil && data["code"].(float64) >= 400 {
						var b bytes.Buffer
						err := json.NewEncoder(&b).Encode(&data)
						if err != nil {
							log.Println(err)
							return nil, err
						}
						return nil, ErrorStartApp{id, strings.Trim(b.String(), "\n")}
					}
				}
			}
			return a, nil
		}
	}
	return nil, ErrorNotFound{id, "app"}
}

func DeleteApp(id string) error {
	if len(apps) == 0 {
		GetAllApp()
	}

	session, err := mgo.Dial(mongo)
	if err != nil {
		log.Println(err)
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dorry").C("app")
	result := App{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err == mgo.ErrNotFound {
		log.Println(err)
		return err
	} else {
		err := c.Remove(bson.M{"id": id})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	for i, a := range apps {
		if a.Id == id {
			copy(apps[i:], apps[i+1:])
			apps[len(apps)-1] = nil
			apps = apps[:len(apps)-1]
			a = nil
			return nil
		}
	}
	return ErrorNotFound{id, "app"}
}

func DownloadApp(id string) (*App, error) {
	// Check if apps is empty
	if len(apps) == 0 {
		GetAllApp()
	}

	// Set the status of downloading item in market.
	i, err := GetItem(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	i.Status = 80

	resp, err := http.Get(market + "/api/getapp/" + id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var data map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
		return nil, err
	}

	// Initialize the app
	var a *App
	a = NewApp()
	a.Id = data["id"].(string)
	a.Name = data["name"].(string)
	a.Version = data["version"].(string)
	a.Author = data["author"].(string)
	a.PictureUrl = data["pictureUrl"].(string)
	a.Description = data["description"].(string)
	a.Configs, err = parseConfig(data["configs"].([]interface{}))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	session, err := mgo.Dial(mongo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dorry").C("app")
	result := App{}
	err = c.Find(bson.M{"id": a.Id}).One(&result)
	if err == mgo.ErrNotFound {
		err = c.Insert(a)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		apps = append(apps, a)
	} else {
		fmt.Printf("%s already exists\n", a.Name)
	}
	a.Status = 98
	i.Status = 100
	return a, nil
}

func GetAllItem() ([]*Item, error) {
	// Check if apps is empty
	if len(apps) == 0 {
		GetAllApp()
	}

	// Check if items is empty. If items is empty, send a request to get items
	// from the server then return it. Otherwise, return items.
	if len(items) == 0 {
		resp, err := http.Get(market + "/api/listallapp")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer resp.Body.Close()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		var data map[string][]map[string]interface{}
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&data); err != nil {
			log.Println(err)
			return nil, err
		}

		// Initialize the items
		for _, d := range data["res"] {
			var i *Item
			i = NewItem()
			err := mapstructure.Decode(d, &i)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			// TODO: Set the status of downloaded items to 100.
			items = append(items, i)
		}
	}
	return items, nil
}

func GetItem(id string) (*Item, error) {
	if len(items) == 0 {
		GetAllItem()
	}

	for _, i := range items {
		if i.Id == id {
			return i, nil
		}
	}
	return nil, ErrorNotFound{id, "item"}
}
