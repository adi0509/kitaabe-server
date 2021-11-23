package model

import "time"

//category struct (Model)
type Category struct {
	// Item_id           string `json:"_id"`
	Category_name string `json:"category_name"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
}

type CreateCategoryInput struct {
	Category_name string `json:"category_name"`
}

type UpdateCategoryInput struct {
	Category_id   string `json:"_id"`
	Category_name string `json:"category_name"`
}

func (category *Category) BeforeCreate() (err error) {
	category.Created_at = time.Now().Local().String()
	category.Updated_at = time.Now().Local().String()
	return nil
}
