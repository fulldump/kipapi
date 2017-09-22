package kipapi

import (
	"github.com/fulldump/golax"

	"encoding/json"

	"github.com/fulldump/kip"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_HookId_ByName(c *C) {

	// Create John
	john := w.Users.Create()
	john.Save()
	john.Patch(&kip.Patch{
		Operation: "set",
		Key:       "name",
		Value:     "John",
	})
	john.Save()

	// Hooks
	w.KipapiUsers.HookId = func(d *Context, c *golax.Context) {
		d.Filter = bson.M{"name": c.Parameter}
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/John").Do()

	// Check
	c.Assert(r.StatusCode, Equals, 200)

	c.Assert(r.BodyJson(), DeepEquals, map[string]interface{}{
		"_id":    john.GetId(),
		"age":    json.Number("18"),
		"email":  "",
		"name":   "John",
		"single": false,
	})
}
