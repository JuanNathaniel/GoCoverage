package main

import "fmt"

func main() {
    // Mendefinisikan variabel untuk angka yang akan dioperasikan
    var num1, num2 float64

    // Meminta pengguna untuk memasukkan dua angka
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)

    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    // Melakukan operasi aritmatika sederhana
    fmt.Printf("\nHasil Penambahan: %.2f\n", num1 + num2)
    fmt.Printf("Hasil Pengurangan: %.2f\n", num1 - num2)
    fmt.Printf("Hasil Perkalian: %.2f\n", num1 * num2)

    // Menangani pembagian dengan cek pembagian oleh nol
    if num2 != 0 {
        fmt.Printf("Hasil Pembagian: %.2f\n", num1 / num2)
    } else {
        fmt.Println("Tidak bisa melakukan pembagian karena angka kedua adalah nol.")
    }
}
