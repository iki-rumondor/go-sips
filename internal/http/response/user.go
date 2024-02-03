package response

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type SubjectsCount struct {
	General   int `json:"general"`
	Practical int `json:"practical"`
}
