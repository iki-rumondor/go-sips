package response

type Department struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Head      string `json:"head"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	User      *User  `json:"user"`
	Major     *Major `json:"major"`
}
