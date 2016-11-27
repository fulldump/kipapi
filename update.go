package kipapi

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func update(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		patches := []*kip.Patch{}
		err := json.NewDecoder(c.Request.Body).Decode(&patches)
		if nil != err {
			c.Error(http.StatusBadRequest, "Body should be a valid JSON array of patches")
			return
		}

		// TODO: Put here the HOOK !

		i := GetItem(c)
		for _, p := range patches {
			if patchErr := i.Patch(p); nil != patchErr {
				c.Error(http.StatusBadRequest, patchErr.Error())
			}
		}

		if nil != i.Save() {
			c.Error(http.StatusInternalServerError, "Unexpected error patching this")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)
	}
}
