package database

import (
	"context"
	"log"
	"nubank/model"
)

func InsertJob(job model.Job) (err error) {

	// select users collection
	c := Db.Collection("jobs")

	// insert user into users collection
	res, err := c.InsertOne(context.TODO(), job)

	// checks if any error occurs insert an user
	if err != nil {
		log.Printf("[ERROR] probleming inserting user: %v %v", err, res.InsertedID)
		return
	}

	return
}

// func SearchUser(l model.LoginUser) (u model.User, err error) {

// 	// select users collection
// 	c := Db.Collection("users")

// 	// create filter
// 	filter := bson.D{{"email", l.Email}, {"pass", l.Pass}}

// 	// try to find user with login
// 	err = c.FindOne(context.TODO(), filter).Decode(&u)

// 	// checks if any error occurs insert an user
// 	if err != nil {
// 		log.Printf("[ERROR] probleming searching user: %v %v", err, u)
// 		return
// 	}

// 	log.Printf("[INFO] searching user: %v %v", err, u)

// 	return
// }
