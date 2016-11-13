package kipapi

import . "gopkg.in/check.v1"

func (w *World) Test_Update_OK(c *C) {
	// // Create John
	// john := w.Users.Create()
	// john.Save()

	// id := john.GetId().(bson.ObjectId)

	// // Request set name
	// r := w.Apitest.Request("PATCH", "/users/"+id.Hex()).
	// 	WithBodyString(`[
	// 		{
	// 			"operation": "set",
	// 			"key": "name",
	// 			"value": "Fulanito",
	// 		}
	// 	]`).Do()

	// // Check
	// c.Assert(r.StatusCode, Equals, http.StatusOK)

	// user := w.Users.FindById(id)
	// c.Assert(user.Value.(*User).Name, Equals, "Fulanito")
}
