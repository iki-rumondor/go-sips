package response

type Pembimbing struct {
	Uuid string `json:"uuid"`
	Nama string `json:"nama"`
	Nip  string `json:"nip"`
}

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
