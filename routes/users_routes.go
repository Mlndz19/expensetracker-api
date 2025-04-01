package routes

import (
	"fmt"
	"net/http"

	"expensetrack/main.go/config"
	"expensetrack/main.go/models"
	utils "expensetrack/main.go/utils/crypto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterUsersRoutes(r *gin.Engine) {
	r.GET("api/user", getAllUsers)
	r.GET("api/user/:id", getUserById)
	r.POST("api/user", createUser)
	r.PUT("api/user/:id", updateUserById)
	r.DELETE("api/user/:id", deleteUserById)
}

func getAllUsers(c *gin.Context) {
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	var users []models.User

	fmt.Println("Test")
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getUserById(c *gin.Context){
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func createUser(c *gin.Context){
	var newUSer models.User

	if err := c.ShouldBindJSON(&newUSer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(newUSer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUSer.Guid = uuid.NewString()
	newUSer.Password = hashedPassword

	if err := config.DB.Create(&newUSer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created correctly",
		"user": newUSer,
	})
}

func updateUserById(c *gin.Context){
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated correctly",
		"user": user,
	})
}

func deleteUserById(c *gin.Context){
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted correctly"})
}

/* func getPersonById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var userToFind *models.User
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
	var newUser *models.User
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

	var userToUpdate *models.User
	var userBody *models.User

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

	var userToDelete *models.User
	for i := range users {
		if users[i].Id == id {
			userToDelete = &users[i]
		}
	}

	if userToDelete == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var newUsers []models.User
	for _, user := range users {
		if user.Id != userToDelete.Id {
			newUsers = append(newUsers, user)
		}
	}

	users = newUsers

	c.JSON(http.StatusOK, gin.H{"message": "User deleted correctly"})
} */