package admin

import (
	"fmt"
	"log"
	"os"

	"github.com/tcnksm/go-input"
	"waifu.pics/util/crypto"
	"waifu.pics/util/database"
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

	passwordHashed, err := crypto.GeneratePassword(password)
	if err != nil {
		log.Fatal("Unable to hash password!")
	}

	err = database.CreateAdmin(username, passwordHashed)
	if err != nil {
		log.Fatal("Unable to create admin!")
	}

	fmt.Println("Admin has been created!")
}
