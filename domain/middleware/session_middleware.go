/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:51
 * Copyright (c) 2020
 */

package middleware

import (
	"github.com/gin-contrib/sessions/mongo"
	"github.com/globalsign/mgo"
	"github.com/oktopriima/lemonilo/domain/config"
)

func InitSessions(conf config.Config) mongo.Store {
	mdb, err := mgo.Dial(conf.GetString("mongodb.address") + ":" + conf.GetString("mongodb.port"))
	if err != nil {
		panic(err)
	}

	coll := mdb.DB(conf.GetString("mongodb.database")).C("sessions")
	store := mongo.NewStore(coll, 3600, true, []byte("mark-one-apps"))

	return store

}
