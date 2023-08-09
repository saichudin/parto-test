# parto-test
Parto Jambe test golang

build with GO 1.18

ubah setting database sesuai kebutuhan difile 
```
.env
```

untuk menjalankan app :
go run main.go

dan otomatis database akan di migrate

Login :
Hit ke endpoint POST {base_url}/login
dengan json :
```
{
  "username" : "admin",
  "password" : "password
}
```
kemudian copy token yang didapat dari login dan tambahkan di header "Token" : {value} untuk setiap request ke endpoint contact

API Documentation
```
https://documenter.getpostman.com/view/12200000/2s9Xy2NX1e
```
