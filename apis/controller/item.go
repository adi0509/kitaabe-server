package controller

import (
	"context"
	"kitaabe2/apis/model"
	"kitaabe2/mongo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var itemDatabase string
var itemCollection string
var itemPrimaryKey string

//create index in the database
func CreateItemIndex(dbName string) {
	itemDatabase = dbName
	itemCollection = "item"
	itemPrimaryKey = "_id"
	if err := mongo.AddTextIndexItem(itemDatabase, itemCollection); err != nil {
		return
	}
}

// Get single users
func GetItemById(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var filter, option interface{}
	filter = bson.D{
		{itemPrimaryKey, id},
	}

	option = bson.D{}
	cursor, err := mongo.Query(itemDatabase, itemCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var items []primitive.M
	for cursor.Next(mongo.Context) {
		var item bson.M
		err := cursor.Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// Get single users
func GetItemByName(ctx *gin.Context) {
	param := ctx.Param("name")

	var filter, option interface{}
	filter = bson.D{
		{"item_name", param},
	}

	option = bson.D{}
	cursor, err := mongo.Query(itemDatabase, itemCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var items []primitive.M
	for cursor.Next(mongo.Context) {
		var item bson.M
		err := cursor.Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// Get item by seller email
func GetItemBySeller(ctx *gin.Context) {
	param := ctx.Param("email")

	var filter, option interface{}
	filter = bson.D{
		{"seller_id", param},
	}

	option = bson.D{}
	cursor, err := mongo.Query(itemDatabase, itemCollection, filter, option)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var items []primitive.M
	for cursor.Next(mongo.Context) {
		var item bson.M
		err := cursor.Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// POST /user
// Create a new user
func CreateItem(c *gin.Context) {
	var input model.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create item
	item := model.Item{Item_name: input.Item_name, Item_description: input.Item_description, Promoted: "", Price: input.Price, Seller_id: input.Seller_id, Contact_info: input.Contact_info, Available_in_city: input.Available_in_city, Category_id: input.Category_id, Subcategory_id: input.Subcategory_id, Status: input.Status, University: input.University, Listed_on: "", Created_at: "", Updated_at: ""}

	item.BeforeCreate()

	//Insert it into mongoDB
	result, err := mongo.InsertOne(itemDatabase, itemCollection, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_id := result.InsertedID
	c.JSON(http.StatusOK, gin.H{"status": item, "_id": _id})
}

// POST /api/item/all
// get all items with filter
func GetItemByFilter(ctx *gin.Context) {
	var input model.FilterItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var filter interface{}
	if input.Search == "" {
		filter = bson.D{
			{"subcategory_id", bson.M{"$regex": input.Subcategory_id, "$options": "i"}},
			{"category_id", bson.M{"$regex": input.Category_id, "$options": "i"}},
			{"status", "1"},
		}
	} else {
		filter = bson.M{
			"$and": []bson.M{
				bson.M{"$text": bson.M{"$search": input.Search}},
				bson.M{"category_id": bson.M{"$regex": input.Category_id, "$options": "i"}},
				bson.M{"subcategory_id": bson.M{"$regex": input.Subcategory_id, "$options": "i"}},
				bson.M{"status": "1"},
			},
		}
	}

	collection := mongo.Client.Database(itemDatabase).Collection(itemCollection)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var items []primitive.M
	for cursor.Next(mongo.Context) {
		var item bson.M
		err := cursor.Decode(&item)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// Update a user
func UpdateItem(c *gin.Context) {
	var input model.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := primitive.ObjectIDFromHex(input.Item_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	filter := bson.D{
		{itemPrimaryKey, id},
	}

	update := bson.D{
		{"$set", bson.D{
			{"item_name", input.Item_name},
			{"item_description", input.Item_description},
			{"price", input.Price},
			{"seller_id", input.Seller_id},
			{"contact_info", input.Contact_info},
			{"available_in_city", input.Available_in_city},
			{"category_id", input.Category_id},
			{"subcategory_id", input.Subcategory_id},
			{"status", input.Status},
			{"university", input.University},
			{"updated_at", time.Now().Local().String()},
		}},
	}

	// Returns result of updated document and a error.
	_, err = mongo.UpdateOne(itemDatabase, itemCollection, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}
