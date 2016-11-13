package kipapi

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/golax"
)

func create(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		o := k.Dao.Create()

		if err := json.NewDecoder(c.Request.Body).Decode(o.Value); nil != err {
			c.Error(http.StatusBadRequest, "Body expected to be JSON: "+err.Error())
			return
		}

		if err := o.Save(); nil != err {
			c.Error(http.StatusInternalServerError, "Unexpected error saving object")
		}

		c.Response.WriteHeader(http.StatusCreated)
		json.NewEncoder(c.Response).Encode(o.Value)

	}
}
