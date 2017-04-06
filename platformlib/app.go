package platformlib

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "platform/kubeapi"
)

var (
	// The url of Kubernetes api server
	apiserver string
	// The Kubernetes namespace used by the platform
	ns string
	// The url of market
	market string
	// The url of mongodb
	mongo string
	// The slice containing all apps on the machine
	apps []App
	// The slice containing all apps on the market
	items []Item
)

type App struct {
	id          string
	name        string
	version     string
	author      string
	pictureUrl  string
	description string
	size        int64
	configs     []interface{}

	// Status code of an app
	//
	// 100: has already installed
	// 80: starting a service
	// 77: error happened when starting a service
	// 50: on deleting
	// 22: error happened when deleting
	// 10: has been successfully deleted
	// 0: initialized
	status int64
}

type Item struct {
	id          string
	name        string
	version     string
	author      string
	pictureUrl  string
	description string

	// Status code of an item
	//
	// 100: has already installed
	// 80: on downloading
	// 77: error happened when downloading
	// 10: not installed
	// 0: initialized
	status int64
}

func init() {
	apiserver = "http://127.0.0.1:8001"
	ns = "dorry-system"
	market = "http://localhost:15000"
	mongo = "http://localhost:27000"
	apps = make([]App, 0)
	items = make([]Item, 0)
}

func NewApp() *App {
	a := new(App)
	a.configs = make([]interface{}, 0)
	a.status = 0

	return a
}

func NewItem() *Item {
	a := new(Item)
	a.status = 0

	return a
}

func DownloadApp(id string) {
	resp, err := http.Get(market + "/api/getapp/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data map[string]interface{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Fatal(err)
	}

	// Initialize the app
	var app *App
	app = NewApp()
	app.id = data["id"].(string)
	app.name = data["name"].(string)
	app.version = data["version"].(string)
	app.author = data["author"].(string)
	app.pictureUrl = data["pictureUrl"].(string)
	app.description = data["description"].(string)
	app.configs = data["configs"].([]interface{})

	// Connect to mongodb
	session, err := mgo.Dial(mongo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dorry").C("app")

	result := App{}
	err = c.Find(bson.M{"id": app.id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	if &result != nil {
		return
	} else {
		err = c.Insert(app)
		if err != nil {
			log.Fatal(err)
		}
	}
}
