package kipapi

import (
	"net/http"

	"fmt"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Delete_OK(c *C) {

	// Create John
	john := w.Users.Create()
	john.Save()

	id := john.GetId()

	// Request John
	r := w.Apitest.Request("DELETE", fmt.Sprintf("/users/%s", id)).Do()

	// Check request
	c.Assert(r.StatusCode, Equals, http.StatusNoContent)
	c.Assert(r.BodyString(), Equals, "")

	// Check db
	user, err := w.Users.FindById(id)
	c.Assert(user, IsNil)
	c.Assert(err, IsNil)
}
