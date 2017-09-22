package kipapi

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/golax"
	"gopkg.in/mgo.v2"
)

func create(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		d := &Context{
			Item: k.Dao.Create(),
		}

		if nil != k.HookCreate {
			if k.HookCreate(d, c); nil != c.LastError {
				return
			}
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

			if mgo.IsDup(err) {
				c.Error(http.StatusConflict, "Duplicated index: "+err.Error())
				return
			}

			c.Error(http.StatusInternalServerError, "Unexpected error saving object")
			return
		}

		c.Response.WriteHeader(http.StatusCreated)

		k.PrintItem(d.Item, c)

	}
}
