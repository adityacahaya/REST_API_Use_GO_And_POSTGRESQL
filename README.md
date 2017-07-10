# REST_API_Use_GO_And_POSTGRESQL

project ini berisi 2 buah file program dengan bahasa Go yang saya gunakan untuk melakukan tes serta belajar bagaimana :

#File TestJsonDbRest.go
1. mengambil data dari database Postgresql
2. menyimpan data ke dalam variable Struct yang sudah di definisikan
3. membentuk data tadi kedalam JSON dengan Marshal
4. membentuk REST API dengan menggunakan HTTP dan memanggil fungsi yang digunakan untuk menampilkan JSON kedalam web localhost dengan port    yang sudah di tentukan

#File TestUnmarshal.go
1. mengakses data REST API dari web dengan HTTP Get dari url website yang sudah di definisikan (dalam program ini adalah localhost dengan port yang    sudah di definisikan di file sebelumnya)
2. mengubah data menjadi bentuk JSON menggunakan Unmarshal
3. menampilkan data yang di dapat dari API ke dalam program Go
