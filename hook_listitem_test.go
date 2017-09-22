package kipapi

import (
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookListItem_RemoveItem(c *C) {
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
	w.KipapiUsers.HookListItem = func(d *Context, c *golax.Context) *kip.Item {
		user := d.Item.Value.(*User)
		if "Mary" == user.Name {
			return nil
		}
		return d.Item
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, 200)

	body := r.BodyJson().([]interface{})
	c.Assert(len(body), Equals, 1)

	user := body[0].(map[string]interface{})
	c.Assert(user["_id"], DeepEquals, john.Id)
}

func (w *World) Test_HookListItem_Error(c *C) {
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
	w.KipapiUsers.HookListItem = func(d *Context, c *golax.Context) *kip.Item {
		user := d.Item.Value.(*User)
		if "Mary" == user.Name {
			c.Error(999, "Error listing this file")
		}
		return d.Item
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, 999)

	c.Assert(r.BodyString(), Equals, "")
}
