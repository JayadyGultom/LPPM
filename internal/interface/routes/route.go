package routes

import (
	"perpustakaan/config"
	"perpustakaan/internal/interface/controllers"
	"perpustakaan/internal/repository"
	"perpustakaan/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// Initialize database
	db := config.GetDB()

	// Initialize repositories
	profilRepo := repository.NewProfilRepository(db)
	visimisiRepo := repository.NewVisiMisiRepository(db)
	strukturRepo := repository.NewStrukturOrganisasiRepository(db)
	penelitianRepo := repository.NewPenelitianRepository(db)
	pkmRepo := repository.NewPKMRepository(db)
	jurnalRepo := repository.NewJurnalRepository(db)
	hkiRepo := repository.NewHKIRepository(db)

	// Initialize use cases
	profilUseCase := usecase.NewProfilUseCase(profilRepo)
	visimisiUseCase := usecase.NewVisiMisiUseCase(visimisiRepo)
	strukturUseCase := usecase.NewStrukturOrganisasiUseCase(strukturRepo)
	penelitianUseCase := usecase.NewPenelitianUseCase(penelitianRepo)
	pkmUseCase := usecase.NewPKMUseCase(pkmRepo)
	jurnalUseCase := usecase.NewJurnalUseCase(jurnalRepo)
	hkiUseCase := usecase.NewHKIUseCase(hkiRepo)

	// Initialize controllers
	profilController := controllers.NewProfilController(profilUseCase)
	visimisiController := controllers.NewVisiMisiController(visimisiUseCase)
	strukturController := controllers.NewStrukturOrganisasiController(strukturUseCase)
	penelitianController := controllers.NewPenelitianController(penelitianUseCase)
	pkmController := controllers.NewPKMController(pkmUseCase)
	jurnalController := controllers.NewJurnalController(jurnalUseCase)
	hkiController := controllers.NewHKIController(hkiUseCase)

	// Setup Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// Root route - API Documentation
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🚀 LPPM REST API Server",
			"version": "1.0.0",
			"status":  "running",
			"server":  "http://localhost:8081",
			"documentation": gin.H{
				"profile": gin.H{
					"read_only": []string{
						"GET /visimisi",
						"GET /visimisi/{id}",
						"GET /struktur-organisasi",
						"GET /struktur-organisasi/{id}",
					},
					"full_crud": []string{
						"GET /profil",
						"GET /profil/{id}",
						"POST /profil",
						"PUT /profil/{id}",
						"DELETE /profil/{id}",
					},
				},
				"research": []string{
					"GET /penelitian/bame",
					"GET /penelitian/hpp",
					"GET /penelitian/lp",
					"GET /penelitian/p3",
					"GET /penelitian/rrp",
					"GET /penelitian/skr",
					"GET /penelitian/stp",
					"GET /penelitian/tcr",
				},
				"community_service": []string{
					"GET /pkm/bame",
					"GET /pkm/hpp",
					"GET /pkm/lp",
					"GET /pkm/p3",
					"GET /pkm/rrp",
					"GET /pkm/skr",
					"GET /pkm/stp",
					"GET /pkm/tcr",
				},
				"journal": []string{
					"GET /jurnal/jk",
					"GET /jurnal/js",
					"GET /jurnal/kiat",
					"GET /jurnal/tajb",
					"GET /jurnal/teknois",
					"GET /jurnal/tmjb",
				},
				"intellectual_property": []string{
					"GET /hki/dosen",
					"GET /hki/mhs",
				},
			},
			"note": "All endpoints support GET, POST, PUT, DELETE methods except Profile endpoints (READ-ONLY)",
		})
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "LPPM API is running",
			"timestamp": "2025-06-21",
		})
	})

	// Profile routes (Read-only)
	visimisiGroup := r.Group("/visimisi")
	{
		visimisiGroup.GET("", visimisiController.GetAll)
		visimisiGroup.GET("/:id", visimisiController.GetByID)
	}

	strukturGroup := r.Group("/struktur-organisasi")
	{
		strukturGroup.GET("", strukturController.GetAll)
		strukturGroup.GET("/:id", strukturController.GetByID)
	}

	// Profil routes (Full CRUD)
	profilGroup := r.Group("/profil")
	{
		profilGroup.GET("", profilController.GetAll)
		profilGroup.GET("/:id", profilController.GetByID)
		profilGroup.POST("", profilController.Create)
		profilGroup.PUT("/:id", profilController.Update)
		profilGroup.DELETE("/:id", profilController.Delete)
	}

	// Penelitian routes
	setupPenelitianRoutes(r, penelitianController)
	
	// PKM routes
	setupPKMRoutes(r, pkmController)
	
	// Jurnal routes
	setupJurnalRoutes(r, jurnalController)
	
	// HKI routes
	setupHKIRoutes(r, hkiController)

	return r
}

func setupPenelitianRoutes(r *gin.Engine, controller *controllers.PenelitianController) {
	// Bame routes
	r.GET("/penelitian/bame", controller.GetAllBame)
	r.POST("/penelitian/bame", controller.CreateBame)
	r.GET("/penelitian/bame/:id", controller.GetBameByID)
	r.PUT("/penelitian/bame/:id", controller.UpdateBame)
	r.DELETE("/penelitian/bame/:id", controller.DeleteBame)

	// Hpp routes
	r.GET("/penelitian/hpp", controller.GetAllHpp)
	r.POST("/penelitian/hpp", controller.CreateHpp)
	r.GET("/penelitian/hpp/:id", controller.GetHppByID)
	r.PUT("/penelitian/hpp/:id", controller.UpdateHpp)
	r.DELETE("/penelitian/hpp/:id", controller.DeleteHpp)

	// Lp routes
	r.GET("/penelitian/lp", controller.GetAllLp)
	r.POST("/penelitian/lp", controller.CreateLp)
	r.GET("/penelitian/lp/:id", controller.GetLpByID)
	r.PUT("/penelitian/lp/:id", controller.UpdateLp)
	r.DELETE("/penelitian/lp/:id", controller.DeleteLp)

	// P3 routes
	r.GET("/penelitian/p3", controller.GetAllP3)
	r.POST("/penelitian/p3", controller.CreateP3)
	r.GET("/penelitian/p3/:id", controller.GetP3ByID)
	r.PUT("/penelitian/p3/:id", controller.UpdateP3)
	r.DELETE("/penelitian/p3/:id", controller.DeleteP3)

	// Rrp routes
	r.GET("/penelitian/rrp", controller.GetAllRrp)
	r.POST("/penelitian/rrp", controller.CreateRrp)
	r.GET("/penelitian/rrp/:id", controller.GetRrpByID)
	r.PUT("/penelitian/rrp/:id", controller.UpdateRrp)
	r.DELETE("/penelitian/rrp/:id", controller.DeleteRrp)

	// SkR routes
	r.GET("/penelitian/skr", controller.GetAllSkR)
	r.POST("/penelitian/skr", controller.CreateSkR)
	r.GET("/penelitian/skr/:id", controller.GetSkRByID)
	r.PUT("/penelitian/skr/:id", controller.UpdateSkR)
	r.DELETE("/penelitian/skr/:id", controller.DeleteSkR)

	// Stp routes
	r.GET("/penelitian/stp", controller.GetAllStp)
	r.POST("/penelitian/stp", controller.CreateStp)
	r.GET("/penelitian/stp/:id", controller.GetStpByID)
	r.PUT("/penelitian/stp/:id", controller.UpdateStp)
	r.DELETE("/penelitian/stp/:id", controller.DeleteStp)

	// Tcr routes
	r.GET("/penelitian/tcr", controller.GetAllTcr)
	r.POST("/penelitian/tcr", controller.CreateTcr)
	r.GET("/penelitian/tcr/:id", controller.GetTcrByID)
	r.PUT("/penelitian/tcr/:id", controller.UpdateTcr)
	r.DELETE("/penelitian/tcr/:id", controller.DeleteTcr)
}

func setupPKMRoutes(r *gin.Engine, controller *controllers.PKMController) {
	// Bame routes
	r.GET("/pkm/bame", controller.GetAllBame)
	r.POST("/pkm/bame", controller.CreateBame)
	r.GET("/pkm/bame/:id", controller.GetBameByID)
	r.PUT("/pkm/bame/:id", controller.UpdateBame)
	r.DELETE("/pkm/bame/:id", controller.DeleteBame)

	// Hpp routes
	r.GET("/pkm/hpp", controller.GetAllHpp)
	r.POST("/pkm/hpp", controller.CreateHpp)
	r.GET("/pkm/hpp/:id", controller.GetHppByID)
	r.PUT("/pkm/hpp/:id", controller.UpdateHpp)
	r.DELETE("/pkm/hpp/:id", controller.DeleteHpp)

	// Lp routes
	r.GET("/pkm/lp", controller.GetAllLp)
	r.POST("/pkm/lp", controller.CreateLp)
	r.GET("/pkm/lp/:id", controller.GetLpByID)
	r.PUT("/pkm/lp/:id", controller.UpdateLp)
	r.DELETE("/pkm/lp/:id", controller.DeleteLp)

	// P3 routes
	r.GET("/pkm/p3", controller.GetAllP3)
	r.POST("/pkm/p3", controller.CreateP3)
	r.GET("/pkm/p3/:id", controller.GetP3ByID)
	r.PUT("/pkm/p3/:id", controller.UpdateP3)
	r.DELETE("/pkm/p3/:id", controller.DeleteP3)

	// Rrp routes
	r.GET("/pkm/rrp", controller.GetAllRrp)
	r.POST("/pkm/rrp", controller.CreateRrp)
	r.GET("/pkm/rrp/:id", controller.GetRrpByID)
	r.PUT("/pkm/rrp/:id", controller.UpdateRrp)
	r.DELETE("/pkm/rrp/:id", controller.DeleteRrp)

	// SkR routes
	r.GET("/pkm/skr", controller.GetAllSkR)
	r.POST("/pkm/skr", controller.CreateSkR)
	r.GET("/pkm/skr/:id", controller.GetSkRByID)
	r.PUT("/pkm/skr/:id", controller.UpdateSkR)
	r.DELETE("/pkm/skr/:id", controller.DeleteSkR)

	// Stp routes
	r.GET("/pkm/stp", controller.GetAllStp)
	r.POST("/pkm/stp", controller.CreateStp)
	r.GET("/pkm/stp/:id", controller.GetStpByID)
	r.PUT("/pkm/stp/:id", controller.UpdateStp)
	r.DELETE("/pkm/stp/:id", controller.DeleteStp)

	// Tcr routes
	r.GET("/pkm/tcr", controller.GetAllTcr)
	r.POST("/pkm/tcr", controller.CreateTcr)
	r.GET("/pkm/tcr/:id", controller.GetTcrByID)
	r.PUT("/pkm/tcr/:id", controller.UpdateTcr)
	r.DELETE("/pkm/tcr/:id", controller.DeleteTcr)
}

func setupJurnalRoutes(r *gin.Engine, controller *controllers.JurnalController) {
	// Jk routes
	r.GET("/jurnal/jk", controller.GetAllJk)
	r.POST("/jurnal/jk", controller.CreateJk)
	r.GET("/jurnal/jk/:id", controller.GetJkByID)
	r.PUT("/jurnal/jk/:id", controller.UpdateJk)
	r.DELETE("/jurnal/jk/:id", controller.DeleteJk)

	// Js routes
	r.GET("/jurnal/js", controller.GetAllJs)
	r.POST("/jurnal/js", controller.CreateJs)
	r.GET("/jurnal/js/:id", controller.GetJsByID)
	r.PUT("/jurnal/js/:id", controller.UpdateJs)
	r.DELETE("/jurnal/js/:id", controller.DeleteJs)

	// Kiat routes
	r.GET("/jurnal/kiat", controller.GetAllKiat)
	r.POST("/jurnal/kiat", controller.CreateKiat)
	r.GET("/jurnal/kiat/:id", controller.GetKiatByID)
	r.PUT("/jurnal/kiat/:id", controller.UpdateKiat)
	r.DELETE("/jurnal/kiat/:id", controller.DeleteKiat)

	// Tajb routes
	r.GET("/jurnal/tajb", controller.GetAllTajb)
	r.POST("/jurnal/tajb", controller.CreateTajb)
	r.GET("/jurnal/tajb/:id", controller.GetTajbByID)
	r.PUT("/jurnal/tajb/:id", controller.UpdateTajb)
	r.DELETE("/jurnal/tajb/:id", controller.DeleteTajb)

	// Teknois routes
	r.GET("/jurnal/teknois", controller.GetAllTeknois)
	r.POST("/jurnal/teknois", controller.CreateTeknois)
	r.GET("/jurnal/teknois/:id", controller.GetTeknoisByID)
	r.PUT("/jurnal/teknois/:id", controller.UpdateTeknois)
	r.DELETE("/jurnal/teknois/:id", controller.DeleteTeknois)

	// Tmjb routes
	r.GET("/jurnal/tmjb", controller.GetAllTmjb)
	r.POST("/jurnal/tmjb", controller.CreateTmjb)
	r.GET("/jurnal/tmjb/:id", controller.GetTmjbByID)
	r.PUT("/jurnal/tmjb/:id", controller.UpdateTmjb)
	r.DELETE("/jurnal/tmjb/:id", controller.DeleteTmjb)
}

func setupHKIRoutes(r *gin.Engine, controller *controllers.HKIController) {
	// Dosen routes
	r.GET("/hki/dosen", controller.GetAllDosen)
	r.POST("/hki/dosen", controller.CreateDosen)
	r.GET("/hki/dosen/:id", controller.GetDosenByID)
	r.PUT("/hki/dosen/:id", controller.UpdateDosen)
	r.DELETE("/hki/dosen/:id", controller.DeleteDosen)

	// Mahasiswa routes
	r.GET("/hki/mhs", controller.GetAllMhs)
	r.POST("/hki/mhs", controller.CreateMhs)
	r.GET("/hki/mhs/:id", controller.GetMhsByID)
	r.PUT("/hki/mhs/:id", controller.UpdateMhs)
	r.DELETE("/hki/mhs/:id", controller.DeleteMhs)
} 