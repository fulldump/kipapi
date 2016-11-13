package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func newInterceptorItem(k *Kipapi) *golax.Interceptor {

	return &golax.Interceptor{
		Documentation: golax.Doc{
			Name:        "Object",
			Description: `Put a valid object in context.`,
		},
		Before: func(c *golax.Context) {

			id := getId(c)

			item := k.Dao.FindById(id)

			if nil == item {
				c.Error(http.StatusNotFound, `Item '`+id.Hex()+`' not found.`)
				return
			}

			c.Set("kipapi_item", item)
		},
	}

}

func getItem(c *golax.Context) *kip.Item {
	object, exists := c.Get("kipapi_item")

	if exists {
		return object.(*kip.Item)
	}

	c.Error(http.StatusInternalServerError, "Something went terribly wrong getting object")
	return nil
}
