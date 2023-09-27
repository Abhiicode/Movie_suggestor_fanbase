package actions

import (
	"github.com/gobuffalo/buffalo"
	"movie-suggestor-fanbase/models"
)

func LoginHandler(context buffalo.Context) error {
	email := context.Param("email")

	password := context.Param("password")

	response := make(map[string]interface{})

	if(len(password) < 1){
		response["message"] = "Password can not be empty.";
		return context.Render(400, r.JSON(response))
	}

	user, err := UserExists(email)

	if(err != nil){
		response["message"] = "No user was found for that email.";
		return context.Render(404, r.JSON(response))
	}

	if(user.Password == password){
		user.Password = "<Hidden>"
		return context.Render(200, r.JSON(user))
	}

	response["message"] = "Password does not match.";
	return context.Render(400, r.JSON(response))
}

func SignUpHandler (context buffalo.Context) error {
	return context.Render(200,r.JSON(map[string]interface{}{"message":"Not yet implemented."}))
}

func UserExists (email string) (*models.User, error) {
	user := models.User{}

	err := models.DB.Where("email = ?",email).First(&user)

	return &user, err;
}