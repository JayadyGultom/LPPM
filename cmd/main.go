package main

import (
	"log"

	"perpustakaan/config"
	"perpustakaan/internal/interface/routes"
)

func main() {
	// Initialize database
	config.InitDB()

	// Setup routes with Gin
	r := routes.SetupRoutes()

	// Start server
	log.Println("🚀 LPPM REST API Server Started")
	log.Println("📍 Server running at: http://localhost:8081")
	log.Println("🌐 Base URL: http://localhost:8081")
	log.Println("📖 API Documentation: http://localhost:8081/")
	log.Println("❤️  Health Check: http://localhost:8081/health")
	log.Println("")
	log.Println("📋 AVAILABLE ENDPOINTS:")
	log.Println("")
	
	// Profile endpoints (Read-only)
	log.Println("🏛️  PROFILE ENDPOINTS (READ-ONLY):")
	log.Println("   GET http://localhost:8081/visimisi")
	log.Println("   GET http://localhost:8081/visimisi/{id}")
	log.Println("   GET http://localhost:8081/struktur-organisasi")
	log.Println("   GET http://localhost:8081/struktur-organisasi/{id}")
	log.Println("")
	
	// Profil endpoints (Full CRUD)
	log.Println("👤 PROFIL ENDPOINTS (FULL CRUD):")
	log.Println("   GET    http://localhost:8081/profil")
	log.Println("   GET    http://localhost:8081/profil/{id}")
	log.Println("   POST   http://localhost:8081/profil")
	log.Println("   PUT    http://localhost:8081/profil/{id}")
	log.Println("   DELETE http://localhost:8081/profil/{id}")
	log.Println("")
	
	// Research endpoints (GET only for display)
	log.Println("🔬 RESEARCH ENDPOINTS (FULL CRUD):")
	log.Println("   GET http://localhost:8081/penelitian/bame")
	log.Println("   GET http://localhost:8081/penelitian/hpp")
	log.Println("   GET http://localhost:8081/penelitian/lp")
	log.Println("   GET http://localhost:8081/penelitian/p3")
	log.Println("   GET http://localhost:8081/penelitian/rrp")
	log.Println("   GET http://localhost:8081/penelitian/skr")
	log.Println("   GET http://localhost:8081/penelitian/stp")
	log.Println("   GET http://localhost:8081/penelitian/tcr")
	log.Println("")
	
	// Community Service endpoints (GET only for display)
	log.Println("🤝 COMMUNITY SERVICE ENDPOINTS (FULL CRUD):")
	log.Println("   GET http://localhost:8081/pkm/bame")
	log.Println("   GET http://localhost:8081/pkm/hpp")
	log.Println("   GET http://localhost:8081/pkm/lp")
	log.Println("   GET http://localhost:8081/pkm/p3")
	log.Println("   GET http://localhost:8081/pkm/rrp")
	log.Println("   GET http://localhost:8081/pkm/skr")
	log.Println("   GET http://localhost:8081/pkm/stp")
	log.Println("   GET http://localhost:8081/pkm/tcr")
	log.Println("")
	
	// Journal endpoints (GET only for display)
	log.Println("📚 JOURNAL ENDPOINTS (FULL CRUD):")
	log.Println("   GET http://localhost:8081/jurnal/jk")
	log.Println("   GET http://localhost:8081/jurnal/js")
	log.Println("   GET http://localhost:8081/jurnal/kiat")
	log.Println("   GET http://localhost:8081/jurnal/tajb")
	log.Println("   GET http://localhost:8081/jurnal/teknois")
	log.Println("   GET http://localhost:8081/jurnal/tmjb")
	log.Println("")
	
	// HKI endpoints (GET only for display)
	log.Println("💡 INTELLECTUAL PROPERTY ENDPOINTS (FULL CRUD):")
	log.Println("   GET http://localhost:8081/hki/dosen")
	log.Println("   GET http://localhost:8081/hki/mhs")
	log.Println("")
	
	log.Println("📝 NOTE: All endpoints support GET, POST, PUT, DELETE methods")
	log.Println("   except Profile endpoints (READ-ONLY)")
	log.Println("")
	log.Println("🔧 CORS enabled for cross-origin requests")
	log.Println("")
	log.Println("⏹️  Press Ctrl+C to stop the server")
	log.Println("")
	
	log.Fatal(r.Run(":8081"))
}
