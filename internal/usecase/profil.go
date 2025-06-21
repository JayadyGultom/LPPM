package usecase

import (
	"perpustakaan/internal/entity"
	"perpustakaan/internal/repository"
)

type ProfilUseCase interface {
	GetAll() ([]entity.Profil, error)
	GetByID(id int) (*entity.Profil, error)
	Create(profil entity.Profil) (*entity.Profil, error)
	Update(id int, profil entity.Profil) (*entity.Profil, error)
	Delete(id int) error
}

type VisiMisiUseCase interface {
	GetAll() ([]entity.VisiMisi, error)
	GetByID(id int) (*entity.VisiMisi, error)
}

type StrukturOrganisasiUseCase interface {
	GetAll() ([]entity.StrukturOrganisasi, error)
	GetByID(id int) (*entity.StrukturOrganisasi, error)
}

type profilUseCase struct {
	profilRepo repository.ProfilRepository
}

type visimisiUseCase struct {
	visimisiRepo repository.VisiMisiRepository
}

type strukturOrganisasiUseCase struct {
	strukturRepo repository.StrukturOrganisasiRepository
}

func NewProfilUseCase(profilRepo repository.ProfilRepository) ProfilUseCase {
	return &profilUseCase{profilRepo: profilRepo}
}

func NewVisiMisiUseCase(visimisiRepo repository.VisiMisiRepository) VisiMisiUseCase {
	return &visimisiUseCase{visimisiRepo: visimisiRepo}
}

func NewStrukturOrganisasiUseCase(strukturRepo repository.StrukturOrganisasiRepository) StrukturOrganisasiUseCase {
	return &strukturOrganisasiUseCase{strukturRepo: strukturRepo}
}

// Profil implementations
func (uc *profilUseCase) GetAll() ([]entity.Profil, error) {
	return uc.profilRepo.GetAll()
}

func (uc *profilUseCase) GetByID(id int) (*entity.Profil, error) {
	return uc.profilRepo.GetByID(id)
}

func (uc *profilUseCase) Create(profil entity.Profil) (*entity.Profil, error) {
	return uc.profilRepo.Create(profil)
}

func (uc *profilUseCase) Update(id int, profil entity.Profil) (*entity.Profil, error) {
	return uc.profilRepo.Update(id, profil)
}

func (uc *profilUseCase) Delete(id int) error {
	return uc.profilRepo.Delete(id)
}

// VisiMisi implementations
func (uc *visimisiUseCase) GetAll() ([]entity.VisiMisi, error) {
	return uc.visimisiRepo.GetAll()
}

func (uc *visimisiUseCase) GetByID(id int) (*entity.VisiMisi, error) {
	return uc.visimisiRepo.GetByID(id)
}

// StrukturOrganisasi implementations
func (uc *strukturOrganisasiUseCase) GetAll() ([]entity.StrukturOrganisasi, error) {
	return uc.strukturRepo.GetAll()
}

func (uc *strukturOrganisasiUseCase) GetByID(id int) (*entity.StrukturOrganisasi, error) {
	return uc.strukturRepo.GetByID(id)
} 