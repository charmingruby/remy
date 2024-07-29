package common

func NewPayloadErr(reason string) error {
	return &ErrPayload{
		Message: reason,
	}
}

type ErrPayload struct {
	Message string `json:"message"`
}

func (e *ErrPayload) Error() string {
	return e.Message
}

func NewNoItemsErr() error {
	return NewPayloadErr("items must have at least one item")
}
