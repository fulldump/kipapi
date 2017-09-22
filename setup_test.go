package kipapi

import (
	"testing"

	"github.com/fulldump/apitest"
	"github.com/fulldump/golax"
	"github.com/fulldump/kip"

	. "gopkg.in/check.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type World struct {
	MongoUri    string
	Api         *golax.Api
	Apitest     *apitest.Apitest
	Database    *kip.Database
	Kip         *kip.Kip
	KipapiUsers *Kipapi
	KipapiBooks *Kipapi
	Users       *kip.Dao
	Books       *kip.Dao
}

var _ = Suite(&World{
	MongoUri: random_with_prefix("localhost/kipapi-"),
})

func (w *World) SetUpSuite(c *C) {

	if db, err := kip.NewDatabase(w.MongoUri); err == nil {
		w.Database = db
	} else {
		panic("Fail creating a TESTING database. Please, check your MongoDB")
	}

	w.Kip = kip.NewKip()
	w.Kip.Define(&kip.Collection{
		Name: "users",
		OnCreate: func() interface{} {
			return &User{
				Id:   bson.NewObjectId().Hex(),
				Name: "unnamed",
				Age:  18,
			}
		},
	})
	w.Users = w.Kip.NewDao("users", w.Database)

	w.Kip.Define(&kip.Collection{
		Name: "books",
		OnCreate: func() interface{} {
			return map[string]interface{}{
				"_id":   bson.NewObjectId(),
				"title": "untitled",
				"pages": 0,
			}
		},
	})
	w.Books = w.Kip.NewDao("books", w.Database)

}

func (w *World) SetUpTest(c *C) {
	// Clean databases
	w.Users.Delete(bson.M{})
	w.Books.Delete(bson.M{})

	// Build api
	w.Api = golax.NewApi()
	w.KipapiUsers = New(w.Api.Root, w.Users)
	w.KipapiBooks = New(w.Api.Root, w.Books)

	w.Apitest = apitest.New(w.Api)
}

func (w *World) TearDownSuite(c *C) {
	// When all tests are finished, drop database
	session, _ := mgo.Dial(w.MongoUri)
	session.SetMode(mgo.Monotonic, true)
	session.DB("").DropDatabase()
	session.Close()
}

// Check world constraints
func (w *World) Test_World_NoUsers(c *C) {
	n, _ := w.KipapiUsers.Dao.Find(nil).Count()
	c.Assert(n, Equals, 0)
}
