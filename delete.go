package kipapi

import (
	"net/http"

	"github.com/fulldump/golax"
)

func delete(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		i := getItem(c)

		if err := i.Delete(); nil != err {
			c.Error(http.StatusInternalServerError, "Unexpected error deleting object")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)
	}
}
