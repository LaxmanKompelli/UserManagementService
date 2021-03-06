package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Users []User

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Username string        `bson:"Username" json:"Username"`
	Password string        `bson:"Password" json:"Password"`
	Email    string        `bson:"Email" json:"Email"`
	Name     string        `bson:"Name" json:"Name"`
	TOC      time.Time     `bson:"TOC" json:"TOC"`
	TTL      int64         `bson:"TTL" json:"TTL"`
	Accesses []AppAccess   `bson:"AppAccess" json:"AppAccess"`
	IsActive bool          `bson:"IsActive" json:"IsActive"`
}

type AppAccess struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Token  string        `bson:"Token" json:"Token"`
	TOC    time.Time     `bson:"TOC" json:"TOC"`
	TTL    int64         `bson:"TTL" json:"TTL"`
	RoleId bson.ObjectId `bson:"RoleId" json:"RoleId"`
	Paths  []AppPath     `bson:"AppPath" json:"AppPath"`
}

type AppPath struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Path        string        `bson:"Path" json:"Path"`
	AccessLevel int8          `bson:"AccessLevel" json:"AccessLevel"`
}
