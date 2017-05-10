package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_HookUpdate(c *C) {

	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookUpdate = func(id *bson.ObjectId, c *golax.Context) {
		c.Error(999, "Not allowed to update D:")
	}

	// Request set name
	r := w.Apitest.Request("PATCH", "/users/"+john.Id.Hex()).
		WithBodyString(`
			[
				{"operation": "set", "key":"name", "value": "Jonny"}
			]
		`).
		Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	c.Assert(r.BodyString(), Equals, "")
}
