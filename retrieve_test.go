package kipapi

import (
	"encoding/json"
	"net/http"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_Retrieve_OK(c *C) {

	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId().(bson.ObjectId).Hex()

	// Request John
	r := w.Apitest.Request("GET", "/users/"+id).Do()

	// Check
	body := *r.BodyJsonMap()
	expected := map[string]interface{}{
		"id":     id,
		"name":   "unnamed",
		"email":  "",
		"age":    json.Number("18"),
		"single": false,
	}

	c.Assert(body, DeepEquals, expected)
	c.Assert(r.StatusCode, Equals, http.StatusOK)
}

func (w *World) Test_Retrieve_MalformedId(c *C) {

	// Request John
	r := w.Apitest.Request("GET", "/users/malformed_id").Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusBadRequest)
	c.Assert(r.BodyString(), Equals, "")
}

func (w *World) Test_Retrieve_NotFound(c *C) {

	id := bson.NewObjectId()

	// Request John
	r := w.Apitest.Request("GET", "/users/"+id.Hex()).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNotFound)
	c.Assert(r.BodyString(), Equals, "")
}
