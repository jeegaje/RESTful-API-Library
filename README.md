# Simple RESTful API Library

Ini adalah proyek sederhana yang mengimplementasikan RESTful API untuk sebuah perpustakaan menggunakan GORM, Gin, dan MySQL. Proyek ini bertujuan untuk menunjukkan bagaimana Anda dapat membuat API sederhana dengan menggunakan teknologi-teknologi tersebut.

## Persyaratan

Sebelum Anda dapat menjalankan proyek ini, pastikan Anda telah menginstal:

- Go
- GORM: `go get -u gorm.io/gorm`
- Gin: `go get -u github.com/gin-gonic/gin`
- MySQL: [https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/)

## API Endpoints

Berikut adalah beberapa endpoint API yang tersedia:

### Health Check
- `GET /health`: Memeriksa status kesehatan server.

### Buku (Book)
- `GET /books`: Mendapatkan daftar semua buku.
- `GET /books/{id}`: Mendapatkan detail buku berdasarkan ID.
- `POST /book`: Menambahkan buku baru.
- `DELETE /book/{id}`: Menghapus buku berdasarkan ID.

### Penulis (Author)
- `GET /authors`: Mendapatkan daftar semua penulis.
- `GET /authors/{id}`: Mendapatkan detail penulis berdasarkan ID.
- `POST /author`: Menambahkan penulis baru.
- `DELETE /author/{id}`: Menghapus penulis berdasarkan ID.

### Genre
- `GET /genres`: Mendapatkan daftar semua genre.
- `GET /genres/{id}`: Mendapatkan detail genre berdasarkan ID.
- `POST /genre`: Menambahkan genre baru.
- `DELETE /genre/{id}`: Menghapus genre berdasarkan ID.

Pastikan untuk merujuk ke implementasi sesungguhnya di kode Anda dan menyesuaikan dengan deskripsi masing-masing endpoint.
