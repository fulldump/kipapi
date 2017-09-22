package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
)

func remove(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		d := &Context{
			Item: GetItem(c),
		}

		if nil != k.HookDelete {
			if k.HookDelete(d, c); nil != c.LastError {
				return
			}
		}

		if err := d.Item.Delete(); nil != err {
			c.Error(http.StatusInternalServerError, "Unexpected error deleting object")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)
	}
}
