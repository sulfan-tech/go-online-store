# Gunakan image resmi Golang sebagai parent image
FROM golang:1.20-alpine AS builder

# Setel direktori kerja saat ini di dalam container
WORKDIR /app

# Salin file yang diperlukan termasuk .env
COPY go.mod go.sum .env ./

# Tampilkan daftar file di direktori kerja
RUN ls -la

# Unduh dependensi
RUN go mod download

# Salin sisa kode sumber ke dalam container
COPY . .

# Bangun aplikasi Go dengan GO111MODULE=on
RUN go build -o myapp ./server/cmd

# Stage 2: Setup the final image
FROM alpine:latest
WORKDIR /root/
# Pastikan file .env disalin ke lokasi yang benar
COPY --from=builder /app/.env /root/.env
COPY --from=builder /app/myapp .
CMD ["./myapp"]