package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PKMUseCase interface {
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

type pkmUseCase struct {
	pkmRepo repository.PKMRepository
}

func NewPKMUseCase(pkmRepo repository.PKMRepository) PKMUseCase {
	return &pkmUseCase{
		pkmRepo: pkmRepo,
	}
}

// Bame implementations
func (uc *pkmUseCase) GetAllBame() ([]entity.PKMBame, error) {
	return uc.pkmRepo.GetAllBame()
}

func (uc *pkmUseCase) GetBameByID(id int) (*entity.PKMBame, error) {
	return uc.pkmRepo.GetBameByID(id)
}

func (uc *pkmUseCase) CreateBame(pkm entity.PKMBame) (*entity.PKMBame, error) {
	return uc.pkmRepo.CreateBame(pkm)
}

func (uc *pkmUseCase) UpdateBame(id int, pkm entity.PKMBame) (*entity.PKMBame, error) {
	return uc.pkmRepo.UpdateBame(id, pkm)
}

func (uc *pkmUseCase) DeleteBame(id int) error {
	return uc.pkmRepo.DeleteBame(id)
}

// Hpp implementations
func (uc *pkmUseCase) GetAllHpp() ([]entity.PKMHpp, error) {
	return uc.pkmRepo.GetAllHpp()
}

func (uc *pkmUseCase) GetHppByID(id int) (*entity.PKMHpp, error) {
	return uc.pkmRepo.GetHppByID(id)
}

func (uc *pkmUseCase) CreateHpp(pkm entity.PKMHpp) (*entity.PKMHpp, error) {
	return uc.pkmRepo.CreateHpp(pkm)
}

func (uc *pkmUseCase) UpdateHpp(id int, pkm entity.PKMHpp) (*entity.PKMHpp, error) {
	return uc.pkmRepo.UpdateHpp(id, pkm)
}

func (uc *pkmUseCase) DeleteHpp(id int) error {
	return uc.pkmRepo.DeleteHpp(id)
}

// Lp implementations
func (uc *pkmUseCase) GetAllLp() ([]entity.PKMLp, error) {
	return uc.pkmRepo.GetAllLp()
}

func (uc *pkmUseCase) GetLpByID(id int) (*entity.PKMLp, error) {
	return uc.pkmRepo.GetLpByID(id)
}

func (uc *pkmUseCase) CreateLp(pkm entity.PKMLp) (*entity.PKMLp, error) {
	return uc.pkmRepo.CreateLp(pkm)
}

func (uc *pkmUseCase) UpdateLp(id int, pkm entity.PKMLp) (*entity.PKMLp, error) {
	return uc.pkmRepo.UpdateLp(id, pkm)
}

func (uc *pkmUseCase) DeleteLp(id int) error {
	return uc.pkmRepo.DeleteLp(id)
}

// P3 implementations
func (uc *pkmUseCase) GetAllP3() ([]entity.PKMP3, error) {
	return uc.pkmRepo.GetAllP3()
}

func (uc *pkmUseCase) GetP3ByID(id int) (*entity.PKMP3, error) {
	return uc.pkmRepo.GetP3ByID(id)
}

func (uc *pkmUseCase) CreateP3(pkm entity.PKMP3) (*entity.PKMP3, error) {
	return uc.pkmRepo.CreateP3(pkm)
}

func (uc *pkmUseCase) UpdateP3(id int, pkm entity.PKMP3) (*entity.PKMP3, error) {
	return uc.pkmRepo.UpdateP3(id, pkm)
}

func (uc *pkmUseCase) DeleteP3(id int) error {
	return uc.pkmRepo.DeleteP3(id)
}

// Rrp implementations
func (uc *pkmUseCase) GetAllRrp() ([]entity.PKMRrp, error) {
	return uc.pkmRepo.GetAllRrp()
}

func (uc *pkmUseCase) GetRrpByID(id int) (*entity.PKMRrp, error) {
	return uc.pkmRepo.GetRrpByID(id)
}

func (uc *pkmUseCase) CreateRrp(pkm entity.PKMRrp) (*entity.PKMRrp, error) {
	return uc.pkmRepo.CreateRrp(pkm)
}

func (uc *pkmUseCase) UpdateRrp(id int, pkm entity.PKMRrp) (*entity.PKMRrp, error) {
	return uc.pkmRepo.UpdateRrp(id, pkm)
}

func (uc *pkmUseCase) DeleteRrp(id int) error {
	return uc.pkmRepo.DeleteRrp(id)
}

// SkR implementations
func (uc *pkmUseCase) GetAllSkR() ([]entity.PKMSkR, error) {
	return uc.pkmRepo.GetAllSkR()
}

func (uc *pkmUseCase) GetSkRByID(id int) (*entity.PKMSkR, error) {
	return uc.pkmRepo.GetSkRByID(id)
}

func (uc *pkmUseCase) CreateSkR(pkm entity.PKMSkR) (*entity.PKMSkR, error) {
	return uc.pkmRepo.CreateSkR(pkm)
}

func (uc *pkmUseCase) UpdateSkR(id int, pkm entity.PKMSkR) (*entity.PKMSkR, error) {
	return uc.pkmRepo.UpdateSkR(id, pkm)
}

func (uc *pkmUseCase) DeleteSkR(id int) error {
	return uc.pkmRepo.DeleteSkR(id)
}

// Stp implementations
func (uc *pkmUseCase) GetAllStp() ([]entity.PKMStp, error) {
	return uc.pkmRepo.GetAllStp()
}

func (uc *pkmUseCase) GetStpByID(id int) (*entity.PKMStp, error) {
	return uc.pkmRepo.GetStpByID(id)
}

func (uc *pkmUseCase) CreateStp(pkm entity.PKMStp) (*entity.PKMStp, error) {
	return uc.pkmRepo.CreateStp(pkm)
}

func (uc *pkmUseCase) UpdateStp(id int, pkm entity.PKMStp) (*entity.PKMStp, error) {
	return uc.pkmRepo.UpdateStp(id, pkm)
}

func (uc *pkmUseCase) DeleteStp(id int) error {
	return uc.pkmRepo.DeleteStp(id)
}

// Tcr implementations
func (uc *pkmUseCase) GetAllTcr() ([]entity.PKMTcr, error) {
	return uc.pkmRepo.GetAllTcr()
}

func (uc *pkmUseCase) GetTcrByID(id int) (*entity.PKMTcr, error) {
	return uc.pkmRepo.GetTcrByID(id)
}

func (uc *pkmUseCase) CreateTcr(pkm entity.PKMTcr) (*entity.PKMTcr, error) {
	return uc.pkmRepo.CreateTcr(pkm)
}

func (uc *pkmUseCase) UpdateTcr(id int, pkm entity.PKMTcr) (*entity.PKMTcr, error) {
	return uc.pkmRepo.UpdateTcr(id, pkm)
}

func (uc *pkmUseCase) DeleteTcr(id int) error {
	return uc.pkmRepo.DeleteTcr(id)
} 