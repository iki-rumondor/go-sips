package migrate

import "github.com/iki-rumondor/sips/internal/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.Pengguna{}},
		{Model: models.Mahasiswa{}},
		{Model: models.PembimbingAkademik{}},
		{Model: models.Role{}},
		{Model: models.Pengaturan{}},
	}
}
