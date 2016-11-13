package kipapi

import (
	"net/http"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_Delete_OK(c *C) {

	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId().(bson.ObjectId)

	// Request John
	r := w.Apitest.Request("DELETE", "/users/"+id.Hex()).Do()

	// Check request
	c.Assert(r.StatusCode, Equals, http.StatusNoContent)
	c.Assert(r.BodyString(), Equals, "")

	// Check db
	user := w.Users.FindById(id)
	c.Assert(user, IsNil)
}
