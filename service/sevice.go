package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/GURUAKASH-MUTHURAJAN/swagger/config"
	"github.com/GURUAKASH-MUTHURAJAN/swagger/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ListUser() []model.User {
	filter := bson.D{}
	cursor, err := config.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []model.User
	for cursor.Next(context.Background()) {
		var profile model.User
		err := cursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles

}
func GetUser(Name string) (*model.User, error) {
	for _, user := range model.Users {
		if user.Name == Name {
			return &user, nil
		}
	}
	return nil, sql.ErrNoRows
}

func CreateUser(user model.User) error {
	inserted, err := config.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted", inserted.InsertedID)
	return nil
}
