package request

type Teacher struct {
	Name string `json:"name" valid:"required~field name tidak ditemukan"`
}
