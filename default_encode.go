package kipapi

import (
	"encoding/json"

	"github.com/fulldump/golax"
)

func DefaultEncode(i interface{}, c *golax.Context) {
	json.NewEncoder(c.Response).Encode(i)
}
