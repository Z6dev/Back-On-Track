package structs

type TODO struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status bool `json:"status"`
}