package kipapi

import (
	"net/http"
	"strings"

	"github.com/fulldump/golax"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func (w *World) Test_Filter(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

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
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s",id)).
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

	user, err := w.Users.FindById(id)
	c.Assert(user.Value.(*User).Name, Equals, "·-{fulanito}-·")
	c.Assert(user.Value.(*User).Email, Equals, "myemail@google.com")
	c.Assert(err, IsNil)
}

func (w *World) Test_Filter_List(c *C) {
	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Create Mary
	mary_item := w.Users.Create()
	mary := mary_item.Value.(*User)
	mary.Name = "Mary"
	mary_item.Save()

	// Hooks
	w.KipapiUsers.HookFilter = func(d *Context, c *golax.Context) {
		d.Filter["name"] = bson.RegEx{Pattern: "M"}
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusOK)

	body := r.BodyJson().([]interface{})
	c.Assert(len(body), Equals, 1)

	user := body[0].(map[string]interface{})
	c.Assert(user["_id"], DeepEquals, mary.Id)

}

func (w *World) Test_Filter_Item(c *C) {
	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookFilter = func(d *Context, c *golax.Context) {
		d.Filter["name"] = bson.RegEx{Pattern: "M"}
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/"+john.Id).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNotFound)

	body := r.BodyString()
	c.Assert(body, Equals, "")

}
