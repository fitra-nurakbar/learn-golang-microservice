# Menggunakan gambar Linux Alpine sebagai dasar
FROM golang:1.21.0-alpine AS production

# Menambahkan metadata tentang gambar (opsional)
LABEL maintainer="Fitra Nurakbar <fitranurakbar378@gmail.com>"

# Menyusun direktori kerja di dalam container
WORKDIR /app

# Menyalin file go.mod dan go.sum untuk mengelola dependensi
COPY go.mod go.sum ./

# Mengambil dependensi yang diperlukan
RUN go mod download

# Menyalin seluruh kode sumber proyek ke dalam container
COPY . .

# Membangun aplikasi Go
RUN go build -o main

# Perintah yang akan dijalankan saat kontainer dimulai
CMD ["./main"]
