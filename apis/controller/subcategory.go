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

var subcategoryDatabase string
var subcategoryCollection string
var subcategoryPrimaryKey string

//create index in the database
func CreateSubcategoryIndex(dbName string) {
	subcategoryDatabase = dbName
	subcategoryCollection = "subcategory"
	subcategoryPrimaryKey = "_id"
}

// Get single subcategory
func GetSubcategoryWithId(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}
	var filter, option interface{}
	filter = bson.D{
		{subcategoryPrimaryKey, id},
	}

	option = bson.D{}
	cursor, err := mongo.Query(subcategoryDatabase, subcategoryCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var subcategories []primitive.M
	for cursor.Next(mongo.Context) {
		var subcategory bson.M
		err := cursor.Decode(&subcategory)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		subcategories = append(subcategories, subcategory)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": subcategories})
}

// Get single subcategory
func GetSubcategoryWithName(ctx *gin.Context) {
	param := ctx.Param("name")

	var filter, option interface{}
	filter = bson.D{
		{"subcategory_name", param},
	}

	option = bson.D{}
	cursor, err := mongo.Query(subcategoryDatabase, subcategoryCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var subcategories []primitive.M
	for cursor.Next(mongo.Context) {
		var subcategory bson.M
		err := cursor.Decode(&subcategory)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		subcategories = append(subcategories, subcategory)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": subcategories})
}

// POST /user
// Create a new user
func CreateSubcategory(c *gin.Context) {
	var input model.CreateSubcategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create subcategory
	subcategory := model.Subcategory{Subcategory_name: input.Subcategory_name, Category_id: input.Category_id, Created_at: "", Updated_at: ""}

	subcategory.BeforeCreate()

	//Insert it into mongoDB
	_, err := mongo.InsertOne(subcategoryDatabase, subcategoryCollection, subcategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": subcategory})
}

// Update a user
func UpdateSubcategory(c *gin.Context) {
	var input model.UpdateSubcategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := primitive.ObjectIDFromHex(input.Subcategory_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	filter := bson.D{
		{subcategoryPrimaryKey, id},
	}

	update := bson.D{
		{"$set", bson.D{
			{"subcategory_name", input.Subcategory_name},
			{"category_id", input.Category_id},
			{"updated_at", time.Now().Local().String()},
		}},
	}

	// Returns result of updated document and a error.
	_, err = mongo.UpdateOne(subcategoryDatabase, subcategoryCollection, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}
