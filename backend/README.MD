## Cara Menjalakan
Masuk Ke Folder backend lalu ketik go run main.go
Atau Bisa  Menggunakan URL ini : https://hicoder.herokuapp.com/

Konfigurasi server terdapat pada file .env 

## LIST API

- `POST`: `/api/auth/login`
- `POST`: `/api/auth/logout`
- `POST`: `/api/auth/register`
- `GET` : `/api/user/profile`
- `PUT` : `/api/user/update/:id`
- `GET` : `/api/user/mentor/mentorlist`
- `GET` : `/api/user/mentor/mentorlist?skill`
- `GET` : `/api/user/mentor/detail/:id`
- `GET` : `/api/user/booking/mentor/:id(mentor)`
- `GET` : `/api/user/booking/status`
- `GET` : `/api/user/booking/mentor/kontak/:bookid`
- `GET` : `/api/user/artikel`
- `GET` : `/api/user/artikel/:id`
- `GET` : `/api/mentor/acc/:bookid`
- `POST`: `/api/admin/registermentor`
- `POST`: `/api/admin/addartikel`




## REQUEST BODY
- `LOGIN : {username ,password}`
- `REGISTER: {username,name,password,role}`
- `UPDATE: {username,name,password,Address,Phone,Email}`
- `ADDARTIKEL: {judul,content}`
- `REGISTERMENTOR: {Name,Address,Phone,Email,Skill,Bio}`

## RESPONSE BODY
- `LOGIN : {status,data}`
- `REGISTER: {status,"Berhasil Register"}`
- `Logout : {status,"Berhasil Logout"}`
- `UPDATE: {status,"Berhasil Update"}`
- `ADDARTIKEL: {status,"Berhasil Tambah Artikel"}`
- `REGISTERMENTOR: {status,"Berhasil Register Mentor"}`
- `MENTORLIST: {status,data}`
- `MENTORDETAIL: {status,data}`
- `BOOKING: {status,"Berhasil Mengirim Request"}`
- `BOOKINGSTATUS: {status,data}`
- `BOOKINGKONTAK: {status,data}`
- `ARTIKEL: {status,data}`
- `ARTIKELDETAIL: {status,data}`