package entity

type HKI struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
	Jenis   string `json:"jenis" gorm:"-"` // For identifying which hki table
}

// Specific HKI entities
type HKIDosen struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type HKIMhs struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

func (HKIDosen) TableName() string {
	return "lppm_hki_dosen"
}

func (HKIMhs) TableName() string {
	return "lppm_hki_mhs"
} 