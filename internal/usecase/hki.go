package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type HKIUseCase interface {
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

type hkiUseCase struct {
	hkiRepo repository.HKIRepository
}

func NewHKIUseCase(hkiRepo repository.HKIRepository) HKIUseCase {
	return &hkiUseCase{
		hkiRepo: hkiRepo,
	}
}

// Dosen implementations
func (uc *hkiUseCase) GetAllDosen() ([]entity.HKIDosen, error) {
	return uc.hkiRepo.GetAllDosen()
}

func (uc *hkiUseCase) GetDosenByID(id int) (*entity.HKIDosen, error) {
	return uc.hkiRepo.GetDosenByID(id)
}

func (uc *hkiUseCase) CreateDosen(hki entity.HKIDosen) (*entity.HKIDosen, error) {
	return uc.hkiRepo.CreateDosen(hki)
}

func (uc *hkiUseCase) UpdateDosen(id int, hki entity.HKIDosen) (*entity.HKIDosen, error) {
	return uc.hkiRepo.UpdateDosen(id, hki)
}

func (uc *hkiUseCase) DeleteDosen(id int) error {
	return uc.hkiRepo.DeleteDosen(id)
}

// Mahasiswa implementations
func (uc *hkiUseCase) GetAllMhs() ([]entity.HKIMhs, error) {
	return uc.hkiRepo.GetAllMhs()
}

func (uc *hkiUseCase) GetMhsByID(id int) (*entity.HKIMhs, error) {
	return uc.hkiRepo.GetMhsByID(id)
}

func (uc *hkiUseCase) CreateMhs(hki entity.HKIMhs) (*entity.HKIMhs, error) {
	return uc.hkiRepo.CreateMhs(hki)
}

func (uc *hkiUseCase) UpdateMhs(id int, hki entity.HKIMhs) (*entity.HKIMhs, error) {
	return uc.hkiRepo.UpdateMhs(id, hki)
}

func (uc *hkiUseCase) DeleteMhs(id int) error {
	return uc.hkiRepo.DeleteMhs(id)
} 