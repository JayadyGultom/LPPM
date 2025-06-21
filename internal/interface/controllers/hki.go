package controllers

import (
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HKIController struct {
	hkiUseCase usecase.HKIUseCase
}

func NewHKIController(hkiUseCase usecase.HKIUseCase) *HKIController {
	return &HKIController{
		hkiUseCase: hkiUseCase,
	}
}

// Dosen handlers
func (c *HKIController) GetAllDosen(ctx *gin.Context) {
	hki, err := c.hkiUseCase.GetAllDosen()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hki)
}

func (c *HKIController) GetDosenByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	hki, err := c.hkiUseCase.GetDosenByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hki)
}

func (c *HKIController) CreateDosen(ctx *gin.Context) {
	var hki entity.HKIDosen
	if err := ctx.ShouldBindJSON(&hki); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.hkiUseCase.CreateDosen(hki)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *HKIController) UpdateDosen(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var hki entity.HKIDosen
	if err := ctx.ShouldBindJSON(&hki); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.hkiUseCase.UpdateDosen(id, hki)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *HKIController) DeleteDosen(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.hkiUseCase.DeleteDosen(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "HKI Dosen deleted successfully"})
}

// Mahasiswa handlers
func (c *HKIController) GetAllMhs(ctx *gin.Context) {
	hki, err := c.hkiUseCase.GetAllMhs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hki)
}

func (c *HKIController) GetMhsByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	hki, err := c.hkiUseCase.GetMhsByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hki)
}

func (c *HKIController) CreateMhs(ctx *gin.Context) {
	var hki entity.HKIMhs
	if err := ctx.ShouldBindJSON(&hki); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.hkiUseCase.CreateMhs(hki)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *HKIController) UpdateMhs(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var hki entity.HKIMhs
	if err := ctx.ShouldBindJSON(&hki); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.hkiUseCase.UpdateMhs(id, hki)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *HKIController) DeleteMhs(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.hkiUseCase.DeleteMhs(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "HKI Mahasiswa deleted successfully"})
} 