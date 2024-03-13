package migrate

import "github.com/iki-rumondor/sips/internal/models"

type Model struct {
	Model interface{}
}

func GetAllModels() []Model {
	return []Model{
		{Model: models.Pengguna{}},
		{Model: models.Mahasiswa{}},
		{Model: models.TahunAjaran{}},
		{Model: models.Percepatan{}},
		{Model: models.Peringatan{}},
		{Model: models.Kelas{}},
		{Model: models.PembimbingAkademik{}},
		{Model: models.Role{}},
	}
}
