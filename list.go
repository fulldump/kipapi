package kipapi

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"github.com/fulldump/golax"
)

func list(k *Kipapi) func(c *golax.Context) {
	return func(c *golax.Context) {

		i := k.Dao.Create().Value

		l := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(i)), 0, 0).Interface()

		fmt.Printf("%#v \n", l)

		k.Dao.Find(bson.M{}).All(&l)

		// json.NewEncoder(c.Response).Encode(l)

	}
}
