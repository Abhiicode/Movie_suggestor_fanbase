package actions

import (
	"github.com/gobuffalo/buffalo"
	"movie-suggestor-fanbase/models"
	"fmt"
)

func LoginHandler(context buffalo.Context) error {
	email := context.Param("email")
	password := context.Param("password")

	if(len(email) < 1){
		return context.Render(400, r.String("Email is either empty or invalid."))
	}

	if(len(password) < 1){
		return context.Render(400, r.String("Password is either empty or invalid."))
	}

	user := models.User{}

	err := models.DB.Where("email = ?",email).First(&user)

	if(err != nil){
		return context.Render(400, r.String("User not found."))
	}

	if(user.Password == password){
		return context.Render(200, r.String(fmt.Sprintf("Authenticated. [%s]",email)))
	}

	return context.Render(400, r.String("Password does not match."))
}