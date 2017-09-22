package kipapi

import "github.com/fulldump/golax"

func retrieve(k *Kipapi) func(c *golax.Context) {

	return func(c *golax.Context) {

		d := &Context{
			Item: GetItem(c),
		}

		if nil != k.HookRetrieve {
			if k.HookRetrieve(d, c); nil != c.LastError {
				return
			}
		}

		k.PrintItem(d.Item, c)
	}
}
