package psapi

import (
	"time"
)

type Category struct {
	Id              int       `json:"id"`
	IdParent        int       `json:"id_parent"`
	Active          int       `json:"active"`
	IdShopDefault   int       `json:"id_shop_default"`
	IsRootCategory  int       `json:"is_root_category"`
	Position        int       `json:"position"`
	Name            string    `json:"name"`
	LinkRewrite     string    `json:"link_rewrite"`
	Description     string    `json:"description"`
	MetaTitle       string    `json:"meta_title"`
	MetaDescription string    `json:"meta_description"`
	MetaKeywords    string    `json:"meta_keywords"`
	DateAdd         time.Date `json:"date_add"`
	DateUpd         time.Date `json:"date_upd"`
}

type Associations struct {
	categoriesId []int
}

type Product struct {
	Id int `json:"id"`
	Quantity int `json:"quantity"`
	Type string `json:"type"`
	//Description string `json:"description"`
	Associations Associations  `json:"associations"`
}

