package kipapi

import "github.com/fulldump/golax"

func retrieve(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		i := GetItem(c)

		k.PrintItem(i, c)
	}
}
