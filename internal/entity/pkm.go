package entity

type PKM struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
	Jenis   string `json:"jenis" gorm:"-"` // For identifying which pkm table
}

// Specific PKM entities
type PKMBame struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMHpp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMLp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMP3 struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMRrp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMSkR struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMStp struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type PKMTcr struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

func (PKMBame) TableName() string {
	return "lppm_pkm_bame"
}

func (PKMHpp) TableName() string {
	return "lppm_pkm_hpp"
}

func (PKMLp) TableName() string {
	return "lppm_pkm_lp"
}

func (PKMP3) TableName() string {
	return "lppm_pkm_p3"
}

func (PKMRrp) TableName() string {
	return "lppm_pkm_rrp"
}

func (PKMSkR) TableName() string {
	return "lppm_pkm_sk_r"
}

func (PKMStp) TableName() string {
	return "lppm_pkm_stp"
}

func (PKMTcr) TableName() string {
	return "lppm_pkm_tcr"
} 