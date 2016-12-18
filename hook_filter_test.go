package kipapi

import (
	"net/http"
	"strings"

	"github.com/fulldump/golax"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (w *World) Test_Filter(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId().(bson.ObjectId)

	// Hooks
	w.KipapiUsers.HookPatch = func(d *Context, c *golax.Context) {
		// All fields are uppercased
		for _, p := range d.Patches {
			p.Value = strings.ToLower(p.Value.(string))
		}
	}

	w.KipapiUsers.HookPatchItem = func(d *Context, c *golax.Context) {
		// Decorate field name :D
		p := d.Patch
		if "set" == p.Operation && "name" == p.Key {
			p.Value = "·-{" + p.Value.(string) + "}-·"
		}
	}

	// Request set name
	r := w.Apitest.Request("PATCH", "/users/"+id.Hex()).
		WithBodyString(`[
			{
				"operation": "set",
				"key": "name",
				"value": "FUlaniTo"
			}
			,
			{
				"operation": "set",
				"key": "email",
				"value": "myEmAil@gOOglE.COM"
			}
		]`).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNoContent)

	user := w.Users.FindById(id)
	c.Assert(user.Value.(*User).Name, Equals, "·-{fulanito}-·")
	c.Assert(user.Value.(*User).Email, Equals, "myemail@google.com")
}
