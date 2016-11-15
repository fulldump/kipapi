package kipapi

import "github.com/fulldump/golax"

func retrieve(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		i := getItem(c)

		k.Print(c, i)
	}
}
