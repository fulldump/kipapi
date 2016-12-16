package kipapi

import (
	"encoding/json"
	"net/http"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

func update(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		d := &Context{
			Patches: []*kip.Patch{},
		}

		err := json.NewDecoder(c.Request.Body).Decode(&d.Patches)
		if nil != err {
			c.Error(http.StatusBadRequest, "Body should be a valid JSON array of patches")
			return
		}

		if nil != k.HookPatch {
			if k.HookPatch(d, c); nil != c.LastError {
				return
			}
		}

		i := GetItem(c)
		for _, d.Patch = range d.Patches {

			if nil != k.HookPatchItem {
				if k.HookPatchItem(d, c); nil != c.LastError {
					return
				}
			}

			if patchErr := i.Patch(d.Patch); nil != patchErr {
				c.Error(http.StatusBadRequest, patchErr.Error())
				return
			}
		}

		if nil != i.Save() {
			c.Error(http.StatusInternalServerError, "Unexpected error patching this")
			return
		}

		c.Response.WriteHeader(http.StatusNoContent)
	}
}
