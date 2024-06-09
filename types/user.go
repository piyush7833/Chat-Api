package types

type UpdateUserType struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Image *string `json:"image,omitempty"`
}

type UserType struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Username  *string `json:"username,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
	Image     *string `json:"image,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}
