package kipapi

import "github.com/fulldump/golax"

func retrieve(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		if nil != k.HookRetrieve {
			id := GetId(c)
			if k.HookRetrieve(id, c); nil != c.LastError {
				return
			}
		}

		i := GetItem(c)

		k.PrintItem(i, c)
	}
}
