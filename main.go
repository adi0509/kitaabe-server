package main

import (
	"kitaabe2/apis/controller"
	"kitaabe2/mongo"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())

	//********USER ROUTES********
	//insert new user
	r.POST("/api/user", controller.CreateUser)
	//get the particular user
	r.GET("/api/user/:id", controller.GetUser)
	//login
	r.POST("/api/user/login", controller.Login)
	//update user
	r.POST("/api/user/update", controller.UpdateUser)

	//********ITEM ROUTES********
	//insert new item
	r.POST("/api/item", controller.CreateItem)
	//get all items with filter
	r.POST("/api/item/all", controller.GetItemByFilter)
	//get item by id
	r.GET("/api/item/id/:id", controller.GetItemById)
	//get item by name
	r.GET("/api/item/name/:name", controller.GetItemByName)
	//get item by name
	r.GET("/api/item/seller/:email", controller.GetItemBySeller)
	//update item
	r.POST("/api/item/update", controller.UpdateItem)

	//********CATEGORY ROUTES********
	//insert new category
	r.POST("/api/category", controller.CreateCategory)
	//get category by id
	r.GET("/api/category/id/:id", controller.GetCategoryWithId)
	//get category by name
	r.GET("/api/category/name/:name", controller.GetCategoryWithName)
	//update category
	r.POST("/api/category/update", controller.UpdateCategory)

	//********SUBCATEGORY ROUTES********
	//insert new subcategory
	r.POST("/api/subcategory", controller.CreateSubcategory)
	//get subcategory by id
	r.GET("/api/subcategory/id/:id", controller.GetSubcategoryWithId)
	//get subcategory by name
	r.GET("/api/subcategory/name/:name", controller.GetSubcategoryWithName)
	//update subcategory
	r.POST("/api/subcategory/update", controller.UpdateSubcategory)

	//********Media ROUTES********
	//insert new media
	r.POST("/api/media", controller.CreateMedia)
	//get subcategory by id
	r.GET("/api/media/mediaid/:id", controller.GetMediaWithMediaId)
	//get subcategory by name
	r.GET("/api/media/itemid/:id", controller.GetMediaWithItemId)
	//update media
	r.POST("/api/media/update", controller.UpdateMedia)

	r.Run()
}
