package entity

type Penelitian struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
	Jenis   string `json:"jenis" gorm:"-"` // For identifying which penelitian table
}

// Specific penelitian entities
type PenelitianBame struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianHpp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianLp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianP3 struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianRrp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianSkR struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianStp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PenelitianTcr struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

func (PenelitianBame) TableName() string {
	return "lppm_penelitian_bame"
}

func (PenelitianHpp) TableName() string {
	return "lppm_penelitian_hpp"
}

func (PenelitianLp) TableName() string {
	return "lppm_penelitian_lp"
}

func (PenelitianP3) TableName() string {
	return "lppm_penelitian_p3"
}

func (PenelitianRrp) TableName() string {
	return "lppm_penelitian_rrp"
}

func (PenelitianSkR) TableName() string {
	return "lppm_penelitian_sk_r"
}

func (PenelitianStp) TableName() string {
	return "lppm_penelitian_stp"
}

func (PenelitianTcr) TableName() string {
	return "lppm_penelitian_tcr"
} 