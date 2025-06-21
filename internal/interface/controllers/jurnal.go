package controllers

import (
	"net/http"
	"perpustakaan/internal/entity"
	"perpustakaan/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JurnalController struct {
	jurnalUseCase usecase.JurnalUseCase
}

func NewJurnalController(jurnalUseCase usecase.JurnalUseCase) *JurnalController {
	return &JurnalController{
		jurnalUseCase: jurnalUseCase,
	}
}

// Jk handlers
func (c *JurnalController) GetAllJk(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllJk()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetJkByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetJkByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateJk(ctx *gin.Context) {
	var jurnal entity.JurnalJk
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateJk(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateJk(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalJk
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateJk(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteJk(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteJk(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Jk deleted successfully"})
}

// Js handlers
func (c *JurnalController) GetAllJs(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllJs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetJsByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetJsByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateJs(ctx *gin.Context) {
	var jurnal entity.JurnalJs
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateJs(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateJs(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalJs
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateJs(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteJs(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteJs(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Js deleted successfully"})
}

// Kiat handlers
func (c *JurnalController) GetAllKiat(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllKiat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetKiatByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetKiatByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateKiat(ctx *gin.Context) {
	var jurnal entity.JurnalKiat
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateKiat(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateKiat(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalKiat
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateKiat(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteKiat(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteKiat(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Kiat deleted successfully"})
}

// Tajb handlers
func (c *JurnalController) GetAllTajb(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllTajb()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetTajbByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetTajbByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateTajb(ctx *gin.Context) {
	var jurnal entity.JurnalTajb
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateTajb(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateTajb(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalTajb
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateTajb(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteTajb(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteTajb(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Tajb deleted successfully"})
}

// Teknois handlers
func (c *JurnalController) GetAllTeknois(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllTeknois()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetTeknoisByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetTeknoisByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateTeknois(ctx *gin.Context) {
	var jurnal entity.JurnalTeknois
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateTeknois(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateTeknois(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalTeknois
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateTeknois(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteTeknois(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteTeknois(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Teknois deleted successfully"})
}

// Tmjb handlers
func (c *JurnalController) GetAllTmjb(ctx *gin.Context) {
	jurnal, err := c.jurnalUseCase.GetAllTmjb()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) GetTmjbByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jurnal, err := c.jurnalUseCase.GetTmjbByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, jurnal)
}

func (c *JurnalController) CreateTmjb(ctx *gin.Context) {
	var jurnal entity.JurnalTmjb
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := c.jurnalUseCase.CreateTmjb(jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *JurnalController) UpdateTmjb(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jurnal entity.JurnalTmjb
	if err := ctx.ShouldBindJSON(&jurnal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := c.jurnalUseCase.UpdateTmjb(id, jurnal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *JurnalController) DeleteTmjb(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.jurnalUseCase.DeleteTmjb(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Jurnal Tmjb deleted successfully"})
} 