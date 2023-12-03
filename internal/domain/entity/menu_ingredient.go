package entity

type MenuIngredient struct {
	ID           int64       `json:"id"`
	IngredientID int64       `json:"-"`
	Ingredient   *Ingredient `json:"ingredient,omitempty"`
	MenuID       int64       `json:"-"`
	Menu         *Menu       `json:"menu,omitempty"`
	Quantity     float64     `json:"quantity"`
}
