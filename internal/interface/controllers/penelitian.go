package controllers

import (
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PenelitianController struct {
	penelitianUseCase usecase.PenelitianUseCase
}

func NewPenelitianController(penelitianUseCase usecase.PenelitianUseCase) *PenelitianController {
	return &PenelitianController{
		penelitianUseCase: penelitianUseCase,
	}
}

// Bame handlers
func (c *PenelitianController) GetAllBame(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllBame()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetBameByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetBameByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateBame(ctx *gin.Context) {
	var penelitian entity.PenelitianBame
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateBame(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateBame(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianBame
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateBame(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteBame(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteBame(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Bame deleted successfully"})
}

// Hpp handlers
func (c *PenelitianController) GetAllHpp(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllHpp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetHppByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetHppByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateHpp(ctx *gin.Context) {
	var penelitian entity.PenelitianHpp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateHpp(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateHpp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianHpp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateHpp(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteHpp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteHpp(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Hpp deleted successfully"})
}

// Lp handlers
func (c *PenelitianController) GetAllLp(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllLp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetLpByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetLpByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateLp(ctx *gin.Context) {
	var penelitian entity.PenelitianLp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateLp(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateLp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianLp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateLp(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteLp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteLp(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Lp deleted successfully"})
}

// P3 handlers
func (c *PenelitianController) GetAllP3(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllP3()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetP3ByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetP3ByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateP3(ctx *gin.Context) {
	var penelitian entity.PenelitianP3
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateP3(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateP3(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianP3
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateP3(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteP3(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteP3(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian P3 deleted successfully"})
}

// Rrp handlers
func (c *PenelitianController) GetAllRrp(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllRrp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetRrpByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetRrpByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateRrp(ctx *gin.Context) {
	var penelitian entity.PenelitianRrp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateRrp(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateRrp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianRrp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateRrp(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteRrp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteRrp(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Rrp deleted successfully"})
}

// SkR handlers
func (c *PenelitianController) GetAllSkR(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllSkR()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetSkRByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetSkRByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateSkR(ctx *gin.Context) {
	var penelitian entity.PenelitianSkR
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateSkR(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateSkR(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianSkR
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateSkR(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteSkR(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteSkR(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian SkR deleted successfully"})
}

// Stp handlers
func (c *PenelitianController) GetAllStp(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllStp()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetStpByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetStpByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateStp(ctx *gin.Context) {
	var penelitian entity.PenelitianStp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateStp(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateStp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianStp
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateStp(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteStp(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteStp(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Stp deleted successfully"})
}

// Tcr handlers
func (c *PenelitianController) GetAllTcr(ctx *gin.Context) {
	penelitian, err := c.penelitianUseCase.GetAllTcr()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) GetTcrByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	penelitian, err := c.penelitianUseCase.GetTcrByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, penelitian)
}

func (c *PenelitianController) CreateTcr(ctx *gin.Context) {
	var penelitian entity.PenelitianTcr
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.penelitianUseCase.CreateTcr(penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PenelitianController) UpdateTcr(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var penelitian entity.PenelitianTcr
	if err := ctx.ShouldBindJSON(&penelitian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.penelitianUseCase.UpdateTcr(id, penelitian)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *PenelitianController) DeleteTcr(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.penelitianUseCase.DeleteTcr(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penelitian Tcr deleted successfully"})
} 