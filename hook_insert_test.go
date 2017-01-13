package kipapi

import (
	"encoding/json"
	"strings"

	"github.com/fulldump/golax"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Create_HookInsert_Sanitizer(c *C) {

	// Hook
	w.KipapiUsers.HookInsert = func(d *Context, c *golax.Context) {

		user := d.Item.Value.(*User)
		user.Name = strings.ToUpper(user.Name)

	}

	// Request set name
	r := w.Apitest.Request("POST", "/users/").
		WithBodyString(`{
			"name": "fulanito"
		}`).
		Do()

	// Check
	body := *r.BodyJsonMap()
	expected := map[string]interface{}{
		"_id":    body["_id"],
		"name":   "FULANITO",
		"email":  "",
		"age":    json.Number("18"),
		"single": false,
	}

	// Check returned object
	c.Assert(body, DeepEquals, expected)

	// Check object has not been inserted
	n, _ := w.KipapiUsers.Dao.Find(nil).Count()
	c.Assert(n, Equals, 1)
}

func (w *World) Test_Create_HookInsert_Validate(c *C) {

	// Hook
	w.KipapiUsers.HookInsert = func(d *Context, c *golax.Context) {

		c.Error(666, "There is a conflict!")

	}

	// Request set name
	r := w.Apitest.Request("POST", "/users/").
		WithBodyString(`{
			"name": "fulanito"
		}`).
		Do()

	// Check hook is executed
	c.Assert(r.StatusCode, Equals, 666)

	// Check object has not been inserted
	n, _ := w.KipapiUsers.Dao.Find(nil).Count()
	c.Assert(n, Equals, 0)
}
