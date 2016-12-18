package kipapi

import (
	"encoding/json"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/fulldump/golax"
)

func list(k *Kipapi) func(c *golax.Context) {
	return func(c *golax.Context) {

		d := &Context{
			Filter: bson.M{},
		}

		if nil != k.HookFilter {
			if k.HookFilter(d, c); nil != c.LastError {
				return
			}
		}

		if nil != k.HookList {
			if k.HookList(d, c); nil != c.LastError {
				return
			}
		}

		l := []interface{}{}

		fields := c.Request.URL.Query().Get("fields")
		f := strings.Split(fields+",id", ",")

		iter := k.Dao.Find(d.Filter).Iter()

		tmp := k.Dao.Create().Value
		for iter.Next(tmp) {
			m := interface2map(tmp)
			m = map_item_fields(m, f)

			l = append(l, m)

			tmp = k.Dao.Create().Value // Optional: ensure do not reuse previous values :)
		}

		if err := iter.Close(); err != nil {
			return
		}

		json.NewEncoder(c.Response).Encode(l)

	}
}
