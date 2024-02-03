package response

type Teacher struct {
	Uuid       string      `json:"uuid"`
	Name       string      `json:"name"`
	Department *Department `json:"department"`
	CreatedAt  int64       `json:"created_at"`
	UpdatedAt  int64       `json:"updated_at"`
}
