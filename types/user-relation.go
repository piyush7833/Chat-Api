package types

type CreateUserRelationType struct {
	UserId        *string `json:"userId,omitempty"`
	RelatedUserId *string `json:"relatedUserId,omitempty"`
	Status        *string `json:"status,omitempty"`
}

type UpdateUserRelationType struct {
	Status *string `json:"status,omitempty"`
}
