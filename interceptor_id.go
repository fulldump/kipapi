package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
	"gopkg.in/mgo.v2/bson"
)

func newInterceptorId() *golax.Interceptor {
	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name:        "ObjectId",
			Description: `Validate a url id and put it in context a ObjectId.`,
		},
		Before: func(c *golax.Context) {
			id := c.Parameter

			if !bson.IsObjectIdHex(id) {
				c.Error(http.StatusBadRequest, "Invalid format id")
				return
			}

			object_id := bson.ObjectIdHex(id)
			c.Set("object_id", &object_id)
		},
	}
}

func GetId(c *golax.Context) *bson.ObjectId {
	object_id, exists := c.Get("object_id")

	if !exists {
		return nil
	}

	return object_id.(*bson.ObjectId)
}
