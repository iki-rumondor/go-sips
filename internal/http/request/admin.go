package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type Pembimbing struct {
	UserUuid string `json:"user_uuid" valid:"required~field user uuid tidak ditemukan"`
	Nama     string `json:"nama" valid:"required~field nama tidak ditemukan"`
	Nip      string `json:"nip" valid:"required~field nip tidak ditemukan"`
}

type UpdatePembimbing struct {
	Nama     string `json:"nama" valid:"required~field nama tidak ditemukan"`
	Nip      string `json:"nip" valid:"required~field nip tidak ditemukan"`
}

type Prodi struct {
	Name     string `json:"name" valid:"required~field nama tidak ditemukan"`
	Kaprodi  string `json:"kaprodi" valid:"required~field kaprodi tidak ditemukan"`
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
}

type KelasRule struct {
	JumlahMahasiswa string `json:"amount" valid:"required~field Jumlah Mahasiswa tidak ditemukan, int~field Jumlah Mahasiswa harus berupa bilangan bulat, range(1|200)~field Jumlah Mahasiswa tidak valid"`
}

type PercepatanCond struct {
	TotalSks    string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk         string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	JumlahError string `json:"jumlah_error" valid:"required~field jumlah_error tidak ditemukan, int~field jumlah_error harus berupa bilangan bulat, range(0|200)~field jumlah_error tidak valid"`
}

type Pengaturan struct {
	TotalSks           string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk                string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	JumlahError        string `json:"jumlah_error" valid:"required~field jumlah_error tidak ditemukan, int~field jumlah_error harus berupa bilangan bulat, range(0|200)~field jumlah_error tidak valid"`
	AngkatanPercepatan string `json:"angkatan_percepatan" valid:"required~field angkatan percepatan tidak ditemukan, int~field angkatan percepatan tidak valid, range(2000|3000)~field angkatan percepatan tidak valid"`
	AngkatanKelas      string `json:"angkatan_kelas" valid:"required~field angkatan kelas tidak ditemukan, int~field angkatan kelas tidak valid, range(2000|3000)~field angkatan kelas tidak valid"`
	JumlahMahasiswa    string `json:"jumlah_mahasiswa" valid:"required~field jumlah mahasiswa per kelas tidak ditemukan, int~field Jumlah Mahasiswa harus berupa bilangan bulat, range(1|200)~field Jumlah Mahasiswa tidak valid"`
}
