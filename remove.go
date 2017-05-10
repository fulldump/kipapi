package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
)

func remove(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		if nil != k.HookDelete {
			id := GetId(c)
			if k.HookDelete(id, c); nil != c.LastError {
				return
			}
		}

		i := GetItem(c)

		if err := i.Delete(); nil != err {
			c.Error(http.StatusInternalServerError, "Unexpected error deleting object")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)
	}
}
