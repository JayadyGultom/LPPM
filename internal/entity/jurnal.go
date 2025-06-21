package entity

type Jurnal struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
	Jenis   string `json:"jenis" gorm:"-"` // For identifying which jurnal table
}

// Specific Jurnal entities
type JurnalJk struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type JurnalJs struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type JurnalKiat struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type JurnalTajb struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type JurnalTeknois struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

type JurnalTmjb struct {
	ID      int    `json:"id" gorm:"primaryKey;column:id"`
	Judul   string `json:"judul" gorm:"column:judul;type:varchar(30)"`
	Konten  string `json:"konten" gorm:"column:konten;type:text"`
}

func (JurnalJk) TableName() string {
	return "lppm_jurnal_jk"
}

func (JurnalJs) TableName() string {
	return "lppm_jurnal_js"
}

func (JurnalKiat) TableName() string {
	return "lppm_jurnal_kiat"
}

func (JurnalTajb) TableName() string {
	return "lppm_jurnal_tajb"
}

func (JurnalTeknois) TableName() string {
	return "lppm_jurnal_teknois"
}

func (JurnalTmjb) TableName() string {
	return "lppm_jurnal_tmjb"
} 