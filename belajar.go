package main

import "fmt"

func main() {
	tipe_data()
	konversi_type_data()
	operasi_matematika()
	array()
	slice()
}

func tipe_data() {

	fmt.Println("Print-------------------")
	// perbedaan Printf dan Println
	// Printf = mencetak dengan format
	// Println = mencetak dengan baris baru
	umur := 20
	fmt.Printf("Umur saya %d tahun\n", umur)
	fmt.Println("Umur saya", umur, "tahun")

	fmt.Println("\nString-------------------")
	// Tipe data string
	var nama string = "Rizki Faris"
	fmt.Println(nama)

	// function untuk string
	fmt.Println(len(nama)) // menghitung panjang string
	fmt.Println(nama[0])   // mengambil karakter pertama (output berupa byte)
	fmt.Println(nama[0:5]) // mengambil karakter dari index 0 s/d 4
	fmt.Println(nama[:5])  // mengambil karakter dari awal s/d index 4
	fmt.Println(nama[6:])  // mengambil karakter dari index 6 s/d akhir

	fmt.Println("\nINT-------------------")

	// Tipe data integer (data harus di sesuaikan dengan kebutuhan untuk mencegah pemborosan memory)
	var int8 int8 = 127          // -128 s/d 127
	int16 := 32767               // -32768 s/d 32767
	int32 := 2147483647          // -2147483648 s/d 2147483647
	int64 := 9223372036854775807 // -9223372036854775808 s/d 9223372036854775807
	fmt.Println(int8, int16, int32, int64)

	// Tipe data unsigned integer (hanya bilangan positif)
	var uint8 uint8 = 255                    // 0 s/d 255
	uint16 := 65535                          // 0 s/d 65535
	uint32 := 4294967295                     // 0 s/d 4294967295
	var uint64 uint64 = 18446744073709551615 // 0 s/d 18446744073709551615
	fmt.Println(uint8, uint16, uint32, uint64)

	fmt.Println("\nFloat-------------------")

	// Tipe data float (untuk bilangan desimal)
	var float32 float32 = 3.14   // kurang akurat
	float64 := 3.141592653589793 // lebih akurat
	fmt.Println(float32, float64)

	// alias tipe data
	// byte = uint8
	// rune = int32
	// int = int32 atau int64 (tergantung sistem operasi)
	// uint = uint32 atau uint64 (tergantung sistem operasi)

	fmt.Println("\nBoolean-------------------")

	// Tipe data boolean
	var benar bool = true
	salah := false
	fmt.Println(benar, salah)

	fmt.Println("\nVariable-------------------")
	// variable
	// macam-macam cara deklarasi variable
	var nama1 string = "Rizki" // deklarasi variable dengan tipe data
	var nama2 = "Faris"        // deklarasi variable tanpa tipe data
	nama3 := "Rizki Faris"     // deklarasi variable tanpa tipe data dan var
	var (                      // deklarasi dengan banyak variable
		fristName string = "Rizki"
		lastName         = "Faris"
	)
	fmt.Println(nama1, nama2, nama3, fristName, lastName)

	// const
	const phi = 3.14 // konstanta adalah variabel yanag tidak bisa diubah
	fmt.Println(phi)
}

func konversi_type_data() {
	fmt.Println("\nkonversi Tipe data int-------------------")
	var nilai32 int32 = 1000000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // hasilnya -32 karena melebihi batas nilai int8 (disebut overflow)

	fmt.Println("\nkonversi Tipe data String-------------------")
	var name string = "Rizki"
	var byteName = name[0]    // mengambil byte dari string
	var bukanbyte = name[0:5] // mengambil byte dari string tidak perlu di konversi kembali
	var stringName = string(byteName)
	fmt.Println(byteName)
	fmt.Println(stringName)
	fmt.Println(bukanbyte)

	fmt.Println("\ntype declaration-------------------")
	type NoKTP string
	var nofaris NoKTP = "1234567890"
	var contohktp = NoKTP("2123")

	fmt.Println(nofaris)
	fmt.Println(contohktp)
}

func operasi_matematika() {
	fmt.Println("\noperasi matematika-------------------")
	// operasi matematika
	var a = 10
	var b = 5
	var tambah = a + b
	var kurang = a - b
	var kali = a * b
	var bagi = a / b
	var modulus = a % b
	fmt.Println(tambah, kurang, kali, bagi, modulus)

	// operasi unary
	var c = 10
	c++ // c = c + 1
	fmt.Println(c)
	c-- // c = c - 1
	fmt.Println(c)
	c += 5 // c = c + 5
	fmt.Println(c)
	c -= 5 // c = c - 5
	fmt.Println(c)
	c *= 5 // c = c * 5
	fmt.Println(c)
	c /= 5 // c = c / 5
	fmt.Println(c)
	c %= 5 // c = c % 5
	fmt.Println(c)

	fmt.Println("\noperasi perbandingan-------------------")
	// operasi perbandingan
	var d = 10
	var e = 5
	fmt.Println(d == e) // sama dengan
	fmt.Println(d != e) // tidak sama dengan
	fmt.Println(d < e)  // kurang dari
	fmt.Println(d > e)  // lebih dari
	fmt.Println(d <= e) // kurang dari sama dengan
	fmt.Println(d >= e) // lebih dari sama dengan

	fmt.Println("\noperasi Boolean-------------------")
	// operasi boolean
	var t = true
	var f = false
	fmt.Println(t && f) // dan
	fmt.Println(t || f) // atau
	fmt.Println(!t)     // kebalikan

}

func array() {
	fmt.Println("\nTipe Array-------------------")
	// array
	var buah [3]string // array dengan panjang 3
	buah[0] = "Apel"
	buah[1] = "Jeruk"
	buah[2] = "Mangga"
	fmt.Println(buah)
	fmt.Println(buah[0])
	fmt.Println(buah[1])
	fmt.Println(buah[2])

	var angka = [3]int{1, 2, 3} // array dengan panjang 3
	fmt.Println(angka)
	var keranjang = [...]string{"naga", "nangka", "pir"} // array dengan panjang dinamis
	fmt.Println(keranjang)

	fmt.Println("\nFunction Array-------------------")
	// function array
	var buah1 = [3]string{"Apel", "Jeruk", "Mangga"}
	fmt.Println(buah1)
	fmt.Println(len(buah1))
	fmt.Println(buah1[0]) // mengambil nilai array
	fmt.Println(buah1[1] == "Jeruk")
	buah1[0] = "Pisang" // mengubah nilai array
	fmt.Println(buah1)
}

func slice() {
	fmt.Println("\nTipe Slice-------------------")
	// slice
	var kelompok = [...]string{"Rizki", "Faris", "Nur", "Fadilah"}
	kelompok1 := kelompok[0:3] // mengambil data dari index 0 s/d 2
	fmt.Println(kelompok1)
	kelompok2 := kelompok[1] // mengambil data dari index 1 s/d akhir
	fmt.Println(kelompok2)
}
