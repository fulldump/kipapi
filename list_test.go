package kipapi

import (
	"net/http"

	. "gopkg.in/check.v1"
)

func (w *World) Test_List_Empty(c *C) {

	r := w.Apitest.Request("GET", "/users/").Do()

	c.Assert(r.StatusCode, Equals, 200)
	c.Assert(r.BodyJson(), DeepEquals, []interface{}{})
}

func (w *World) Test_List_SeveralItems(c *C) {

	// Create sample users
	user1 := w.Users.Create()
	value1 := user1.Value.(*User)
	value1.Name = "John"
	user1.Save()

	user2 := w.Users.Create()
	value2 := user2.Value.(*User)
	value2.Name = "Peter"
	user2.Save()

	// Do request
	r := w.Apitest.Request("GET", "/users/").Do()

	// Check
	c.Assert(r.StatusCode, Equals, http.StatusOK)

	expected_body := []interface{}{
		map[string]interface{}{
			"_id": value1.Id,
		},
		map[string]interface{}{
			"_id": value2.Id,
		},
	}

	c.Assert(r.BodyJson(), DeepEquals, expected_body)

}
