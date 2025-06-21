package entity

import "time"

type Profil struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Nama        string    `json:"nama"`
    NIDN        string    `json:"nidn"`
    Jabatan     string    `json:"jabatan"`
    Prodi       string    `json:"prodi"`
    Fakultas    string    `json:"fakultas"`
    Email       string    `json:"email"`
    Telepon     string    `json:"telepon"`
    Alamat      string    `json:"alamat"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type VisiMisi struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type StrukturOrganisasi struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

func (Profil) TableName() string {
    return "profil"
}

func (VisiMisi) TableName() string {
	return "lppm_profile_visimisi"
}

func (StrukturOrganisasi) TableName() string {
	return "lppm_profile_so"
}
