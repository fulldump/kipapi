package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookCreate_OK(c *C) {

	// Hooks
	w.KipapiUsers.HookCreate = func(d *Context, c *golax.Context) {
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

func (w *World) Test_HookCreate_Conditional(c *C) {

	// Hooks
	w.KipapiUsers.HookCreate = func(d *Context, c *golax.Context) {

		i_am_admin_from_golax_context := true
		if i_am_admin_from_golax_context {
			// Add default value
			d.Item.Value.(*User).Email = "administrator@email.com"
		}
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
	c.Assert(r.StatusCode, Equals, http.StatusCreated)

	body := *r.BodyJsonMap()

	c.Assert(body["email"], Equals, "administrator@email.com")

}
