package kipapi

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/golax"
)

func create(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		if nil != k.HookCreate {
			if k.HookCreate(c); nil != c.LastError {
				return
			}
		}

		d := &Context{
			Item: k.Dao.Create(),
		}

		if err := json.NewDecoder(c.Request.Body).Decode(d.Item.Value); nil != err {
			c.Error(http.StatusBadRequest, "Body expected to be JSON: "+err.Error())
			return
		}

		if nil != k.HookInsert {
			if k.HookInsert(d, c); nil != c.LastError {
				return
			}
		}

		if err := d.Item.Save(); nil != err {
			c.Error(http.StatusInternalServerError, "Unexpected error saving object")
		}

		c.Response.WriteHeader(http.StatusCreated)

		k.PrintItem(d.Item, c)

	}
}
