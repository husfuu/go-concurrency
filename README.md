#### Paralel Calculation
- Menghitung total penjumlahan angka 1-100
- Perhitungan dibagi 4 bagian, 1-25, 26-50, 51-75 dan 76-100. Artinya ada 4 goroutine yang digunakan secara bersamaan.
- function `calculate` menerima channel yang nantinya akan dimasukan hasil perhitungan dari function tersebut.
- karena 4 perhitungan tersebut saling dependant (hasil-hasilnya akan digabungkan lagi), maka kita menggunakan channel `resultChan` untuk mengumpulkan hasil-hasil perhitungan setiap goroutine.
- untuk block code ini, di code ini tidak ada jaminan eksplisit kalo kode ini dijalankan sebelum `resultChannel` diisi oleh semua goroutine. Namun, cara ini tetap berfungsi karena pengambilan nilai dari channel (<-resultChan) akan memblokir (menunggu) hingga ada nilai yang dikirimkan ke channel oleh goroutine yang sedang berjalan.
    ```go
	totalSum := 0
	for i := 0; i < 4; i++ {
		totalSum += <-resultChannel
	}
    ```


<!-- - **ChatGpt Chat**: https://chatgpt.com/share/c0eb8dba-db9a-4224-8e68-73ac197de2f6 -->