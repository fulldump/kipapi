package kipapi

import (
	"encoding/json"
	"net/http"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Create_OK(c *C) {

	// Request create
	r := w.Apitest.Request("POST", "/users/").
		WithBodyString(`{
			"name": "fulanito",
			"unexisting": "invented value for an invented field"
		}`).
		Do()

	// Check
	body := *r.BodyJsonMap()
	expected := map[string]interface{}{
		"_id":    body["_id"],
		"name":   "fulanito",
		"email":  "",
		"age":    json.Number("18"),
		"single": false,
	}

	c.Assert(body, DeepEquals, expected)
	c.Assert(r.StatusCode, Equals, http.StatusCreated)
}

func (w *World) Test_Create_MalformedBody(c *C) {

	// Request create
	r := w.Apitest.Request("POST", "/users/").
		WithBodyString(`}`).
		Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusBadRequest)

	body := r.BodyString()
	c.Assert(body, DeepEquals, "")
}
