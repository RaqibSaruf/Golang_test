package main

import (
	"fmt"

	"github.com/RaqibSaruf/learnPackage/auth"
	"github.com/RaqibSaruf/learnPackage/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("learnPackage", "learn")

	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "learnPackage@example.com",
		Name:  "Learn Package User",
	}

	fmt.Println("user", user)

	color.Green("name", user.Name)
}
