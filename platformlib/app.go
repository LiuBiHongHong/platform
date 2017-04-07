package platformlib

import (
	"encoding/json"
	"log"
	"net/http"

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

type App struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Author      string        `json:"author"`
	PictureUrl  string        `json:"pictureUrl"`
	Description string        `json:"description"`
	Size        int64         `json:"size"`
	Configs     []interface{} `json:"configs"`

	// Status code of an app
	//
	// 100: ready
	// 80: starting a service
	// 77: error happened when starting a service
	// 50: on deleting
	// 22: error happened when deleting
	// 10: has been successfully deleted
	// 0: initialized
	Status int64 `json:"status"`
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
	// 10: not installed
	// 0: initialized
	Status int64 `json:"status"`
}

func init() {
	apps = make([]*App, 0)
	items = make([]*Item, 0)
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

func GetAllApp() {}

func GetApp() {}

func StartApp() {}

func DeleteApp() {}

func DownloadApp(id string) (*App, error) {
	// Check if apps is empty
	if len(apps) == 0 {
		// TODO: Call GetAllApp() here if apps is empty.
	}

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
	a.Configs = data["configs"].([]interface{})

	// Connect to mongodb
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
			log.Fatal(err)
		}
		apps = append(apps, a)
	} else {
		log.Println(a.Name + " already exists")
	}
	a.Status = 100
	return a, nil
}

func GetAllItem() {}

func GetItem() {}

func GetAllService() {}

func GetService() {}

func DeleteService() {}
