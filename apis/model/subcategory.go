package model

import "time"

//category struct (Model)
type Subcategory struct {
	// Item_id           string `json:"_id"`
	Subcategory_name string `json:"subcategory_name"`
	Category_id      string `json:"category_id"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
}

type CreateSubcategoryInput struct {
	Subcategory_name string `json:"subcategory_name"`
	Category_id      string `json:"category_id"`
}

type UpdateSubcategoryInput struct {
	Subcategory_id   string `json:"_id"`
	Category_id      string `json:"category_id"`
	Subcategory_name string `json:"subcategory_name"`
}

func (subcategory *Subcategory) BeforeCreate() (err error) {
	subcategory.Created_at = time.Now().Local().String()
	subcategory.Updated_at = time.Now().Local().String()
	return nil
}
