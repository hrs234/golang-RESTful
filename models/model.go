package models

type Kategori struct {
	Id       string `form:"id" json:"id"`
	Kategori string `form:"kategori" json:"kategori"`
}

type Item struct {
	Id          string `form:"id" json:"id"`
	Item_id     string `form:"item_id" json:"item_id"`
	Kategori_id string `form:"kategori_id" json:"kategori_id"`
}

type Sign struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type ResponseKategori struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Kategori
}

type ResponseItem struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Item
}
