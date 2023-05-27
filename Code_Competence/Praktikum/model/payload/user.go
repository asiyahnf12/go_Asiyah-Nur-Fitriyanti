package payload

type (
	LoginUserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	Item struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Stock       uint   `json:"stock" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
		CategoryID  uint   `json:"category_id" validate:"required"`
	}

	CreateUserRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	CreateCategory struct {
		Name string `json:"name" form:"name" validate:"required"`
	}

	CategoryRequest struct {
		Name string `json:"name" validate:"required"`
	}

	CategoryResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name" validate:"required"`
	}
)
