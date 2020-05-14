package contracts

type Filter struct {
	Sort string // [{"order": "asc"}]
	Limit int32 // 1
	Search string // Partial search inside Item.Content
}