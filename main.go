package main

import (
	"fmt"
	"golang-campaign/user"
	"golang-campaign/user/auth"
	"golang-campaign/user/handler"
	"golang-campaign/user/handler/campaigns"
	"golang-campaign/user/helper"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := "root:@tcp(127.0.0.1:3306)/campaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//fmt.Println("SQL Connection established")

	campaignRepository := campaigns.NewRepository(db)
	userRepository := user.NewRepository(db)

	campaigns, err := campaignRepository.FindAll()

	for _, campaigns := range campaigns {
		fmt.Println(campaigns.Name)
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.GGvyPxIF-qtj9SFrDoHXpgIwRmYUKjpqh1BkysASSUk")
	if err != nil {
		// fmt.Println("Error")
		// fmt.Println("Error")
		// fmt.Println("Error")
	}
	if token.Valid {
		// fmt.Println("Valid")
		// fmt.Println("Valid")
		// fmt.Println("Valid")
	} else {
		// fmt.Println("Invalid")
		// fmt.Println("Invalid")
		// fmt.Println("Invalid")
	}

	//userService.SaveAvatar(1, "images/1-profile.png")

	fmt.Println(authService.GenerateToken(1001))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/findName", userHandler.LoginByName)
	api.POST("/email_chekers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleWare(authService, userService), userHandler.UploadAvatar)
	router.Run()

}

func authMiddleWare(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer ") {
			resonse := helper.APIResponse("Unauthenticated", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, resonse)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)

	}

}
