package types

type CreateReminderType struct {
	Message    *string `json:"message,omitempty"`
	Tune       *string `json:"tune,omitempty"`
	Time       *string `json:"time,omitempty"`
	SenderId   *string `json:"senderId,omitempty"`
	ReceiverId *string `json:"receiverId,omitempty"`
	CreatedAt  *string `json:"createdAt,omitempty"`
	UpdatedAt  *string `json:"updatedAt,omitempty"`
}
type UpdateReminderType struct {
	Message *string `json:"message,omitempty"`
	Tune    *string `json:"tune,omitempty"`
	Time    *string `json:"time,omitempty"`
}
