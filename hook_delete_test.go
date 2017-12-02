package kipapi

import (
	"fmt"

	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookDelete(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Hooks
	w.KipapiUsers.HookDelete = func(d *Context, c *golax.Context) {
		c.Error(999, "Not authorized to do this :_(")
	}

	// Request set name
	r := w.Apitest.Request("DELETE", fmt.Sprintf("/users/%s", id)).Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	user, err := w.Users.FindById(id)
	c.Assert(user, NotNil)
	c.Assert(err, IsNil)
}
