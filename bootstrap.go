package main

import (
	validator "gopkg.in/go-playground/validator.v9"
	mgo "gopkg.in/mgo.v2"
)

// Global variables for the application
var validate = validator.New()
var collection = InitDatabase()

// InitDatabase start the database and create the collection
func InitDatabase() *mgo.Collection {
	session, _ := mgo.Dial("mongodb://localhost")
	collection := session.DB("pedidosya").C("comments")
	index := mgo.Index{
		Key:    []string{"purchase"},
		Unique: true,
	}
	_ = collection.EnsureIndex(index)
	return collection
}
