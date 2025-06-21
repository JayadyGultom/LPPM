package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type PenelitianUseCase interface {
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

type penelitianUseCase struct {
	penelitianRepo repository.PenelitianRepository
}

func NewPenelitianUseCase(penelitianRepo repository.PenelitianRepository) PenelitianUseCase {
	return &penelitianUseCase{
		penelitianRepo: penelitianRepo,
	}
}

// Bame implementations
func (uc *penelitianUseCase) GetAllBame() ([]entity.PenelitianBame, error) {
	return uc.penelitianRepo.GetAllBame()
}

func (uc *penelitianUseCase) GetBameByID(id int) (*entity.PenelitianBame, error) {
	return uc.penelitianRepo.GetBameByID(id)
}

func (uc *penelitianUseCase) CreateBame(penelitian entity.PenelitianBame) (*entity.PenelitianBame, error) {
	return uc.penelitianRepo.CreateBame(penelitian)
}

func (uc *penelitianUseCase) UpdateBame(id int, penelitian entity.PenelitianBame) (*entity.PenelitianBame, error) {
	return uc.penelitianRepo.UpdateBame(id, penelitian)
}

func (uc *penelitianUseCase) DeleteBame(id int) error {
	return uc.penelitianRepo.DeleteBame(id)
}

// Hpp implementations
func (uc *penelitianUseCase) GetAllHpp() ([]entity.PenelitianHpp, error) {
	return uc.penelitianRepo.GetAllHpp()
}

func (uc *penelitianUseCase) GetHppByID(id int) (*entity.PenelitianHpp, error) {
	return uc.penelitianRepo.GetHppByID(id)
}

func (uc *penelitianUseCase) CreateHpp(penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error) {
	return uc.penelitianRepo.CreateHpp(penelitian)
}

func (uc *penelitianUseCase) UpdateHpp(id int, penelitian entity.PenelitianHpp) (*entity.PenelitianHpp, error) {
	return uc.penelitianRepo.UpdateHpp(id, penelitian)
}

func (uc *penelitianUseCase) DeleteHpp(id int) error {
	return uc.penelitianRepo.DeleteHpp(id)
}

// Lp implementations
func (uc *penelitianUseCase) GetAllLp() ([]entity.PenelitianLp, error) {
	return uc.penelitianRepo.GetAllLp()
}

func (uc *penelitianUseCase) GetLpByID(id int) (*entity.PenelitianLp, error) {
	return uc.penelitianRepo.GetLpByID(id)
}

func (uc *penelitianUseCase) CreateLp(penelitian entity.PenelitianLp) (*entity.PenelitianLp, error) {
	return uc.penelitianRepo.CreateLp(penelitian)
}

func (uc *penelitianUseCase) UpdateLp(id int, penelitian entity.PenelitianLp) (*entity.PenelitianLp, error) {
	return uc.penelitianRepo.UpdateLp(id, penelitian)
}

func (uc *penelitianUseCase) DeleteLp(id int) error {
	return uc.penelitianRepo.DeleteLp(id)
}

// P3 implementations
func (uc *penelitianUseCase) GetAllP3() ([]entity.PenelitianP3, error) {
	return uc.penelitianRepo.GetAllP3()
}

func (uc *penelitianUseCase) GetP3ByID(id int) (*entity.PenelitianP3, error) {
	return uc.penelitianRepo.GetP3ByID(id)
}

func (uc *penelitianUseCase) CreateP3(penelitian entity.PenelitianP3) (*entity.PenelitianP3, error) {
	return uc.penelitianRepo.CreateP3(penelitian)
}

func (uc *penelitianUseCase) UpdateP3(id int, penelitian entity.PenelitianP3) (*entity.PenelitianP3, error) {
	return uc.penelitianRepo.UpdateP3(id, penelitian)
}

func (uc *penelitianUseCase) DeleteP3(id int) error {
	return uc.penelitianRepo.DeleteP3(id)
}

// Rrp implementations
func (uc *penelitianUseCase) GetAllRrp() ([]entity.PenelitianRrp, error) {
	return uc.penelitianRepo.GetAllRrp()
}

func (uc *penelitianUseCase) GetRrpByID(id int) (*entity.PenelitianRrp, error) {
	return uc.penelitianRepo.GetRrpByID(id)
}

func (uc *penelitianUseCase) CreateRrp(penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error) {
	return uc.penelitianRepo.CreateRrp(penelitian)
}

func (uc *penelitianUseCase) UpdateRrp(id int, penelitian entity.PenelitianRrp) (*entity.PenelitianRrp, error) {
	return uc.penelitianRepo.UpdateRrp(id, penelitian)
}

func (uc *penelitianUseCase) DeleteRrp(id int) error {
	return uc.penelitianRepo.DeleteRrp(id)
}

// SkR implementations
func (uc *penelitianUseCase) GetAllSkR() ([]entity.PenelitianSkR, error) {
	return uc.penelitianRepo.GetAllSkR()
}

func (uc *penelitianUseCase) GetSkRByID(id int) (*entity.PenelitianSkR, error) {
	return uc.penelitianRepo.GetSkRByID(id)
}

func (uc *penelitianUseCase) CreateSkR(penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error) {
	return uc.penelitianRepo.CreateSkR(penelitian)
}

func (uc *penelitianUseCase) UpdateSkR(id int, penelitian entity.PenelitianSkR) (*entity.PenelitianSkR, error) {
	return uc.penelitianRepo.UpdateSkR(id, penelitian)
}

func (uc *penelitianUseCase) DeleteSkR(id int) error {
	return uc.penelitianRepo.DeleteSkR(id)
}

// Stp implementations
func (uc *penelitianUseCase) GetAllStp() ([]entity.PenelitianStp, error) {
	return uc.penelitianRepo.GetAllStp()
}

func (uc *penelitianUseCase) GetStpByID(id int) (*entity.PenelitianStp, error) {
	return uc.penelitianRepo.GetStpByID(id)
}

func (uc *penelitianUseCase) CreateStp(penelitian entity.PenelitianStp) (*entity.PenelitianStp, error) {
	return uc.penelitianRepo.CreateStp(penelitian)
}

func (uc *penelitianUseCase) UpdateStp(id int, penelitian entity.PenelitianStp) (*entity.PenelitianStp, error) {
	return uc.penelitianRepo.UpdateStp(id, penelitian)
}

func (uc *penelitianUseCase) DeleteStp(id int) error {
	return uc.penelitianRepo.DeleteStp(id)
}

// Tcr implementations
func (uc *penelitianUseCase) GetAllTcr() ([]entity.PenelitianTcr, error) {
	return uc.penelitianRepo.GetAllTcr()
}

func (uc *penelitianUseCase) GetTcrByID(id int) (*entity.PenelitianTcr, error) {
	return uc.penelitianRepo.GetTcrByID(id)
}

func (uc *penelitianUseCase) CreateTcr(penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error) {
	return uc.penelitianRepo.CreateTcr(penelitian)
}

func (uc *penelitianUseCase) UpdateTcr(id int, penelitian entity.PenelitianTcr) (*entity.PenelitianTcr, error) {
	return uc.penelitianRepo.UpdateTcr(id, penelitian)
}

func (uc *penelitianUseCase) DeleteTcr(id int) error {
	return uc.penelitianRepo.DeleteTcr(id)
} 