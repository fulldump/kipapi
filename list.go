package kipapi

import (
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/fulldump/golax"
)

func list(k *Kipapi) func(c *golax.Context) {
	return func(c *golax.Context) {

		d := &Context{
			Filter: bson.M{},
		}

		if nil != k.HookList {
			if k.HookList(d, c); nil != c.LastError {
				return
			}
		}

		if nil != k.HookFilter {
			if k.HookFilter(d, c); nil != c.LastError {
				return
			}
		}

		l := []interface{}{}

		fields := c.Request.URL.Query().Get("fields")
		f := strings.Split(fields+",_id", ",")

		iter, db := k.Dao.Find(d.Filter).Iter()
		defer db.Close()

		i := k.Dao.Create()
		for iter.Next(i.Value) {

			if nil != k.HookListItem {
				d.Item = i
				j := k.HookListItem(d, c)
				if nil != c.LastError {
					return
				}
				if nil == j {
					continue
				}
				i = j
			}

			m := k.Map(i, c)

			if fields != "" {
				m = map_item_fields(m, f)
			}

			l = append(l, m)

			i = k.Dao.Create() // Optional: ensure do not reuse previous values :)
		}

		if err := iter.Close(); err != nil {
			return
		}

		var p interface{} = l

		if nil != k.HookPrintList {
			d.PrintedList = l
			if p = k.HookPrintList(d, c); nil != c.LastError {
				return
			}
		}

		k.Encode(p, c)

	}
}
