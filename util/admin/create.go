package admin

import (
	"context"
	"fmt"
	"github.com/tcnksm/go-input"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"waifu.pics/util/crypto"
)

func CreateAdmin(database *mongo.Database) {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	username, _ := ui.Ask("Enter username", &input.Options{
		Required: true,
		Loop:     true,
	})

	password, _ := ui.Ask("Enter password", &input.Options{
		Required: true,
		Loop:     true,
	})

	passwordHashed, err := crypto.GeneratePassword(password)
	if err != nil {
		log.Fatal("Unable to hash password!")
	}

	_, err = database.Collection("admins").InsertOne(context.TODO(), bson.M{"username": username, "password": passwordHashed, "token": crypto.GenUUID()})
	if err != nil {
		log.Fatal("Unable to create admin!")
	}

	fmt.Println("Admin has been created!")
}
