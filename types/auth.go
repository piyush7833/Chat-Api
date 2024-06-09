package types

type SignupType struct {
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Image    *string `json:"image,omitempty"`
}

type SignInType struct {
	Email    *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
}
