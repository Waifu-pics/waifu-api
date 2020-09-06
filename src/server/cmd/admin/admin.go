package admin

import (
	"log"
	"os"

	"github.com/Riku32/waifu.pics/src/database"
	"github.com/alexedwards/argon2id"
	"github.com/tcnksm/go-input"
)

func CreateAdmin(database database.Database) {
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

	if username == "" || password == "" {
		log.Fatal("Invalid username or password!")
	}

	passwordHashed, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal("Unable to hash password!")
	}

	err = database.CreateAdmin(username, passwordHashed)
	if err != nil {
		log.Fatal("Unable to create admin!")
	}

	log.Println("Admin has been created!")
}
