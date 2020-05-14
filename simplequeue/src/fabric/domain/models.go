package domain

//import "encoding/json"

type Item struct {
	ObjectType string `json:"docType"`
	ID         string `json:"id"`
	Content    string `json:"content"`
	Order      int    `json:"order"`
	//json.RawMessage `json:"data"`
}
