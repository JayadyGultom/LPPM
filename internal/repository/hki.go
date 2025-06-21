package repository

import (
	"perpustakaan/config"
	"perpustakaan/internal/entity"
)

type HKIRepository interface {
	// Dosen
	GetAllDosen() ([]entity.HKIDosen, error)
	GetDosenByID(id int) (*entity.HKIDosen, error)
	CreateDosen(hki entity.HKIDosen) (*entity.HKIDosen, error)
	UpdateDosen(id int, hki entity.HKIDosen) (*entity.HKIDosen, error)
	DeleteDosen(id int) error

	// Mahasiswa
	GetAllMhs() ([]entity.HKIMhs, error)
	GetMhsByID(id int) (*entity.HKIMhs, error)
	CreateMhs(hki entity.HKIMhs) (*entity.HKIMhs, error)
	UpdateMhs(id int, hki entity.HKIMhs) (*entity.HKIMhs, error)
	DeleteMhs(id int) error
}

type hkiRepository struct {
	db *config.Database
}

func NewHKIRepository(db *config.Database) HKIRepository {
	return &hkiRepository{db: db}
}

// Dosen implementations
func (r *hkiRepository) GetAllDosen() ([]entity.HKIDosen, error) {
	var hki []entity.HKIDosen
	err := r.db.DB.Find(&hki).Error
	return hki, err
}

func (r *hkiRepository) GetDosenByID(id int) (*entity.HKIDosen, error) {
	var hki entity.HKIDosen
	err := r.db.DB.Where("id = ?", id).First(&hki).Error
	if err != nil {
		return nil, err
	}
	return &hki, nil
}

func (r *hkiRepository) CreateDosen(hki entity.HKIDosen) (*entity.HKIDosen, error) {
	err := r.db.DB.Create(&hki).Error
	if err != nil {
		return nil, err
	}
	return &hki, nil
}

func (r *hkiRepository) UpdateDosen(id int, hki entity.HKIDosen) (*entity.HKIDosen, error) {
	err := r.db.DB.Model(&entity.HKIDosen{}).Where("id = ?", id).Updates(hki).Error
	if err != nil {
		return nil, err
	}
	return r.GetDosenByID(id)
}

func (r *hkiRepository) DeleteDosen(id int) error {
	return r.db.DB.Delete(&entity.HKIDosen{}, id).Error
}

// Mahasiswa implementations
func (r *hkiRepository) GetAllMhs() ([]entity.HKIMhs, error) {
	var hki []entity.HKIMhs
	err := r.db.DB.Find(&hki).Error
	return hki, err
}

func (r *hkiRepository) GetMhsByID(id int) (*entity.HKIMhs, error) {
	var hki entity.HKIMhs
	err := r.db.DB.Where("id = ?", id).First(&hki).Error
	if err != nil {
		return nil, err
	}
	return &hki, nil
}

func (r *hkiRepository) CreateMhs(hki entity.HKIMhs) (*entity.HKIMhs, error) {
	err := r.db.DB.Create(&hki).Error
	if err != nil {
		return nil, err
	}
	return &hki, nil
}

func (r *hkiRepository) UpdateMhs(id int, hki entity.HKIMhs) (*entity.HKIMhs, error) {
	err := r.db.DB.Model(&entity.HKIMhs{}).Where("id = ?", id).Updates(hki).Error
	if err != nil {
		return nil, err
	}
	return r.GetMhsByID(id)
}

func (r *hkiRepository) DeleteMhs(id int) error {
	return r.db.DB.Delete(&entity.HKIMhs{}, id).Error
} 