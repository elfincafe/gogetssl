package gogetssl

type (
	Error struct {
		Error       bool   `json:"error"`
		Message     string `json:"message"`
		Description string `json:"description"`
	}
)
