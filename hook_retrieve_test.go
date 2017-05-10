package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_HookRetrieve(c *C) {

	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookRetrieve = func(id *bson.ObjectId, c *golax.Context) {
		c.Error(999, "Not allowed to create :D")
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/"+john.Id.Hex()).Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	c.Assert(r.BodyString(), Equals, "")
}
