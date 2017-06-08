package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookList(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	// Hooks
	w.KipapiUsers.HookList = func(d *Context, c *golax.Context) {
		c.Error(999, "Not authorized to list :_(")
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)
}
