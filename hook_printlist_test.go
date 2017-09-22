package kipapi

import (
	"net/http"

	"encoding/json"

	"github.com/fulldump/golax"
	. "gopkg.in/check.v1"
)

func (w *World) Test_HookPrintList_OK(c *C) {

	// Create John
	john_item := w.Users.Create()
	john := john_item.Value.(*User)
	john.Name = "John"
	john_item.Save()

	// Hooks
	w.KipapiUsers.HookPrintList = func(d *Context, c *golax.Context) interface{} {
		return map[string]interface{}{
			"users": d.PrintedList,
		}
	}

	// Request set name
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusOK)

	body := *r.BodyJsonMap()
	obtained := body["users"].([]interface{})[0].(map[string]interface{})
	expected := map[string]interface{}{
		"_id":    john.Id,
		"age":    json.Number("18"),
		"email":  "",
		"name":   "John",
		"single": false,
	}

	for k, v := range expected {
		c.Assert(v, DeepEquals, obtained[k])
	}

}
