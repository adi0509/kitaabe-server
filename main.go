package main

import (
	"kitaabe2/apis/controller"
	"kitaabe2/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	_, _, _, err := mongo.Connect("mongodb+srv://pustak:Asdfghjkl@kitaabe.gvwt3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	database := "kitaabe"
	controller.CreateUserIndex(database)
	controller.CreateItemIndex(database)
	controller.CreateCategoryIndex(database)
	controller.CreateSubcategoryIndex(database)
	controller.CreateMediaIndex(database)

	r := gin.Default()

	//********USER ROUTES********
	//insert new user
	r.POST("/api/user", controller.CreateUser)
	//get the particular user data
	r.GET("/api/user/:id", controller.GetUser)
	//login
	r.POST("/api/user/login", controller.Login)
	//update user
	r.POST("/api/user/update", controller.UpdateUser)

	//********ITEM ROUTES********
	//insert new item
	r.POST("/api/item", controller.CreateItem)
	//get the particular item data with id
	r.GET("/api/item/id/:id", controller.GetItemById)
	//get the particular item data with name
	r.GET("/api/item/name/:name", controller.GetItemByName)
	//update item
	r.POST("/api/item/update", controller.UpdateItem)

	//********CATEGORY ROUTES********
	//insert new category
	r.POST("/api/category", controller.CreateCategory)
	//get the particular category data with id
	r.GET("/api/category/id/:id", controller.GetCategoryWithId)
	//get the particular category data with name
	r.GET("/api/category/name/:name", controller.GetCategoryWithName)
	//update item
	r.POST("/api/category/update", controller.UpdateCategory)

	//********SUBCATEGORY ROUTES********
	//insert new item
	r.POST("/api/subcategory", controller.CreateSubcategory)
	//get the particular subcategory data with id
	r.GET("/api/subcategory/id/:id", controller.GetSubcategoryWithId)
	//get the particular subcategory data with name
	r.GET("/api/subcategory/name/:name", controller.GetSubcategoryWithName)
	//update item
	r.POST("/api/subcategory/update", controller.UpdateSubcategory)

	//********Media ROUTES********
	//insert new item
	r.POST("/api/media", controller.CreateMedia)
	//get the particular subcategory data with id
	r.GET("/api/media/mediaid/:id", controller.GetMediaWithMediaId)
	//get the particular subcategory data with name
	r.GET("/api/media/itemid/:id", controller.GetMediaWithItemId)
	//update item
	r.POST("/api/media/update", controller.UpdateMedia)

	r.Run()
}
