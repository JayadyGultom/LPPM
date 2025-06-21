package repository

import (
	"perpustakaan/config"
	"perpustakaan/internal/entity"
)

type PenelitianRepository interface {
	// Bame
	GetAllBame() ([]entity.PenelitianBame, error)
	GetBameByID(id int) (*entity.PenelitianBame, error)
	CreateBame(penelitian entity.PenelitianBame) (*entity.PenelitianBame, error)
	UpdateBame(id int, penelitian entity.PenelitianBame) (*entity.PenelitianBame, error)
	DeleteBame(id int) error

	// Hpp
	GetAllHpp() ([]entity.PenelitianHpp, error)
	GetHppByID(id int) (*entity.PenelitianHpp, error)
	CreateHpp(penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error)
	UpdateHpp(id int, penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error)
	DeleteHpp(id int) error

	// Lp
	GetAllLp() ([]entity.PenelitianLp, error)
	GetLpByID(id int) (*entity.PenelitianLp, error)
	CreateLp(penelitian entity.PenelitianLp) (*entity.PenelitianLp, error)
	UpdateLp(id int, penelitian entity.PenelitianLp) (*entity.PenelitianLp, error)
	DeleteLp(id int) error

	// P3
	GetAllP3() ([]entity.PenelitianP3, error)
	GetP3ByID(id int) (*entity.PenelitianP3, error)
	CreateP3(penelitian entity.PenelitianP3) (*entity.PenelitianP3, error)
	UpdateP3(id int, penelitian entity.PenelitianP3) (*entity.PenelitianP3, error)
	DeleteP3(id int) error

	// Rrp
	GetAllRrp() ([]entity.PenelitianRrp, error)
	GetRrpByID(id int) (*entity.PenelitianRrp, error)
	CreateRrp(penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error)
	UpdateRrp(id int, penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error)
	DeleteRrp(id int) error

	// SkR
	GetAllSkR() ([]entity.PenelitianSkR, error)
	GetSkRByID(id int) (*entity.PenelitianSkR, error)
	CreateSkR(penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error)
	UpdateSkR(id int, penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error)
	DeleteSkR(id int) error

	// Stp
	GetAllStp() ([]entity.PenelitianStp, error)
	GetStpByID(id int) (*entity.PenelitianStp, error)
	CreateStp(penelitian entity.PenelitianStp) (*entity.PenelitianStp, error)
	UpdateStp(id int, penelitian entity.PenelitianStp) (*entity.PenelitianStp, error)
	DeleteStp(id int) error

	// Tcr
	GetAllTcr() ([]entity.PenelitianTcr, error)
	GetTcrByID(id int) (*entity.PenelitianTcr, error)
	CreateTcr(penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error)
	UpdateTcr(id int, penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error)
	DeleteTcr(id int) error
}

type penelitianRepository struct {
	db *config.Database
}

func NewPenelitianRepository(db *config.Database) PenelitianRepository {
	return &penelitianRepository{db: db}
}

// Bame implementations
func (r *penelitianRepository) GetAllBame() ([]entity.PenelitianBame, error) {
	var penelitian []entity.PenelitianBame
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetBameByID(id int) (*entity.PenelitianBame, error) {
	var penelitian entity.PenelitianBame
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateBame(penelitian entity.PenelitianBame) (*entity.PenelitianBame, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateBame(id int, penelitian entity.PenelitianBame) (*entity.PenelitianBame, error) {
	err := r.db.DB.Model(&entity.PenelitianBame{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetBameByID(id)
}

func (r *penelitianRepository) DeleteBame(id int) error {
	return r.db.DB.Delete(&entity.PenelitianBame{}, id).Error
}

// Hpp implementations
func (r *penelitianRepository) GetAllHpp() ([]entity.PenelitianHpp, error) {
	var penelitian []entity.PenelitianHpp
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetHppByID(id int) (*entity.PenelitianHpp, error) {
	var penelitian entity.PenelitianHpp
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateHpp(penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateHpp(id int, penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error) {
	err := r.db.DB.Model(&entity.PenelitianHpp{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetHppByID(id)
}

func (r *penelitianRepository) DeleteHpp(id int) error {
	return r.db.DB.Delete(&entity.PenelitianHpp{}, id).Error
}

// Lp implementations
func (r *penelitianRepository) GetAllLp() ([]entity.PenelitianLp, error) {
	var penelitian []entity.PenelitianLp
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetLpByID(id int) (*entity.PenelitianLp, error) {
	var penelitian entity.PenelitianLp
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateLp(penelitian entity.PenelitianLp) (*entity.PenelitianLp, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateLp(id int, penelitian entity.PenelitianLp) (*entity.PenelitianLp, error) {
	err := r.db.DB.Model(&entity.PenelitianLp{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetLpByID(id)
}

func (r *penelitianRepository) DeleteLp(id int) error {
	return r.db.DB.Delete(&entity.PenelitianLp{}, id).Error
}

// P3 implementations
func (r *penelitianRepository) GetAllP3() ([]entity.PenelitianP3, error) {
	var penelitian []entity.PenelitianP3
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetP3ByID(id int) (*entity.PenelitianP3, error) {
	var penelitian entity.PenelitianP3
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateP3(penelitian entity.PenelitianP3) (*entity.PenelitianP3, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateP3(id int, penelitian entity.PenelitianP3) (*entity.PenelitianP3, error) {
	err := r.db.DB.Model(&entity.PenelitianP3{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetP3ByID(id)
}

func (r *penelitianRepository) DeleteP3(id int) error {
	return r.db.DB.Delete(&entity.PenelitianP3{}, id).Error
}

// Rrp implementations
func (r *penelitianRepository) GetAllRrp() ([]entity.PenelitianRrp, error) {
	var penelitian []entity.PenelitianRrp
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetRrpByID(id int) (*entity.PenelitianRrp, error) {
	var penelitian entity.PenelitianRrp
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateRrp(penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateRrp(id int, penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error) {
	err := r.db.DB.Model(&entity.PenelitianRrp{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetRrpByID(id)
}

func (r *penelitianRepository) DeleteRrp(id int) error {
	return r.db.DB.Delete(&entity.PenelitianRrp{}, id).Error
}

// SkR implementations
func (r *penelitianRepository) GetAllSkR() ([]entity.PenelitianSkR, error) {
	var penelitian []entity.PenelitianSkR
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetSkRByID(id int) (*entity.PenelitianSkR, error) {
	var penelitian entity.PenelitianSkR
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateSkR(penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateSkR(id int, penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error) {
	err := r.db.DB.Model(&entity.PenelitianSkR{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetSkRByID(id)
}

func (r *penelitianRepository) DeleteSkR(id int) error {
	return r.db.DB.Delete(&entity.PenelitianSkR{}, id).Error
}

// Stp implementations
func (r *penelitianRepository) GetAllStp() ([]entity.PenelitianStp, error) {
	var penelitian []entity.PenelitianStp
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetStpByID(id int) (*entity.PenelitianStp, error) {
	var penelitian entity.PenelitianStp
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateStp(penelitian entity.PenelitianStp) (*entity.PenelitianStp, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateStp(id int, penelitian entity.PenelitianStp) (*entity.PenelitianStp, error) {
	err := r.db.DB.Model(&entity.PenelitianStp{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetStpByID(id)
}

func (r *penelitianRepository) DeleteStp(id int) error {
	return r.db.DB.Delete(&entity.PenelitianStp{}, id).Error
}

// Tcr implementations
func (r *penelitianRepository) GetAllTcr() ([]entity.PenelitianTcr, error) {
	var penelitian []entity.PenelitianTcr
	err := r.db.DB.Find(&penelitian).Error
	return penelitian, err
}

func (r *penelitianRepository) GetTcrByID(id int) (*entity.PenelitianTcr, error) {
	var penelitian entity.PenelitianTcr
	err := r.db.DB.Where("id = ?", id).First(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) CreateTcr(penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error) {
	err := r.db.DB.Create(&penelitian).Error
	if err != nil {
		return nil, err
	}
	return &penelitian, nil
}

func (r *penelitianRepository) UpdateTcr(id int, penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error) {
	err := r.db.DB.Model(&entity.PenelitianTcr{}).Where("id = ?", id).Updates(penelitian).Error
	if err != nil {
		return nil, err
	}
	return r.GetTcrByID(id)
}

func (r *penelitianRepository) DeleteTcr(id int) error {
	return r.db.DB.Delete(&entity.PenelitianTcr{}, id).Error
} 