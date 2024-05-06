package testing

import (
	"errors"
	"testing"

	"github.com/iki-rumondor/sips/internal/mocks"
	"github.com/iki-rumondor/sips/internal/models"
	"github.com/iki-rumondor/sips/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSinkronPercepatan_Success(t *testing.T) {
	mockRepo := &mocks.MahasiswaRepoInterface{}
	var mahasiswa = []models.Mahasiswa{
		{
			ID:                   1,
			PembimbingAkademikID: 1,
			Uuid:                 mock.Anything,
			Nim:                  mock.Anything,
			Nama:                 mock.Anything,
			Angkatan:             2000,
			Ipk:                  1.3,
			JumlahError:          0,
			PenggunaID:           1,
			TotalSks:             123,
			CreatedAt:            1,
			UpdatedAt:            1,
		},
	}

	mockRepo.On("FindAllMahasiswa", "").Return(&mahasiswa, nil)
	mockRepo.On("UpdatePercepatan").Return(nil)

	mahasiswaService := services.NewMahasiswaService(mockRepo)

	err := mahasiswaService.SinkronPercepatan()
	assert.NoError(t, err)
}

func TestSinkronPercepatan_MahasiswaKosong(t *testing.T) {
	mockRepo := &mocks.MahasiswaRepoInterface{}
	var mahasiswa = []models.Mahasiswa{}

	mockRepo.On("FindAllMahasiswa", "").Return(&mahasiswa, nil)
	mockRepo.On("UpdatePercepatan").Return(nil)

	mahasiswaService := services.NewMahasiswaService(mockRepo)

	err := mahasiswaService.SinkronPercepatan()
	assert.Error(t, err)
	assert.Equal(t, "404: Mahasiswa Masih Kosong", err.Error())
}

func TestSinkronPercepatan_FailFindAllMahasiswa(t *testing.T) {
	mockRepo := &mocks.MahasiswaRepoInterface{}
	mockRepo.On("FindAllMahasiswa", "").Return(nil, errors.New(mock.Anything))
	mockRepo.On("UpdatePercepatan").Return(nil)

	mahasiswaService := services.NewMahasiswaService(mockRepo)

	err := mahasiswaService.SinkronPercepatan()
	assert.Error(t, err)
	assert.Equal(t, "500: ServiceError: Terjadi Kesalahan Sistem", err.Error())
}

func TestSinkronPercepatan_FailUpdate(t *testing.T) {
	mockRepo := &mocks.MahasiswaRepoInterface{}
	var mahasiswa = []models.Mahasiswa{
		{
			ID:                   1,
			PembimbingAkademikID: 1,
			Uuid:                 mock.Anything,
			Nim:                  mock.Anything,
			Nama:                 mock.Anything,
			Angkatan:             2000,
			Ipk:                  1.3,
			JumlahError:          0,
			PenggunaID:           1,
			TotalSks:             123,
			CreatedAt:            1,
			UpdatedAt:            1,
		},
	}

	mockRepo.On("FindAllMahasiswa", "").Return(&mahasiswa, nil)
	mockRepo.On("UpdatePercepatan").Return(errors.New(mock.Anything))

	mahasiswaService := services.NewMahasiswaService(mockRepo)

	err := mahasiswaService.SinkronPercepatan()
	assert.Error(t, err)
	assert.Equal(t, "500: ServiceError: Terjadi Kesalahan Sistem", err.Error())
}
