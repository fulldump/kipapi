package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookPrint(c *C) {

	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookPrint = func(d *Context, c *golax.Context) {
		d.Printed = map[string]interface{}{
			"id": "This is the id: " + d.Item.Value.(*User).Id,
		}
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/"+john.Id).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusOK)

	c.Assert(r.BodyString(), Equals, `{"id":"This is the id: `+john.Id+`"}`+"\n")
}
