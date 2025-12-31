# Aplikasi Donasi

Layanan backend untuk aplikasi donasi, dibangun menggunakan Go. Proyek ini menyediakan API untuk mengelola donasi, pengguna, dan kampanye.

## Daftar Isi

- [Fitur](#fitur)
- [Memulai](#memulai)
  - [Prasyarat](#prasyarat)
  - [Instalasi](#instalasi)
- [Pengaturan Database](#pengaturan-database)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Endpoint API](#endpoint-api)
- [Berkontribusi](#berkontribusi)
- [Lisensi](#lisensi)

## Fitur

- Pemrosesan donasi yang aman.
- Manajemen akun pengguna.
- Pembuatan dan pengelolaan kampanye donasi.
- Pelacakan riwayat donasi.

## Memulai

Untuk menjalankan proyek ini di lingkungan lokal Anda, ikuti langkah-langkah di bawah ini.

### Prasyarat

Pastikan perangkat Anda telah terpasang perangkat lunak berikut:

- [Go](https://golang.org/dl/) (versi 1.18 atau lebih baru)
- [MySQL](https://dev.mysql.com/downloads/installer/) (versi 8.0 atau lebih baru)
- [Git](https://git-scm.com/downloads/)

### Instalasi

1.  Salin repositori ini ke mesin lokal Anda.
    ```sh
    git clone https://github.com/nama-pengguna-anda/nama-repositori-anda.git
    ```

2.  Masuk ke direktori proyek.
    ```sh
    cd golang
    ```

3.  Unduh dependensi yang diperlukan.
    ```sh
    go mod tidy
    ```

## Pengaturan Database

Proyek ini menggunakan MySQL sebagai databasenya.

1.  Masuk ke shell MySQL Anda.
    ```sh
    mysql -u root -p
    ```

2.  Buat database baru untuk aplikasi.
    ```sql
    CREATE DATABASE donation_app;
    ```

3.  Buat pengguna baru dan berikan hak akses ke database tersebut. Ganti `password` dengan kata sandi yang kuat.
    ```sql
    CREATE USER 'donation_user'@'localhost' IDENTIFIED BY 'password';
    GRANT ALL PRIVILEGES ON donation_app.* TO 'donation_user'@'localhost';
    FLUSH PRIVILEGES;
    ```

4.  Keluar dari shell MySQL.
    ```sql
    EXIT;
    ```

5.  Jalankan skrip SQL berikut untuk membuat tabel yang diperlukan. Anda dapat menjalankan ini dengan mengimpor file `.sql` atau menyalin isinya ke klien MySQL Anda.

    Buat tabel `users`:
    ```sql
    USE donation_app;

    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    ```

    Buat tabel `campaigns`:
    ```sql
    CREATE TABLE campaigns (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        goal_amount DECIMAL(15, 2) NOT NULL,
        current_amount DECIMAL(15, 2) DEFAULT 0.00,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INT,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );
    ```

    Buat tabel `donations`:
    ```sql
    CREATE TABLE donations (
        id INT AUTO_INCREMENT PRIMARY KEY,
        amount DECIMAL(15, 2) NOT NULL,
        status VARCHAR(50) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INT,
        campaign_id INT,
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
    );
    ```

## Menjalankan Aplikasi

Setelah konfigurasi selesai, Anda dapat menjalankan server aplikasi.

```sh
go run main.go
```

Server akan berjalan di `http://localhost:8080` secara default.

## Endpoint API

Berikut adalah beberapa contoh endpoint API yang tersedia.

| Metode | Endpoint                | Deskripsi                               |
|--------|---------------------------|-------------------------------------------|
| `POST` | `/api/users/register`     | Mendaftarkan pengguna baru.               |
| `POST` | `/api/users/login`        | Masuk ke akun pengguna.                   |
| `POST` | `/api/donations`          | Membuat donasi baru.                      |
| `GET`  | `/api/donations/{id}`     | Mendapatkan detail donasi berdasarkan ID. |
| `GET`  | `/api/campaigns`          | Mendapatkan daftar semua kampanye.        |
| `GET`  | `/api/campaigns/{id}`     | Mendapatkan detail kampanye berdasarkan ID.|

## Berkontribusi

Kontribusi dari Anda sangat kami harapkan. Jika Anda ingin berkontribusi, silakan buat *fork* dari repositori ini, buat cabang fitur baru, dan kirimkan *pull request*.

1.  Buat *fork* dari proyek.
2.  Buat cabang fitur baru (`git checkout -b fitur/FiturBaru`).
3.  Lakukan perubahan Anda (`git commit -m 'Menambahkan FiturBaru'`).
4.  Unggah ke cabang (`git push origin fitur/FiturBaru`).
5.  Buka *Pull Request*.

## Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT. Lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.
