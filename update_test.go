package kipapi

import (
	"net/http"
	"strings"

	"github.com/fulldump/golax"

	"fmt"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Update_OK(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Request set name
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s", id)).
		WithBodyString(`[
			{
				"operation": "set",
				"key": "name",
				"value": "Fulanito"
			}
		]`).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNoContent)

	user := w.Users.FindById(id)
	c.Assert(user.Value.(*User).Name, Equals, "Fulanito")
}

func (w *World) Test_Update_BadRequest(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Request set name
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s", id)).
		WithBodyString(`}`).
		Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusBadRequest)

	body := r.BodyString()
	c.Assert(body, Equals, "")
}

func (w *World) Test_Update_HookPatch(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Hook
	w.KipapiUsers.HookPatch = func(d *Context, c *golax.Context) {
		for _, p := range d.Patches {
			p.Value = strings.ToUpper(p.Value.(string))
		}
	}

	// Request set name
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s", id)).
		WithBodyString(`[
			{
				"operation": "set",
				"key": "name",
				"value": "Fulanito"
			}
		]`).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNoContent)

	user := w.Users.FindById(id)
	c.Assert(user.Value.(*User).Name, Equals, "FULANITO")
}

func (w *World) Test_Update_BadPatch(c *C) {
	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Add interceptor to print api error:
	w.Api.Root.Interceptor(golax.InterceptorError)

	// Request set name
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s", id)).
		WithBodyString(`[
			{
				"operation": "seta",
				"key": "name",
				"value": "FUlaniTo"
			}
		]`).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusBadRequest)

	body := r.BodyJsonMap()
	c.Assert((*body)["Description"], DeepEquals, "invalid operation")
}

func (w *World) Test_Update_HookPatchCombined(c *C) {
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
	r := w.Apitest.Request("PATCH", fmt.Sprintf("/users/%s", id)).
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
