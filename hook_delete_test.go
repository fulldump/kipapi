package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_HookDelete(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId().(bson.ObjectId)

	// Hooks
	w.KipapiUsers.HookDelete = func(id *bson.ObjectId, c *golax.Context) {
		c.Error(999, "Not authorized to do this :_(")
	}

	// Request set name
	r := w.Apitest.Request("DELETE", "/users/"+id.Hex()).Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	user := w.Users.FindById(id)
	c.Assert(user, NotNil)
}
