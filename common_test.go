package kipapi

import (
	"strconv"
	"time"
)

func random_with_prefix(prefix string) string {
	return prefix + "_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}

type User struct {
	Id      string `bson:"_id"    json:"_id"`
	Name    string `bson:"name"   json:"name"`
	Email   string `bson:"email"  json:"email"`
	Age     int    `bson:"age"    json:"age"`
	Single  bool   `bson:"single" json:"single"`
	Private string `bson:"__private" json:"__private"`
}
