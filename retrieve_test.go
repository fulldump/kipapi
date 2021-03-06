package kipapi

import (
	"encoding/json"
	"net/http"

	"fmt"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Retrieve_OK(c *C) {

	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Request John
	r := w.Apitest.Request("GET", fmt.Sprintf("/users/%s", id)).Do()

	// Check
	body := *r.BodyJsonMap()
	expected := map[string]interface{}{
		"_id":    id,
		"name":   "unnamed",
		"email":  "",
		"age":    json.Number("18"),
		"single": false,
	}

	c.Assert(body, DeepEquals, expected)
	c.Assert(r.StatusCode, Equals, http.StatusOK)
}

func (w *World) Test_Retrieve_NotFound(c *C) {

	id := "invented-id"

	// Request John
	r := w.Apitest.Request("GET", "/users/"+id).Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusNotFound)
	c.Assert(r.BodyString(), Equals, "")
}
