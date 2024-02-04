package services

import (
	"errors"
	"log"
	"math"
	"strconv"

	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type MahasiswaService struct {
	Repo interfaces.MahasiswaRepoInterface
}

func NewMahasiswaService(repo interfaces.MahasiswaRepoInterface) interfaces.MahasiswaServiceInterface {
	return &MahasiswaService{
		Repo: repo,
	}
}

func (s *MahasiswaService) ImportMahasiswa(pathFile string) (*[]response.FailedImport, error) {
	f, err := excelize.OpenFile(pathFile)
	if err != nil {
		log.Println("Gagal Membuka File")
		return nil, response.SERVICE_INTERR
	}
	defer f.Close()

	rows, err := f.GetRows("Mahasiswa")
	if err != nil {
		log.Println("Failed to get rows Mahasiswa")
		return nil, response.SERVICE_INTERR
	}

	var failedImport []response.FailedImport

	for i := 1; i < len(rows); i++ {
		cols := rows[i]
		angkatan, err := strconv.Atoi(cols[2])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Angkatan Bukan Angka",
			})
			continue
		}

		totalSks, err := strconv.Atoi(cols[3])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Total SKS Bukan Angka",
			})
			continue
		}

		ipk, err := strconv.ParseFloat(cols[4], 32)
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "IPK Bukan Angka",
			})
			continue
		}

		jumlahError, err := strconv.Atoi(cols[5])
		if err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Jumlah Error Bukan Angka",
			})
			continue
		}

		mahasiswa := models.Mahasiswa{
			Nim:         cols[0],
			Nama:        cols[1],
			Angkatan:    uint(angkatan),
			TotalSks:    uint(totalSks),
			Ipk:         math.Round(ipk*100) / 100,
			JumlahError: uint(jumlahError),
		}

		if err := s.Repo.CreateMahasiswa(&mahasiswa); err != nil {
			failedImport = append(failedImport, response.FailedImport{
				Nim:   cols[0],
				Nama:  cols[1],
				Pesan: "Gagal Menambahkan Mahasiswa: " + err.Error(),
			})
			continue
		}
	}

	return &failedImport, nil
}

func (s *MahasiswaService) GetAllMahasiswa() (*[]models.Mahasiswa, error) {
	result, err := s.Repo.FindAllMahasiswa()
	if err != nil {
		log.Println(err.Error())
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MahasiswaService) GetMahasiswa(uuid string) (*models.Mahasiswa, error) {
	result, err := s.Repo.FindMahasiswaByUuid(uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Mahasiswa Tidak Ditemukan")
		}
		return nil, response.SERVICE_INTERR
	}

	return result, nil
}

func (s *MahasiswaService) UpdateMahasiswa(uuid string, req *request.Mahasiswa) error {
	result, err := s.GetMahasiswa(uuid)
	if err != nil {
		return err
	}

	angkatan, _ := strconv.Atoi(req.Angkatan)
	totalSks, _ := strconv.Atoi(req.TotalSks)
	ipk, _ := strconv.ParseFloat(req.Ipk, 64)
	jumlahError, _ := strconv.Atoi(req.JumlahError)

	model := models.Mahasiswa{
		ID:          result.ID,
		Nim:         req.Nim,
		Nama:        req.Nama,
		Angkatan:    uint(angkatan),
		TotalSks:    uint(totalSks),
		Ipk:         ipk,
		JumlahError: uint(jumlahError),
	}

	if err := s.Repo.UpdateMahasiswa(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *MahasiswaService) DeleteMahasiswa(uuid string) error {
	result, err := s.GetMahasiswa(uuid)
	if err != nil {
		return err
	}

	model := models.Mahasiswa{
		ID: result.ID,
	}

	if err := s.Repo.DeleteMahasiswa(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return response.VIOLATED_ERR
		}
		return response.SERVICE_INTERR
	}

	return nil
}
