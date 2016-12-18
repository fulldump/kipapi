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
	HookList       func(d *Context, c *golax.Context)
	HookPrint      func(d *Context, c *golax.Context)
}

type Context struct {
	Filter  bson.M
	Item    *kip.Item
	Printed map[string]interface{}
}

func New(pn *golax.Node, d *kip.Dao) *Kipapi {
	k := &Kipapi{
		Dao:        d,
		ParentNode: pn,
	}

	k.CollectionNode = pn.
		Node(d.Collection.Name).
		Method("GET", list(k)).
		Method("POST", create(k))

	k.ItemNode = k.CollectionNode.
		Node("{id}").
		Interceptor(newInterceptorId()).
		Interceptor(newInterceptorItem(k)).
		Method("GET", retrieve(k)).
		Method("DELETE", remove(k)).
		Method("PATCH", update(k))

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
