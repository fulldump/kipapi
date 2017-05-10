package kipapi

import (
	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookCreate(c *C) {

	// Hooks
	w.KipapiUsers.HookCreate = func(c *golax.Context) {
		c.Error(999, "Not allowed to create :D")
	}

	// Request set name
	r := w.Apitest.Request("POST", "/users/").
		WithBodyString(`
			{
				"name": "fulanez"
			}
		`).
		Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	c.Assert(r.BodyString(), Equals, "")
}
