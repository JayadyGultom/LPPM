package repository

import (
	"perpustakaan/config"
	"perpustakaan/internal/entity"
)

type ProfilRepository interface {
	GetAll() ([]entity.Profil, error)
	GetByID(id int) (*entity.Profil, error)
	Create(profil entity.Profil) (*entity.Profil, error)
	Update(id int, profil entity.Profil) (*entity.Profil, error)
	Delete(id int) error
}

type VisiMisiRepository interface {
	GetAll() ([]entity.VisiMisi, error)
	GetByID(id int) (*entity.VisiMisi, error)
}

type StrukturOrganisasiRepository interface {
	GetAll() ([]entity.StrukturOrganisasi, error)
	GetByID(id int) (*entity.StrukturOrganisasi, error)
}

type profilRepository struct {
	db *config.Database
}

type visimisiRepository struct {
	db *config.Database
}

type strukturOrganisasiRepository struct {
	db *config.Database
}

func NewProfilRepository(db *config.Database) ProfilRepository {
	return &profilRepository{db: db}
}

func NewVisiMisiRepository(db *config.Database) VisiMisiRepository {
	return &visimisiRepository{db: db}
}

func NewStrukturOrganisasiRepository(db *config.Database) StrukturOrganisasiRepository {
	return &strukturOrganisasiRepository{db: db}
}

// Profil implementations
func (r *profilRepository) GetAll() ([]entity.Profil, error) {
	var profil []entity.Profil
	err := r.db.DB.Find(&profil).Error
	return profil, err
}

func (r *profilRepository) GetByID(id int) (*entity.Profil, error) {
	var profil entity.Profil
	err := r.db.DB.Where("id = ?", id).First(&profil).Error
	if err != nil {
		return nil, err
	}
	return &profil, nil
}

func (r *profilRepository) Create(profil entity.Profil) (*entity.Profil, error) {
	err := r.db.DB.Create(&profil).Error
	if err != nil {
		return nil, err
	}
	return &profil, nil
}

func (r *profilRepository) Update(id int, profil entity.Profil) (*entity.Profil, error) {
	err := r.db.DB.Model(&entity.Profil{}).Where("id = ?", id).Updates(profil).Error
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *profilRepository) Delete(id int) error {
	return r.db.DB.Delete(&entity.Profil{}, id).Error
}

// VisiMisi implementations
func (r *visimisiRepository) GetAll() ([]entity.VisiMisi, error) {
	var visimisi []entity.VisiMisi
	err := r.db.DB.Find(&visimisi).Error
	return visimisi, err
}

func (r *visimisiRepository) GetByID(id int) (*entity.VisiMisi, error) {
	var visimisi entity.VisiMisi
	err := r.db.DB.Where("id = ?", id).First(&visimisi).Error
	if err != nil {
		return nil, err
	}
	return &visimisi, nil
}

// StrukturOrganisasi implementations
func (r *strukturOrganisasiRepository) GetAll() ([]entity.StrukturOrganisasi, error) {
	var struktur []entity.StrukturOrganisasi
	err := r.db.DB.Find(&struktur).Error
	return struktur, err
}

func (r *strukturOrganisasiRepository) GetByID(id int) (*entity.StrukturOrganisasi, error) {
	var struktur entity.StrukturOrganisasi
	err := r.db.DB.Where("id = ?", id).First(&struktur).Error
	if err != nil {
		return nil, err
	}
	return &struktur, nil
} 