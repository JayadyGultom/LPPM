package controllers

import (
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfilController struct {
	profilUC usecase.ProfilUseCase
}

type VisiMisiController struct {
	visimisiUC usecase.VisiMisiUseCase
}

type StrukturOrganisasiController struct {
	strukturUC usecase.StrukturOrganisasiUseCase
}

func NewProfilController(profilUC usecase.ProfilUseCase) *ProfilController {
	return &ProfilController{profilUC: profilUC}
}

func NewVisiMisiController(visimisiUC usecase.VisiMisiUseCase) *VisiMisiController {
	return &VisiMisiController{visimisiUC: visimisiUC}
}

func NewStrukturOrganisasiController(strukturUC usecase.StrukturOrganisasiUseCase) *StrukturOrganisasiController {
	return &StrukturOrganisasiController{strukturUC: strukturUC}
}

// Profil handlers
func (c *ProfilController) GetAll(ctx *gin.Context) {
	profil, err := c.profilUC.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profil)
}

func (c *ProfilController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	profil, err := c.profilUC.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Profil not found"})
		return
	}
	ctx.JSON(http.StatusOK, profil)
}

func (c *ProfilController) Create(ctx *gin.Context) {
	var profil entity.Profil
	if err := ctx.ShouldBindJSON(&profil); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProfil, err := c.profilUC.Create(profil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdProfil)
}

func (c *ProfilController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var profil entity.Profil
	if err := ctx.ShouldBindJSON(&profil); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProfil, err := c.profilUC.Update(id, profil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedProfil)
}

func (c *ProfilController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.profilUC.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profil deleted successfully"})
}

// VisiMisi handlers
func (c *VisiMisiController) GetAll(ctx *gin.Context) {
	visimisi, err := c.visimisiUC.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, visimisi)
}

func (c *VisiMisiController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	visimisi, err := c.visimisiUC.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "VisiMisi not found"})
		return
	}
	ctx.JSON(http.StatusOK, visimisi)
}

// StrukturOrganisasi handlers
func (c *StrukturOrganisasiController) GetAll(ctx *gin.Context) {
	struktur, err := c.strukturUC.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, struktur)
}

func (c *StrukturOrganisasiController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	struktur, err := c.strukturUC.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "StrukturOrganisasi not found"})
		return
	}
	ctx.JSON(http.StatusOK, struktur)
} 