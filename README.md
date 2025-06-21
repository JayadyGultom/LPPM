# LPPM REST API

A RESTful API for managing LPPM (Lembaga Penelitian dan Pengabdian kepada Masyarakat) data built with Go, following clean architecture principles.

## Features

- **Profile Management**: Read-only access to vision, mission, and organizational structure
- **Research Management**: Full CRUD operations for various research categories
- **Community Service Management**: Full CRUD operations for community service activities
- **Journal Management**: Full CRUD operations for various journal types
- **Intellectual Property Management**: Full CRUD operations for HKI (Hak Kekayaan Intelektual)

## Database Schema

The API is built around the LPPM database with the following main categories:

### Profile Tables (Read-Only)
- `lppm_profile_visimisi` - Vision and Mission
- `lppm_profile_so` - Organizational Structure

### Research Tables (Full CRUD)
- `lppm_penelitian_bame` - BAME Research
- `lppm_penelitian_hpp` - HPP Research
- `lppm_penelitian_lp` - LP Research
- `lppm_penelitian_p3` - P3 Research
- `lppm_penelitian_rrp` - RRP Research
- `lppm_penelitian_sk_r` - SK-R Research
- `lppm_penelitian_stp` - STP Research
- `lppm_penelitian_tcr` - TCR Research

### Community Service Tables (Full CRUD)
- `lppm_pkm_bame` - BAME Community Service
- `lppm_pkm_hpp` - HPP Community Service
- `lppm_pkm_lp` - LP Community Service
- `lppm_pkm_p3` - P3 Community Service
- `lppm_pkm_rrp` - RRP Community Service
- `lppm_pkm_sk_r` - SK-R Community Service
- `lppm_pkm_stp` - STP Community Service
- `lppm_pkm_tcr` - TCR Community Service

### Journal Tables (Full CRUD)
- `lppm_jurnal_jk` - JK Journal
- `lppm_jurnal_js` - JS Journal
- `lppm_jurnal_kiat` - KIAT Journal
- `lppm_jurnal_tajb` - TAJB Journal
- `lppm_jurnal_teknois` - TEKNOIS Journal
- `lppm_jurnal_tmjb` - TMJB Journal

### Intellectual Property Tables (Full CRUD)
- `lppm_hki_dosen` - Lecturer HKI
- `lppm_hki_mhs` - Student HKI

## API Endpoints

### Profile Endpoints (Read-Only)

#### Vision & Mission
- `GET /visimisi` - Get all vision and mission data
- `GET /visimisi/{id}` - Get specific vision and mission by ID

#### Organizational Structure
- `GET /struktur-organisasi` - Get all organizational structure data
- `GET /struktur-organisasi/{id}` - Get specific organizational structure by ID

### Research Endpoints (Full CRUD)

#### BAME Research
- `GET /penelitian/bame` - Get all BAME research data
- `POST /penelitian/bame` - Create new BAME research
- `GET /penelitian/bame/{id}` - Get specific BAME research by ID
- `PUT /penelitian/bame/{id}` - Update BAME research by ID
- `DELETE /penelitian/bame/{id}` - Delete BAME research by ID

#### HPP Research
- `GET /penelitian/hpp` - Get all HPP research data
- `POST /penelitian/hpp` - Create new HPP research
- `GET /penelitian/hpp/{id}` - Get specific HPP research by ID
- `PUT /penelitian/hpp/{id}` - Update HPP research by ID
- `DELETE /penelitian/hpp/{id}` - Delete HPP research by ID

#### LP Research
- `GET /penelitian/lp` - Get all LP research data
- `POST /penelitian/lp` - Create new LP research
- `GET /penelitian/lp/{id}` - Get specific LP research by ID
- `PUT /penelitian/lp/{id}` - Update LP research by ID
- `DELETE /penelitian/lp/{id}` - Delete LP research by ID

#### P3 Research
- `GET /penelitian/p3` - Get all P3 research data
- `POST /penelitian/p3` - Create new P3 research
- `GET /penelitian/p3/{id}` - Get specific P3 research by ID
- `PUT /penelitian/p3/{id}` - Update P3 research by ID
- `DELETE /penelitian/p3/{id}` - Delete P3 research by ID

#### RRP Research
- `GET /penelitian/rrp` - Get all RRP research data
- `POST /penelitian/rrp` - Create new RRP research
- `GET /penelitian/rrp/{id}` - Get specific RRP research by ID
- `PUT /penelitian/rrp/{id}` - Update RRP research by ID
- `DELETE /penelitian/rrp/{id}` - Delete RRP research by ID

#### SK-R Research
- `GET /penelitian/skr` - Get all SK-R research data
- `POST /penelitian/skr` - Create new SK-R research
- `GET /penelitian/skr/{id}` - Get specific SK-R research by ID
- `PUT /penelitian/skr/{id}` - Update SK-R research by ID
- `DELETE /penelitian/skr/{id}` - Delete SK-R research by ID

#### STP Research
- `GET /penelitian/stp` - Get all STP research data
- `POST /penelitian/stp` - Create new STP research
- `GET /penelitian/stp/{id}` - Get specific STP research by ID
- `PUT /penelitian/stp/{id}` - Update STP research by ID
- `DELETE /penelitian/stp/{id}` - Delete STP research by ID

#### TCR Research
- `GET /penelitian/tcr` - Get all TCR research data
- `POST /penelitian/tcr` - Create new TCR research
- `GET /penelitian/tcr/{id}` - Get specific TCR research by ID
- `PUT /penelitian/tcr/{id}` - Update TCR research by ID
- `DELETE /penelitian/tcr/{id}` - Delete TCR research by ID

### Community Service Endpoints (Full CRUD)

#### BAME Community Service
- `GET /pkm/bame` - Get all BAME community service data
- `POST /pkm/bame` - Create new BAME community service
- `GET /pkm/bame/{id}` - Get specific BAME community service by ID
- `PUT /pkm/bame/{id}` - Update BAME community service by ID
- `DELETE /pkm/bame/{id}` - Delete BAME community service by ID

#### HPP Community Service
- `GET /pkm/hpp` - Get all HPP community service data
- `POST /pkm/hpp` - Create new HPP community service
- `GET /pkm/hpp/{id}` - Get specific HPP community service by ID
- `PUT /pkm/hpp/{id}` - Update HPP community service by ID
- `DELETE /pkm/hpp/{id}` - Delete HPP community service by ID

#### LP Community Service
- `GET /pkm/lp` - Get all LP community service data
- `POST /pkm/lp` - Create new LP community service
- `GET /pkm/lp/{id}` - Get specific LP community service by ID
- `PUT /pkm/lp/{id}` - Update LP community service by ID
- `DELETE /pkm/lp/{id}` - Delete LP community service by ID

#### P3 Community Service
- `GET /pkm/p3` - Get all P3 community service data
- `POST /pkm/p3` - Create new P3 community service
- `GET /pkm/p3/{id}` - Get specific P3 community service by ID
- `PUT /pkm/p3/{id}` - Update P3 community service by ID
- `DELETE /pkm/p3/{id}` - Delete P3 community service by ID

#### RRP Community Service
- `GET /pkm/rrp` - Get all RRP community service data
- `POST /pkm/rrp` - Create new RRP community service
- `GET /pkm/rrp/{id}` - Get specific RRP community service by ID
- `PUT /pkm/rrp/{id}` - Update RRP community service by ID
- `DELETE /pkm/rrp/{id}` - Delete RRP community service by ID

#### SK-R Community Service
- `GET /pkm/skr` - Get all SK-R community service data
- `POST /pkm/skr` - Create new SK-R community service
- `GET /pkm/skr/{id}` - Get specific SK-R community service by ID
- `PUT /pkm/skr/{id}` - Update SK-R community service by ID
- `DELETE /pkm/skr/{id}` - Delete SK-R community service by ID

#### STP Community Service
- `GET /pkm/stp` - Get all STP community service data
- `POST /pkm/stp` - Create new STP community service
- `GET /pkm/stp/{id}` - Get specific STP community service by ID
- `PUT /pkm/stp/{id}` - Update STP community service by ID
- `DELETE /pkm/stp/{id}` - Delete STP community service by ID

#### TCR Community Service
- `GET /pkm/tcr` - Get all TCR community service data
- `POST /pkm/tcr` - Create new TCR community service
- `GET /pkm/tcr/{id}` - Get specific TCR community service by ID
- `PUT /pkm/tcr/{id}` - Update TCR community service by ID
- `DELETE /pkm/tcr/{id}` - Delete TCR community service by ID

### Journal Endpoints (Full CRUD)

#### JK Journal
- `GET /jurnal/jk` - Get all JK journal data
- `POST /jurnal/jk` - Create new JK journal
- `GET /jurnal/jk/{id}` - Get specific JK journal by ID
- `PUT /jurnal/jk/{id}` - Update JK journal by ID
- `DELETE /jurnal/jk/{id}` - Delete JK journal by ID

#### JS Journal
- `GET /jurnal/js` - Get all JS journal data
- `POST /jurnal/js` - Create new JS journal
- `GET /jurnal/js/{id}` - Get specific JS journal by ID
- `PUT /jurnal/js/{id}` - Update JS journal by ID
- `DELETE /jurnal/js/{id}` - Delete JS journal by ID

#### KIAT Journal
- `GET /jurnal/kiat` - Get all KIAT journal data
- `POST /jurnal/kiat` - Create new KIAT journal
- `GET /jurnal/kiat/{id}` - Get specific KIAT journal by ID
- `PUT /jurnal/kiat/{id}` - Update KIAT journal by ID
- `DELETE /jurnal/kiat/{id}` - Delete KIAT journal by ID

#### TAJB Journal
- `GET /jurnal/tajb` - Get all TAJB journal data
- `POST /jurnal/tajb` - Create new TAJB journal
- `GET /jurnal/tajb/{id}` - Get specific TAJB journal by ID
- `PUT /jurnal/tajb/{id}` - Update TAJB journal by ID
- `DELETE /jurnal/tajb/{id}` - Delete TAJB journal by ID

#### TEKNOIS Journal
- `GET /jurnal/teknois` - Get all TEKNOIS journal data
- `POST /jurnal/teknois` - Create new TEKNOIS journal
- `GET /jurnal/teknois/{id}` - Get specific TEKNOIS journal by ID
- `PUT /jurnal/teknois/{id}` - Update TEKNOIS journal by ID
- `DELETE /jurnal/teknois/{id}` - Delete TEKNOIS journal by ID

#### TMJB Journal
- `GET /jurnal/tmjb` - Get all TMJB journal data
- `POST /jurnal/tmjb` - Create new TMJB journal
- `GET /jurnal/tmjb/{id}` - Get specific TMJB journal by ID
- `PUT /jurnal/tmjb/{id}` - Update TMJB journal by ID
- `DELETE /jurnal/tmjb/{id}` - Delete TMJB journal by ID

### Intellectual Property Endpoints (Full CRUD)

#### Lecturer HKI
- `GET /hki/dosen` - Get all lecturer HKI data
- `POST /hki/dosen` - Create new lecturer HKI
- `GET /hki/dosen/{id}` - Get specific lecturer HKI by ID
- `PUT /hki/dosen/{id}` - Update lecturer HKI by ID
- `DELETE /hki/dosen/{id}` - Delete lecturer HKI by ID

#### Student HKI
- `GET /hki/mhs` - Get all student HKI data
- `POST /hki/mhs` - Create new student HKI
- `GET /hki/mhs/{id}` - Get specific student HKI by ID
- `PUT /hki/mhs/{id}` - Update student HKI by ID
- `DELETE /hki/mhs/{id}` - Delete student HKI by ID

## Data Structure

All entities follow the same basic structure:

```json
{
  "id": 1,
  "judul": "Sample Title",
  "konten": "Sample content text"
}
```

## Installation & Setup

1. **Prerequisites**
   - Go 1.19 or higher
   - MySQL 5.7 or higher

2. **Database Setup**
   - Create a MySQL database named `lppm`
   - Import the provided SQL file: `lppm (4).sql`

3. **Environment Configuration**
   - Set the database connection string via environment variable:
     ```bash
     export DB_DSN="username:password@tcp(localhost:3306)/lppm?charset=utf8mb4&parseTime=True&loc=Local"
     ```
   - Or use the default: `root:@tcp(localhost:3306)/lppm?charset=utf8mb4&parseTime=True&loc=Local`

4. **Install Dependencies**
   ```bash
   go mod tidy
   ```

5. **Run the Application**
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8081`

## Architecture

The application follows Clean Architecture principles:

- **Entity Layer**: Data models and business entities
- **Repository Layer**: Data access and persistence
- **Use Case Layer**: Business logic and application rules
- **Interface Layer**: Controllers and routing
- **Infrastructure Layer**: Database configuration and external services

## Features

- **CORS Support**: Cross-origin requests are enabled
- **RESTful Design**: Follows REST principles
- **Error Handling**: Proper HTTP status codes and error messages
- **Database Integration**: Uses GORM for database operations
- **Modular Design**: Clean separation of concerns

## Usage Examples

### Get all vision and mission data
```bash
curl -X GET http://localhost:8081/visimisi
```

### Create new research entry
```bash
curl -X POST http://localhost:8081/penelitian/bame \
  -H "Content-Type: application/json" \
  -d '{"judul": "New Research", "konten": "Research content"}'
```

### Update existing entry
```bash
curl -X PUT http://localhost:8081/penelitian/bame/1 \
  -H "Content-Type: application/json" \
  -d '{"judul": "Updated Research", "konten": "Updated content"}'
```

### Delete entry
```bash
curl -X DELETE http://localhost:8081/penelitian/bame/1
```

## License

This project is developed for educational purposes. 