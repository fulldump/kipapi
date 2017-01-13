package kipapi

import (
	"encoding/json"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/fulldump/golax"
	"github.com/fulldump/kip"
)

type Kipapi struct {
	Dao            *kip.Dao
	ParentNode     *golax.Node
	CollectionNode *golax.Node
	ItemNode       *golax.Node
	HookFilter     func(d *Context, c *golax.Context)
	HookList       func(d *Context, c *golax.Context)
	HookPrint      func(d *Context, c *golax.Context)
	HookPatch      func(d *Context, c *golax.Context)
	HookPatchItem  func(d *Context, c *golax.Context)
	HookInsert     func(d *Context, c *golax.Context)
}

type Context struct {
	Filter  bson.M
	Item    *kip.Item
	Printed map[string]interface{}
	Patches []*kip.Patch
	Patch   *kip.Patch
}

func New(pn *golax.Node, d *kip.Dao) *Kipapi {
	k := &Kipapi{
		Dao:        d,
		ParentNode: pn,
	}

	k.CollectionNode = pn.
		Node(d.Collection.Name).
		Method("GET", list(k), golax.Doc{
			Name: `List collection`,
			Description: `

			Retrieve all items inside a collection.

			**Fields**

			By default, only ´_id´ field is returned. If you want to get more
			fields, add the query parameter ´fields´ with the list of fields
			you want. For example:

			´´´
			$API_URL/api/v1/users/?fields=name,age
			´´´

			The special field ´*´ will show all fields.

			**Filtering**

			- not implemented -

			**Sorting**

			- not implemented -

			**Pagination**

			- not implemented -

			Example:

			´´´sh
			curl $API_URL/api/v1/users/?fields=name
			´´´

			Response:

			´´´json
			[
				{"_id":"57f3fbfba4cc8b6afe878240","name":"Fulanez"},
				{"_id":"57f40219ce507665708f33a8","name":"Menganez"},
				{"_id":"57f40256ce507665708f33a9","name":"Matusalen"}
			]
			´´´

			`,
		}).
		Method("POST", create(k), golax.Doc{
			Name: `Create item`,
			Description: `
			Create one item inside a collection.

			The ´_id´ field is autogenerated.

			Example curl:

			´´´sh
			curl $API_URL/api/v1/users/ -d '{"name":"Zutanez"}'
			´´´

			Response:
			
			´´´json
			{"_id":"58013557ce507602551f5faa","name":"Zutanez"}
			´´´
			`,
		})

	k.ItemNode = k.CollectionNode.
		Node("{id}").
		Interceptor(newInterceptorId()).
		Interceptor(newInterceptorItem(k)).
		Method("GET", retrieve(k), golax.Doc{
			Name: `Retrieve item`,
			Description: `
			Retrieve a single item with all its fields. Each item is
			identified by the field ´_id´.

			Example:

			´´´sh
			curl $API_URL/api/v1/users/58014862ce50765183348a14
			´´´

			Response:

			´´´json
			{"_id":"58014862ce50765183348a14","name":"Zutanez"}
			´´´
			`,
		}).
		Method("PATCH", update(k), golax.Doc{
			Name: `Update item`,
			Description: `
			This will update an item with a bunch or patches. A patch is only
			one atomic modification, it consists on:

			* ´operation´ - for example ´set´, ´unset´, ´inc´...
			* ´key´ - the key to be modified in dot notation, for example, in 
			a document with nested objects: ´friend.address.street´.
			* ´value´ - value for the operation.

			´´´json
			[
				{"operation":"set", "key":"name",    "value":"Fulano"},
				{"operation":"set", "key":"surname", "value":"Menganez"},
				{"operation":"inc", "key":"age",     "value":1},
				{"operation":"mul", "key":"price",   "value":1.21},
			]
			´´´

			Implemented operations:

			**set** put a value in a key

			´´´json
			{"operation":"set", "key":"address.street.number", "value":3}
			´´´

			**unset** remove a key

			´´´json
			{"operation":"unset", "key":"address.street.number"}
			´´´

			**inc** increment a numeric key

			´´´json
			{"operation":"inc", "key":"address.street.number", "value":-2}
			´´´

			**mul** multiply a numeric key

			´´´json
			{"operation":"mul", "key":"price", "value":1.21}
			´´´

			**push** push a value in array

			´´´json
			{"operation":"push", "key":"products", "value":"apple"}
			´´´

			**pull** remove a value from array

			´´´json
			{"operation":"push", "key":"products", "value":"pear"}
			´´´

			**Future work**
			´add´, ´pop´

			`,
		}).
		Method("DELETE", remove(k), golax.Doc{
			Name: `Delete item`,
			Description: `
			Remove document.

			This action is permanent by default.

			Example:

			´´´sh
			curl -X DELETE $API_URL/api/v1/users/58014862ce50765183348a14
			´´´

			Response:

			´´´json
			< empty response >
			´´´

			`,
		})

	return k
}

func (k *Kipapi) Map(i *kip.Item, c *golax.Context) bson.M {

	m := interface_to_map(i.Value)

	for k, _ := range m {
		if strings.HasPrefix(k, "__") {
			delete(m, k)
		}
	}

	if nil != k.HookPrint {
		d := &Context{
			Item:    i,
			Printed: m,
		}
		k.HookPrint(d, c)
		m = d.Printed
	}

	return m
}

func (k *Kipapi) PrintItem(i *kip.Item, c *golax.Context) {

	m := k.Map(i, c)

	json.NewEncoder(c.Response).Encode(m)
}
