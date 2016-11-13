package kipapi

import (
	"encoding/json"

	"github.com/fulldump/golax"
)

func retrieve(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		i := getItem(c)

		json.NewEncoder(c.Response).Encode(i.Value)
	}
}
