package types

type Student struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Email string `json:"email" validate:"required,email,unique=students;email"`
}
