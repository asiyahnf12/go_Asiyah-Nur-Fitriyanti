package payload

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}
