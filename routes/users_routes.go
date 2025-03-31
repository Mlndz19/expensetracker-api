package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Id       int       `json:"id"`
	Guid     string    `json:"guid"`
	Created  time.Time `json:"created"`
	Email    string    `json:"email" binding:"required,email"`
	Username string    `json:"username" binding:"required,min=5,max=25"`
	Password string    `json:"password" binding:"required"`
	Active   bool      `json:"active"`
}

var users = []User{
	{Id: 1, Guid: "550e8400-e29b-41d4-a716-446655440000", Created: time.Now(), Email: "user1@example.com", Username: "user1", Password: "pass1", Active: true},
	{Id: 2, Guid: "550e8400-e29b-41d4-a716-446655440001", Created: time.Now(), Email: "user2@example.com", Username: "user2", Password: "pass2", Active: false},
	{Id: 3, Guid: "550e8400-e29b-41d4-a716-446655440002", Created: time.Now(), Email: "user3@example.com", Username: "user3", Password: "pass3", Active: true},
	{Id: 4, Guid: "550e8400-e29b-41d4-a716-446655440003", Created: time.Now(), Email: "user4@example.com", Username: "user4", Password: "pass4", Active: false},
	{Id: 5, Guid: "550e8400-e29b-41d4-a716-446655440004", Created: time.Now(), Email: "user5@example.com", Username: "user5", Password: "pass5", Active: true},
}

func RegisterUsersRoutes(r *gin.Engine) {
	r.GET("api/user", getAllUsers)
	r.GET("api/user/:id", getPersonById)
	r.POST("api/user", createUser)
	r.PUT("api/user/:id", updateUser)
	r.DELETE("api/user/:id", deleteUser)
}

func getAllUsers(c *gin.Context){
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"users": []User{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getPersonById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var userToFind *User
	for i := range users {
		if users[i].Id == id {
			userToFind = &users[i]
		}
	}

	if userToFind == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userToFind})
}

func createUser(c *gin.Context){
	var newUser *User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.Id = len(users) + 1
	newUser.Guid = uuid.NewString()
	newUser.Created = time.Now()
	newUser.Active = true

	users = append(users, *newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created correctly",
		"user": newUser,
	})
}

func updateUser(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var userToUpdate *User
	var userBody *User

	for i := range users {
		if users[i].Id == id {
			userToUpdate = &users[i]
		}
	}

	if userToUpdate == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userToUpdate.Email = userBody.Email
	userToUpdate.Username = userBody.Username
	userToUpdate.Password = userBody.Password
	userToUpdate.Active = userBody.Active

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated correctly",
		"user": userToUpdate,
	})
}

func deleteUser(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var userToDelete *User
	for i := range users {
		if users[i].Id == id {
			userToDelete = &users[i]
		}
	}

	if userToDelete == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var newUsers []User
	for _, user := range users {
		if user.Id != userToDelete.Id {
			newUsers = append(newUsers, user)
		}
	}

	users = newUsers

	c.JSON(http.StatusOK, gin.H{"message": "User deleted correctly"})
}