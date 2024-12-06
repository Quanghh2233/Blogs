package dto

type category struct {
	Name string `json:"name" binding:"required,max=100"`
}
