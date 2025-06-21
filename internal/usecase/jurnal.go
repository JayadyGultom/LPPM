package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type JurnalUseCase interface {
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

type jurnalUseCase struct {
	jurnalRepo repository.JurnalRepository
}

func NewJurnalUseCase(jurnalRepo repository.JurnalRepository) JurnalUseCase {
	return &jurnalUseCase{
		jurnalRepo: jurnalRepo,
	}
}

// Jk implementations
func (uc *jurnalUseCase) GetAllJk() ([]entity.JurnalJk, error) {
	return uc.jurnalRepo.GetAllJk()
}

func (uc *jurnalUseCase) GetJkByID(id int) (*entity.JurnalJk, error) {
	return uc.jurnalRepo.GetJkByID(id)
}

func (uc *jurnalUseCase) CreateJk(jurnal entity.JurnalJk) (*entity.JurnalJk, error) {
	return uc.jurnalRepo.CreateJk(jurnal)
}

func (uc *jurnalUseCase) UpdateJk(id int, jurnal entity.JurnalJk) (*entity.JurnalJk, error) {
	return uc.jurnalRepo.UpdateJk(id, jurnal)
}

func (uc *jurnalUseCase) DeleteJk(id int) error {
	return uc.jurnalRepo.DeleteJk(id)
}

// Js implementations
func (uc *jurnalUseCase) GetAllJs() ([]entity.JurnalJs, error) {
	return uc.jurnalRepo.GetAllJs()
}

func (uc *jurnalUseCase) GetJsByID(id int) (*entity.JurnalJs, error) {
	return uc.jurnalRepo.GetJsByID(id)
}

func (uc *jurnalUseCase) CreateJs(jurnal entity.JurnalJs) (*entity.JurnalJs, error) {
	return uc.jurnalRepo.CreateJs(jurnal)
}

func (uc *jurnalUseCase) UpdateJs(id int, jurnal entity.JurnalJs) (*entity.JurnalJs, error) {
	return uc.jurnalRepo.UpdateJs(id, jurnal)
}

func (uc *jurnalUseCase) DeleteJs(id int) error {
	return uc.jurnalRepo.DeleteJs(id)
}

// Kiat implementations
func (uc *jurnalUseCase) GetAllKiat() ([]entity.JurnalKiat, error) {
	return uc.jurnalRepo.GetAllKiat()
}

func (uc *jurnalUseCase) GetKiatByID(id int) (*entity.JurnalKiat, error) {
	return uc.jurnalRepo.GetKiatByID(id)
}

func (uc *jurnalUseCase) CreateKiat(jurnal entity.JurnalKiat) (*entity.JurnalKiat, error) {
	return uc.jurnalRepo.CreateKiat(jurnal)
}

func (uc *jurnalUseCase) UpdateKiat(id int, jurnal entity.JurnalKiat) (*entity.JurnalKiat, error) {
	return uc.jurnalRepo.UpdateKiat(id, jurnal)
}

func (uc *jurnalUseCase) DeleteKiat(id int) error {
	return uc.jurnalRepo.DeleteKiat(id)
}

// Tajb implementations
func (uc *jurnalUseCase) GetAllTajb() ([]entity.JurnalTajb, error) {
	return uc.jurnalRepo.GetAllTajb()
}

func (uc *jurnalUseCase) GetTajbByID(id int) (*entity.JurnalTajb, error) {
	return uc.jurnalRepo.GetTajbByID(id)
}

func (uc *jurnalUseCase) CreateTajb(jurnal entity.JurnalTajb) (*entity.JurnalTajb, error) {
	return uc.jurnalRepo.CreateTajb(jurnal)
}

func (uc *jurnalUseCase) UpdateTajb(id int, jurnal entity.JurnalTajb) (*entity.JurnalTajb, error) {
	return uc.jurnalRepo.UpdateTajb(id, jurnal)
}

func (uc *jurnalUseCase) DeleteTajb(id int) error {
	return uc.jurnalRepo.DeleteTajb(id)
}

// Teknois implementations
func (uc *jurnalUseCase) GetAllTeknois() ([]entity.JurnalTeknois, error) {
	return uc.jurnalRepo.GetAllTeknois()
}

func (uc *jurnalUseCase) GetTeknoisByID(id int) (*entity.JurnalTeknois, error) {
	return uc.jurnalRepo.GetTeknoisByID(id)
}

func (uc *jurnalUseCase) CreateTeknois(jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error) {
	return uc.jurnalRepo.CreateTeknois(jurnal)
}

func (uc *jurnalUseCase) UpdateTeknois(id int, jurnal entity.JurnalTeknois) (*entity.JurnalTeknois, error) {
	return uc.jurnalRepo.UpdateTeknois(id, jurnal)
}

func (uc *jurnalUseCase) DeleteTeknois(id int) error {
	return uc.jurnalRepo.DeleteTeknois(id)
}

// Tmjb implementations
func (uc *jurnalUseCase) GetAllTmjb() ([]entity.JurnalTmjb, error) {
	return uc.jurnalRepo.GetAllTmjb()
}

func (uc *jurnalUseCase) GetTmjbByID(id int) (*entity.JurnalTmjb, error) {
	return uc.jurnalRepo.GetTmjbByID(id)
}

func (uc *jurnalUseCase) CreateTmjb(jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error) {
	return uc.jurnalRepo.CreateTmjb(jurnal)
}

func (uc *jurnalUseCase) UpdateTmjb(id int, jurnal entity.JurnalTmjb) (*entity.JurnalTmjb, error) {
	return uc.jurnalRepo.UpdateTmjb(id, jurnal)
}

func (uc *jurnalUseCase) DeleteTmjb(id int) error {
	return uc.jurnalRepo.DeleteTmjb(id)
} 