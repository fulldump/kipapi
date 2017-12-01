package kipapi

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"fmt"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
	"gopkg.in/mgo.v2"
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
				if k.HookId(d, c); nil != c.LastError {
					return
				}
			}

			if nil != k.HookFilter {
				if k.HookFilter(d, c); nil != c.LastError {
					return
				}
			}

			item, err := k.Dao.FindOne(d.Filter)

			if mgo.ErrNotFound == err {
				m := fmt.Sprintf("Item `%s` not found", id)
				c.Error(http.StatusNotFound, m)
				return
			}

			if err != nil {
				c.Error(http.StatusInternalServerError, err.Error())
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
