package handler

import (
	"fmt"
	"golang-campaign/user"
	"golang-campaign/user/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		// var errors []string
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errors = append(errors, e.Error())
		// }

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newUser, "jwtToken")
	response_helper := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response_helper)

}

func (h *userHandler) Login(c *gin.Context) {
	/*
		- user memasukan input ( email dan password)
		- input ditangkap handler
		- mapping dari input user ke input struct
		- input struct passing servcie
		- di service mencari dengan bantuan repository user dgn email x......
		- mencocokan password
	*/
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors: ": errors}

		response := helper.APIResponse("Login Failed x", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err2 := h.userService.Login(input)

	fmt.Println(loggedinUser)
	fmt.Println(err2)

	if err2 != nil {
		errorMessage2 := gin.H{"errors": err2.Error()}

		response2 := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage2)
		c.JSON(http.StatusUnprocessableEntity, response2)
		return
	}

	formatter2 := user.FormatUser(loggedinUser, "jwtToken")
	response_helper2 := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter2)

	c.JSON(http.StatusOK, response_helper2)

}

func (h *userHandler) LoginByName(c *gin.Context) {
	var inputName user.LoginByName

	err := c.ShouldBindJSON(&inputName)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors: ": errors}

		response := helper.APIResponse("Login Failed x", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	varinputName, err2 := h.userService.LoginName(inputName)

	fmt.Println("varinputName : ", varinputName)
	fmt.Println("err2 login: ", err2)
	/*
		if err2 != nil {
			errorMessage2 := gin.H{"errors": err2.Error()}

			response2 := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage2)
			c.JSON(http.StatusUnprocessableEntity, response2)
			return
		}

		formatter2 := user.FormatUser(loggedinUser, "jwtToken")
		response_helper2 := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter2)

		c.JSON(http.StatusOK, response_helper2)
	*/
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response_helper := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response_helper)
		return
	}

	isEmailAvailability, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response_helper := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response_helper)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailability,
	}

	var metaMessage string
	if isEmailAvailability {
		metaMessage = "Email is Available"
	} else {
		metaMessage = "Email has been registered"
	}

	response_helper := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusUnprocessableEntity, response_helper)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response_helper := helper.APIResponse("Failed to upload avatar image 1", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response_helper)
		return
	}

	pathimages := "images/" + file.Filename

	fmt.Println("pathimages: ", pathimages)

	err = c.SaveUploadedFile(file, pathimages)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response_helper := helper.APIResponse("Failed to upload avatar image 2", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response_helper)
		return
	}

	userID := 1
	_, err = h.userService.SaveAvatar(userID, pathimages)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response_helper := helper.APIResponse("Failed to upload avatar image 3", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response_helper)
		return
	}

	data := gin.H{"is_uploaded": true}
	response_helper := helper.APIResponse("avatar succesfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response_helper)

}
