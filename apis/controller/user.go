package controller

import (
	"fmt"
	"kitaabe2/apis/model"
	"kitaabe2/mongo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userDatabase string
var userCollection string

//create index in the database
func CreateUserIndex(dbName string) {
	userDatabase = dbName
	userCollection = "user"
	mongo.AddIndex(userDatabase, userCollection, "email")
}

// Get single users
func GetUser(ctx *gin.Context) {
	param := ctx.Param("id")
	var filter, option interface{}
	filter = bson.D{
		{"email", param},
	}
	option = bson.D{}
	cursor, err := mongo.Query(userDatabase, userCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var users []primitive.M
	for cursor.Next(mongo.Context) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		users = append(users, user)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /user
// Create a new user
func CreateUser(c *gin.Context) {
	var input model.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := model.User{Name: input.Name, Email: input.Email, Password: input.Password, University: input.University, Created_at: "", Updated_at: ""}
	user.BeforeCreate()

	//Insert it into mongoDB
	_, err := mongo.InsertOne(userDatabase, userCollection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully inserted"})
}

// Update a user
func UpdateUser(c *gin.Context) {
	var input model.UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.D{
		{"email", input.Email},
	}

	update := bson.D{
		{"$set", bson.D{
			{"name", input.Name},
			{"university", input.University},
			{"updated_at", time.Now().Local().String()},
		}},
	}

	// Returns result of updated document and a error.
	_, err := mongo.UpdateOne(userDatabase, userCollection, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}

//Login
func Login(c *gin.Context) {
	var input model.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var filter, option interface{}
	filter = bson.D{
		{"email", input.Email},
	}
	option = bson.D{}

	// get the user details
	cursor, err := mongo.Query(userDatabase, userCollection, filter, option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user primitive.M
	cursor.Next(mongo.Context)
	err = cursor.Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pass := fmt.Sprint(user["password"])

	// check password
	err = model.ComparePassword(pass, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not matched"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "valid user"})
}

// {
// 	"userid": "1",
// 	"name": "1",
// 	"email": "1",
// 	"password": "1",
// 	"university": "1",
// 	"created_at": "1",
// 	"updated_at": "1"
// }
