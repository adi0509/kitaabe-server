package model

import "time"

//user struct (Model)
type Item struct {
	// Item_id           string `json:"_id"`
	Item_name         string `json:"item_name"`
	Item_description  string `json:"item_description"`
	Price             string `json:"price"`
	Seller_id         string `json:"seller_id"`
	Available_in_city string `json:"available_in_city"`
	Category_id       string `json:"category_id"`
	Subcategory_id    string `json:"subcategory_id"`
	Status            string `json:"status"`
	University        string `json:"university"`
	Listed_on         string `json:"listed_on"`
	Created_at        string `json:"created_at"`
	Updated_at        string `json:"updated_at"`
}

type CreateItemInput struct {
	Item_name         string `json:"item_name"`
	Item_description  string `json:"item_description"`
	Price             string `json:"price"`
	Seller_id         string `json:"seller_id"`
	Available_in_city string `json:"available_in_city"`
	Category_id       string `json:"category_id"`
	Subcategory_id    string `json:"subcategory_id"`
	Status            string `json:"status"`
	University        string `json:"university"`
}

type UpdateItemInput struct {
	Item_id           string `json:"_id"`
	Item_name         string `json:"item_name"`
	Item_description  string `json:"item_description"`
	Price             string `json:"price"`
	Seller_id         string `json:"seller_id"`
	Available_in_city string `json:"available_in_city"`
	Category_id       string `json:"category_id"`
	Subcategory_id    string `json:"subcategory_id"`
	Status            string `json:"status"`
	University        string `json:"university"`
}

type FilterItemInput struct {
	Category_id    string `json:"category_id"`
	Subcategory_id string `json:"subcategory_id"`
	Search         string `json:"search"`
	University     string `json:"university"`
}

func (item *Item) BeforeCreate() (err error) {
	if err != nil {
		return err
	}
	item.Listed_on = time.Now().Local().String()
	item.Created_at = time.Now().Local().String()
	item.Updated_at = time.Now().Local().String()
	return nil
}
