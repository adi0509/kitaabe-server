package controller

import (
	"kitaabe2/apis/model"
	"kitaabe2/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mediaDatabase string
var mediaCollection string
var mediaPrimaryKey string

//create index in the database
func CreateMediaIndex(dbName string) {
	mediaDatabase = dbName
	mediaCollection = "media"
	mediaPrimaryKey = "_id"
}

// Get single media
func GetMediaWithMediaId(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var filter, option interface{}
	filter = bson.D{
		{mediaPrimaryKey, id},
	}

	option = bson.D{}
	cursor, err := mongo.Query(mediaDatabase, mediaCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var medias []primitive.M
	for cursor.Next(mongo.Context) {
		var media bson.M
		err := cursor.Decode(&media)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		medias = append(medias, media)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": medias})
}

// Get single media
func GetMediaWithItemId(ctx *gin.Context) {
	param := ctx.Param("id")

	var filter, option interface{}
	filter = bson.D{
		{"item_id", param},
	}

	option = bson.D{}
	cursor, err := mongo.Query(mediaDatabase, mediaCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var categories []primitive.M
	for cursor.Next(mongo.Context) {
		var media bson.M
		err := cursor.Decode(&media)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, media)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categories})
}

// POST /user
// Create a new user
func CreateMedia(c *gin.Context) {
	var input model.Media
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create media
	media := model.Media{Url: input.Url, Item_id: input.Item_id}

	//Insert it into mongoDB
	_, err := mongo.InsertOne(mediaDatabase, mediaCollection, media)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": media})
}

// Update a user
func UpdateMedia(c *gin.Context) {
	var input model.UpdateMediaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := primitive.ObjectIDFromHex(input.Media_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	filter := bson.D{
		{mediaPrimaryKey, id},
	}

	update := bson.D{
		{"$set", bson.D{
			{"url", input.Url},
			{"item_id", input.Item_id},
		}},
	}

	// Returns result of updated document and a error.
	_, err = mongo.UpdateOne(mediaDatabase, mediaCollection, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}
