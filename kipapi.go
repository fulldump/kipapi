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
}

type Context struct {
	Filter bson.M
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

func (k *Kipapi) Print(c *golax.Context, i *kip.Item) {

	m := interface2map(i.Value)

	for k, _ := range m {
		if strings.HasPrefix(k, "__") {
			delete(m, k)
		}
	}

	json.NewEncoder(c.Response).Encode(m)
}
