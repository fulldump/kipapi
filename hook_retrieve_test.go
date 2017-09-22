package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookRetrieve(c *C) {

	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookRetrieve = func(d *Context, c *golax.Context) {
		c.Error(999, "Not allowed to create :D")
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/"+john.Id).Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	c.Assert(r.BodyString(), Equals, "")
}
