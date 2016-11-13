package kipapi

import . "gopkg.in/check.v1"

func (w *World) Test_List_Empty(c *C) {

	r := w.Apitest.Request("GET", "/users/").Do()

	c.Assert(r.StatusCode, Equals, 200)
	c.Assert(r.BodyJson(), DeepEquals, []interface{}{})
}

func (w *World) Test_List_SeveralItems(c *C) {

	// // Create sample users
	// user1 := w.Users.Create()
	// value1 := user1.Value.(*User)
	// value1.Name = "John"
	// user1.Save()

	// user2 := w.Users.Create()
	// value2 := user1.Value.(*User)
	// value2.Name = "Peter"
	// user2.Save()

	// // Do request
	// r := w.Apitest.Request("GET", "/users/").Do()

	// // Check
	// c.Assert(r.StatusCode, Equals, http.StatusOK)

	// obtained := r.BodyJson()
	// expected := []interface{}{
	// 	map[string]interface{}{
	// 		"id": value1.Id.Hex(),
	// 	},
	// 	map[string]interface{}{
	// 		"id": value2.Id.Hex(),
	// 	},
	// }

	// c.Assert(obtained, DeepEquals, expected)

	// fmt.Println(r.BodyString())
	// c.FailNow()

}
