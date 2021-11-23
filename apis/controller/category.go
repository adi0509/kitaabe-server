package controller

import (
	"kitaabe2/apis/model"
	"kitaabe2/mongo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var categoryDatabase string
var categoryCollection string
var categoryPrimaryKey string

//create index in the database
func CreateCategoryIndex(dbName string) {
	categoryDatabase = dbName
	categoryCollection = "category"
	categoryPrimaryKey = "_id"
}

// Get single category
func GetCategoryWithId(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	var filter, option interface{}
	filter = bson.D{
		{categoryPrimaryKey, id},
	}

	option = bson.D{}
	cursor, err := mongo.Query(categoryDatabase, categoryCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var categories []primitive.M
	for cursor.Next(mongo.Context) {
		var category bson.M
		err := cursor.Decode(&category)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		categories = append(categories, category)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// Get single category
func GetCategoryWithName(ctx *gin.Context) {
	param := ctx.Param("name")

	var filter, option interface{}
	filter = bson.D{
		{"category_name", param},
	}

	option = bson.D{}
	cursor, err := mongo.Query(categoryDatabase, categoryCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var categories []primitive.M
	for cursor.Next(mongo.Context) {
		var category bson.M
		err := cursor.Decode(&category)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		categories = append(categories, category)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// POST /user
// Create a new user
func CreateCategory(c *gin.Context) {
	var input model.CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	category := model.Category{Category_name: input.Category_name, Created_at: "", Updated_at: ""}

	category.BeforeCreate()

	//Insert it into mongoDB
	_, err := mongo.InsertOne(categoryDatabase, categoryCollection, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": category})
}

// Update a user
func UpdateCategory(c *gin.Context) {
	var input model.UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := primitive.ObjectIDFromHex(input.Category_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	filter := bson.D{
		{categoryPrimaryKey, id},
	}

	update := bson.D{
		{"$set", bson.D{
			{"category_name", input.Category_name},
			{"updated_at", time.Now().Local().String()},
		}},
	}

	// Returns result of updated document and a error.
	_, err = mongo.UpdateOne(categoryDatabase, categoryCollection, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}

// {
//     "_id": "619516e22a5cec25d461fc6a",
// 	"category_name": "hello update"
// }
