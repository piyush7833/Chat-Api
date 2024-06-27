package types

type CreateTagType struct {
	MessageId  *string `json:"messageId,omitempty"`
	ReminderId *string `json:"reminderId,omitempty"`
	Title      *string `json:"title,omitempty"`
	CreatedAt  *string `json:"createdAt,omitempty"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
}
