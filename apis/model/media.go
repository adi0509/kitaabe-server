package model

//category struct (Model)
type Media struct {
	Url     string `json:"url"`
	Item_id string `json:"item_id"`
}

type UpdateMediaInput struct {
	Media_id string `json:"_id"`
	Url      string `json:"url"`
	Item_id  string `json:"item_id"`
}
