package kipapi

import (
	"encoding/json"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func interface2map(i interface{}) map[string]interface{} {

	t, marshal_err := json.Marshal(i)
	if nil != marshal_err {
		panic(marshal_err)
	}

	o := map[string]interface{}{}

	unmarshal_err := json.Unmarshal(t, &o)
	if nil != unmarshal_err {
		panic(unmarshal_err)
	}

	return o
}

func map_item_fields(d bson.M, f []string) bson.M {
	r := bson.M{}

	all := wordInArray("*", f)

	for k, _ := range d {
		if "_id" == k {
			r["id"] = d["_id"]
			continue
		}

		if strings.HasPrefix(k, "__") {
			continue
		}

		if all || wordInArray(k, f) {
			r[k] = d[k]
		}
	}

	return r
}

/**
 * Return true if `w` is contained in `a`
 */
func wordInArray(w string, a []string) bool {
	for _, v := range a {
		if w == v {
			return true
		}
	}

	return false
}
