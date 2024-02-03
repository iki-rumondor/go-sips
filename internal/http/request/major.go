package request

type Major struct {
	Name string `json:"name" valid:"required~field name tidak ditemukan"`
}
