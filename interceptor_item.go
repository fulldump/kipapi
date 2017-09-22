package kipapi

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"fmt"
)

func newInterceptorItem(k *Kipapi) *golax.Interceptor {

	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name:        "Object",
			Description: `Put a valid object in context.`,
		},
		Before: func(c *golax.Context) {

			id := c.Parameter

			d := &Context{
				Filter: bson.M{
					"_id": id,
				},
			}
			if nil != k.HookId {
				if k.HookId(d,c); nil != c.LastError {
					return
				}
			}

			if nil != k.HookFilter {
				if k.HookFilter(d, c); nil != c.LastError {
					return
				}
			}

			item := k.Dao.FindOne(d.Filter)

			if nil == item {
				m := fmt.Sprintf("Item `%s` not found", id)
				c.Error(http.StatusNotFound, m)
				return
			}

			c.Set("kipapi_item", item)
		},
	}

}

func GetItem(c *golax.Context) *kip.Item {
	object, exists := c.Get("kipapi_item")

	if exists {
		return object.(*kip.Item)
	}

	c.Error(http.StatusInternalServerError, "Something went terribly wrong getting object")
	return nil
}
