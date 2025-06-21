package repository

import (
	"perpustakaan/config"
	"perpustakaan/internal/entity"
)

type PKMRepository interface {
	// Bame
	GetAllBame() ([]entity.PKMBame, error)
	GetBameByID(id int) (*entity.PKMBame, error)
	CreateBame(pkm entity.PKMBame) (*entity.PKMBame, error)
	UpdateBame(id int, pkm entity.PKMBame) (*entity.PKMBame, error)
	DeleteBame(id int) error

	// Hpp
	GetAllHpp() ([]entity.PKMHpp, error)
	GetHppByID(id int) (*entity.PKMHpp, error)
	CreateHpp(pkm entity.PKMHpp) (*entity.PKMHpp, error)
	UpdateHpp(id int, pkm entity.PKMHpp) (*entity.PKMHpp, error)
	DeleteHpp(id int) error

	// Lp
	GetAllLp() ([]entity.PKMLp, error)
	GetLpByID(id int) (*entity.PKMLp, error)
	CreateLp(pkm entity.PKMLp) (*entity.PKMLp, error)
	UpdateLp(id int, pkm entity.PKMLp) (*entity.PKMLp, error)
	DeleteLp(id int) error

	// P3
	GetAllP3() ([]entity.PKMP3, error)
	GetP3ByID(id int) (*entity.PKMP3, error)
	CreateP3(pkm entity.PKMP3) (*entity.PKMP3, error)
	UpdateP3(id int, pkm entity.PKMP3) (*entity.PKMP3, error)
	DeleteP3(id int) error

	// Rrp
	GetAllRrp() ([]entity.PKMRrp, error)
	GetRrpByID(id int) (*entity.PKMRrp, error)
	CreateRrp(pkm entity.PKMRrp) (*entity.PKMRrp, error)
	UpdateRrp(id int, pkm entity.PKMRrp) (*entity.PKMRrp, error)
	DeleteRrp(id int) error

	// SkR
	GetAllSkR() ([]entity.PKMSkR, error)
	GetSkRByID(id int) (*entity.PKMSkR, error)
	CreateSkR(pkm entity.PKMSkR) (*entity.PKMSkR, error)
	UpdateSkR(id int, pkm entity.PKMSkR) (*entity.PKMSkR, error)
	DeleteSkR(id int) error

	// Stp
	GetAllStp() ([]entity.PKMStp, error)
	GetStpByID(id int) (*entity.PKMStp, error)
	CreateStp(pkm entity.PKMStp) (*entity.PKMStp, error)
	UpdateStp(id int, pkm entity.PKMStp) (*entity.PKMStp, error)
	DeleteStp(id int) error

	// Tcr
	GetAllTcr() ([]entity.PKMTcr, error)
	GetTcrByID(id int) (*entity.PKMTcr, error)
	CreateTcr(pkm entity.PKMTcr) (*entity.PKMTcr, error)
	UpdateTcr(id int, pkm entity.PKMTcr) (*entity.PKMTcr, error)
	DeleteTcr(id int) error
}

type pkmRepository struct {
	db *config.Database
}

func NewPKMRepository(db *config.Database) PKMRepository {
	return &pkmRepository{db: db}
}

// Bame implementations
func (r *pkmRepository) GetAllBame() ([]entity.PKMBame, error) {
	var pkm []entity.PKMBame
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetBameByID(id int) (*entity.PKMBame, error) {
	var pkm entity.PKMBame
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateBame(pkm entity.PKMBame) (*entity.PKMBame, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateBame(id int, pkm entity.PKMBame) (*entity.PKMBame, error) {
	err := r.db.DB.Model(&entity.PKMBame{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetBameByID(id)
}

func (r *pkmRepository) DeleteBame(id int) error {
	return r.db.DB.Delete(&entity.PKMBame{}, id).Error
}

// Hpp implementations
func (r *pkmRepository) GetAllHpp() ([]entity.PKMHpp, error) {
	var pkm []entity.PKMHpp
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetHppByID(id int) (*entity.PKMHpp, error) {
	var pkm entity.PKMHpp
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateHpp(pkm entity.PKMHpp) (*entity.PKMHpp, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateHpp(id int, pkm entity.PKMHpp) (*entity.PKMHpp, error) {
	err := r.db.DB.Model(&entity.PKMHpp{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetHppByID(id)
}

func (r *pkmRepository) DeleteHpp(id int) error {
	return r.db.DB.Delete(&entity.PKMHpp{}, id).Error
}

// Lp implementations
func (r *pkmRepository) GetAllLp() ([]entity.PKMLp, error) {
	var pkm []entity.PKMLp
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetLpByID(id int) (*entity.PKMLp, error) {
	var pkm entity.PKMLp
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateLp(pkm entity.PKMLp) (*entity.PKMLp, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateLp(id int, pkm entity.PKMLp) (*entity.PKMLp, error) {
	err := r.db.DB.Model(&entity.PKMLp{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetLpByID(id)
}

func (r *pkmRepository) DeleteLp(id int) error {
	return r.db.DB.Delete(&entity.PKMLp{}, id).Error
}

// P3 implementations
func (r *pkmRepository) GetAllP3() ([]entity.PKMP3, error) {
	var pkm []entity.PKMP3
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetP3ByID(id int) (*entity.PKMP3, error) {
	var pkm entity.PKMP3
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateP3(pkm entity.PKMP3) (*entity.PKMP3, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateP3(id int, pkm entity.PKMP3) (*entity.PKMP3, error) {
	err := r.db.DB.Model(&entity.PKMP3{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetP3ByID(id)
}

func (r *pkmRepository) DeleteP3(id int) error {
	return r.db.DB.Delete(&entity.PKMP3{}, id).Error
}

// Rrp implementations
func (r *pkmRepository) GetAllRrp() ([]entity.PKMRrp, error) {
	var pkm []entity.PKMRrp
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetRrpByID(id int) (*entity.PKMRrp, error) {
	var pkm entity.PKMRrp
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateRrp(pkm entity.PKMRrp) (*entity.PKMRrp, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateRrp(id int, pkm entity.PKMRrp) (*entity.PKMRrp, error) {
	err := r.db.DB.Model(&entity.PKMRrp{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetRrpByID(id)
}

func (r *pkmRepository) DeleteRrp(id int) error {
	return r.db.DB.Delete(&entity.PKMRrp{}, id).Error
}

// SkR implementations
func (r *pkmRepository) GetAllSkR() ([]entity.PKMSkR, error) {
	var pkm []entity.PKMSkR
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetSkRByID(id int) (*entity.PKMSkR, error) {
	var pkm entity.PKMSkR
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateSkR(pkm entity.PKMSkR) (*entity.PKMSkR, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateSkR(id int, pkm entity.PKMSkR) (*entity.PKMSkR, error) {
	err := r.db.DB.Model(&entity.PKMSkR{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetSkRByID(id)
}

func (r *pkmRepository) DeleteSkR(id int) error {
	return r.db.DB.Delete(&entity.PKMSkR{}, id).Error
}

// Stp implementations
func (r *pkmRepository) GetAllStp() ([]entity.PKMStp, error) {
	var pkm []entity.PKMStp
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetStpByID(id int) (*entity.PKMStp, error) {
	var pkm entity.PKMStp
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateStp(pkm entity.PKMStp) (*entity.PKMStp, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateStp(id int, pkm entity.PKMStp) (*entity.PKMStp, error) {
	err := r.db.DB.Model(&entity.PKMStp{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetStpByID(id)
}

func (r *pkmRepository) DeleteStp(id int) error {
	return r.db.DB.Delete(&entity.PKMStp{}, id).Error
}

// Tcr implementations
func (r *pkmRepository) GetAllTcr() ([]entity.PKMTcr, error) {
	var pkm []entity.PKMTcr
	err := r.db.DB.Find(&pkm).Error
	return pkm, err
}

func (r *pkmRepository) GetTcrByID(id int) (*entity.PKMTcr, error) {
	var pkm entity.PKMTcr
	err := r.db.DB.Where("id = ?", id).First(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) CreateTcr(pkm entity.PKMTcr) (*entity.PKMTcr, error) {
	err := r.db.DB.Create(&pkm).Error
	if err != nil {
		return nil, err
	}
	return &pkm, nil
}

func (r *pkmRepository) UpdateTcr(id int, pkm entity.PKMTcr) (*entity.PKMTcr, error) {
	err := r.db.DB.Model(&entity.PKMTcr{}).Where("id = ?", id).Updates(pkm).Error
	if err != nil {
		return nil, err
	}
	return r.GetTcrByID(id)
}

func (r *pkmRepository) DeleteTcr(id int) error {
	return r.db.DB.Delete(&entity.PKMTcr{}, id).Error
} 