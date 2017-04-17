package platform

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func TestNewApp(t *testing.T) {
	var a *App
	a = NewApp()
	assert.Equal(t, 1, a.Status)
	assert.NotNil(t, a.Configs)
}

func TestNewItem(t *testing.T) {
	var i *Item
	i = NewItem()
	assert.Equal(t, 1, i.Status)
}

func TestGetAllApp(t *testing.T) {
	apps, err := GetAllApp()
	assert.Nil(t, err)
	assert.NotNil(t, apps)
}

func TestGetApp(t *testing.T) {
	a := new(App)
	a.Id = "10000000"
	apps = append(apps, a)
	v, err := GetApp("10000000")
	assert.Nil(t, err)
	assert.Equal(t, "10000000", v.Id)
}

func TestStartApp(t *testing.T) {}

func TestDeleteApp(t *testing.T) {
	err := DeleteApp("10000000")
	assert.NotNil(t, err)

	session, err := mgo.Dial(mongo)
	assert.Nil(t, err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dorry").C("app")
	err = c.Insert(&App{"10000000", "", "", "", "", "", 0, nil, 0})
	assert.Nil(t, err)
	err = DeleteApp("10000000")
	assert.Nil(t, err)
	b := func() bool {
		for _, a := range apps {
			if a.Id == "10000000" {
				return false
			}
		}
		return true
	}()
	assert.Equal(t, true, b)
}

func TestDownloadApp(t *testing.T) {}

func TestGetAllItem(t *testing.T) {
	items, err := GetAllItem()
	assert.Nil(t, err)
	assert.NotNil(t, items)
}

func TestGetItem(t *testing.T) {
	i := new(Item)
	i.Id = "10000000"
	items = append(items, i)
	v, err := GetItem("10000000")
	assert.Nil(t, err)
	assert.Equal(t, "10000000", v.Id)
}
