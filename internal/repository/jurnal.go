package repository

import (
	"perpustakaan/config"
	"perpustakaan/internal/entity"
)

type JurnalRepository interface {
	// Jk
	GetAllJk() ([]entity.JurnalJk, error)
	GetJkByID(id int) (*entity.JurnalJk, error)
	CreateJk(jurnal entity.JurnalJk) (*entity.JurnalJk, error)
	UpdateJk(id int, jurnal entity.JurnalJk) (*entity.JurnalJk, error)
	DeleteJk(id int) error

	// Js
	GetAllJs() ([]entity.JurnalJs, error)
	GetJsByID(id int) (*entity.JurnalJs, error)
	CreateJs(jurnal entity.JurnalJs) (*entity.JurnalJs, error)
	UpdateJs(id int, jurnal entity.JurnalJs) (*entity.JurnalJs, error)
	DeleteJs(id int) error

	// Kiat
	GetAllKiat() ([]entity.JurnalKiat, error)
	GetKiatByID(id int) (*entity.JurnalKiat, error)
	CreateKiat(jurnal entity.JurnalKiat) (*entity.JurnalKiat, error)
	UpdateKiat(id int, jurnal entity.JurnalKiat) (*entity.JurnalKiat, error)
	DeleteKiat(id int) error

	// Tajb
	GetAllTajb() ([]entity.JurnalTajb, error)
	GetTajbByID(id int) (*entity.JurnalTajb, error)
	CreateTajb(jurnal entity.JurnalTajb) (*entity.JurnalTajb, error)
	UpdateTajb(id int, jurnal entity.JurnalTajb) (*entity.JurnalTajb, error)
	DeleteTajb(id int) error

	// Teknois
	GetAllTeknois() ([]entity.JurnalTeknois, error)
	GetTeknoisByID(id int) (*entity.JurnalTeknois, error)
	CreateTeknois(jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error)
	UpdateTeknois(id int, jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error)
	DeleteTeknois(id int) error

	// Tmjb
	GetAllTmjb() ([]entity.JurnalTmjb, error)
	GetTmjbByID(id int) (*entity.JurnalTmjb, error)
	CreateTmjb(jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error)
	UpdateTmjb(id int, jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error)
	DeleteTmjb(id int) error
}

type jurnalRepository struct {
	db *config.Database
}

func NewJurnalRepository(db *config.Database) JurnalRepository {
	return &jurnalRepository{db: db}
}

// Jk implementations
func (r *jurnalRepository) GetAllJk() ([]entity.JurnalJk, error) {
	var jurnal []entity.JurnalJk
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetJkByID(id int) (*entity.JurnalJk, error) {
	var jurnal entity.JurnalJk
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateJk(jurnal entity.JurnalJk) (*entity.JurnalJk, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateJk(id int, jurnal entity.JurnalJk) (*entity.JurnalJk, error) {
	err := r.db.DB.Model(&entity.JurnalJk{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetJkByID(id)
}

func (r *jurnalRepository) DeleteJk(id int) error {
	return r.db.DB.Delete(&entity.JurnalJk{}, id).Error
}

// Js implementations
func (r *jurnalRepository) GetAllJs() ([]entity.JurnalJs, error) {
	var jurnal []entity.JurnalJs
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetJsByID(id int) (*entity.JurnalJs, error) {
	var jurnal entity.JurnalJs
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateJs(jurnal entity.JurnalJs) (*entity.JurnalJs, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateJs(id int, jurnal entity.JurnalJs) (*entity.JurnalJs, error) {
	err := r.db.DB.Model(&entity.JurnalJs{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetJsByID(id)
}

func (r *jurnalRepository) DeleteJs(id int) error {
	return r.db.DB.Delete(&entity.JurnalJs{}, id).Error
}

// Kiat implementations
func (r *jurnalRepository) GetAllKiat() ([]entity.JurnalKiat, error) {
	var jurnal []entity.JurnalKiat
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetKiatByID(id int) (*entity.JurnalKiat, error) {
	var jurnal entity.JurnalKiat
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateKiat(jurnal entity.JurnalKiat) (*entity.JurnalKiat, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateKiat(id int, jurnal entity.JurnalKiat) (*entity.JurnalKiat, error) {
	err := r.db.DB.Model(&entity.JurnalKiat{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetKiatByID(id)
}

func (r *jurnalRepository) DeleteKiat(id int) error {
	return r.db.DB.Delete(&entity.JurnalKiat{}, id).Error
}

// Tajb implementations
func (r *jurnalRepository) GetAllTajb() ([]entity.JurnalTajb, error) {
	var jurnal []entity.JurnalTajb
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetTajbByID(id int) (*entity.JurnalTajb, error) {
	var jurnal entity.JurnalTajb
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateTajb(jurnal entity.JurnalTajb) (*entity.JurnalTajb, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateTajb(id int, jurnal entity.JurnalTajb) (*entity.JurnalTajb, error) {
	err := r.db.DB.Model(&entity.JurnalTajb{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetTajbByID(id)
}

func (r *jurnalRepository) DeleteTajb(id int) error {
	return r.db.DB.Delete(&entity.JurnalTajb{}, id).Error
}

// Teknois implementations
func (r *jurnalRepository) GetAllTeknois() ([]entity.JurnalTeknois, error) {
	var jurnal []entity.JurnalTeknois
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetTeknoisByID(id int) (*entity.JurnalTeknois, error) {
	var jurnal entity.JurnalTeknois
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateTeknois(jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateTeknois(id int, jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error) {
	err := r.db.DB.Model(&entity.JurnalTeknois{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetTeknoisByID(id)
}

func (r *jurnalRepository) DeleteTeknois(id int) error {
	return r.db.DB.Delete(&entity.JurnalTeknois{}, id).Error
}

// Tmjb implementations
func (r *jurnalRepository) GetAllTmjb() ([]entity.JurnalTmjb, error) {
	var jurnal []entity.JurnalTmjb
	err := r.db.DB.Find(&jurnal).Error
	return jurnal, err
}

func (r *jurnalRepository) GetTmjbByID(id int) (*entity.JurnalTmjb, error) {
	var jurnal entity.JurnalTmjb
	err := r.db.DB.Where("id = ?", id).First(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) CreateTmjb(jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error) {
	err := r.db.DB.Create(&jurnal).Error
	if err != nil {
		return nil, err
	}
	return &jurnal, nil
}

func (r *jurnalRepository) UpdateTmjb(id int, jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error) {
	err := r.db.DB.Model(&entity.JurnalTmjb{}).Where("id = ?", id).Updates(jurnal).Error
	if err != nil {
		return nil, err
	}
	return r.GetTmjbByID(id)
}

func (r *jurnalRepository) DeleteTmjb(id int) error {
	return r.db.DB.Delete(&entity.JurnalTmjb{}, id).Error
} 