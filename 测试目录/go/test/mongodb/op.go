package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name  string
	Age   int
	Email string
}

// 添加一个用户
func AddUser(user User) {
	client := Connect().(*mongo.Client)
	// 向test数据库 新建一个danmu表
	collection := client.Database("test").Collection("danmu")
	// 添加一条数据
	doc := bson.M{
		"name":  user.Name,
		"age":   user.Age,
		"email": user.Email,
	}
	insertResult, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
