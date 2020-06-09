package controllers

// CreateRequest ...
type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Category string `json:"category" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

// type CreateRequest struct {
// 	Name     string `json:"name" binding:"require"`
// 	Author   string `json:"author"`
// 	Category string `json:"category"`
// 	Price    int    `json:"price"`
// }
