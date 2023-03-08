1. Recursive adalah teknik pemrograman di mana sebuah fungsi memanggil dirinya sendiri secara berulang-ulang sampai kondisi berhenti tercapai. Contohnya : 
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // output: 120
}
Contoh di atas, fungsi factorial memanggil dirinya sendiri dengan parameter n-1 sampai mencapai kondisi n == 0 yang akan mengembalikan nilai 1 dan mengakhiri pemanggilan rekursif.
2. Number theory adalah cabang matematika yang mempelajari sifat-sifat bilangan bulat dan hubungan antara mereka. Contohnya menggunakan fungsi math.Sqrt untuk menghitung akar kuadrat dari sebuah bilangan bulat, math.Min untuk mencari nilai minimum antara dua bilangan bulat, dan math.Pow untuk menghitung bilangan pangkat. Selain itu, juga menggunakan loop for dan operator if untuk melakukan operasi dan komputasi pada bilangan bulat, seperti menentukan apakah suatu bilangan bulat adalah bilangan prima dan menghitung GCD dari dua bilangan bulat.
3. Searching dibagi menjadi 2 yaitu ada Linier Search dan Builtins Search. 
    - Linier Search adalah teknik pencarian sederhana yang mencari suatu elemen pada suatu koleksi data secara     berurutan, dari awal sampai akhir. Pada Golang, linear search dapat diterapkan pada slice. Contohnya 
    func linearSearch(arr []int, target int) bool {
    for i := 0; i < len(arr); i++ {
        if arr[i] == target {
            return true
        }
    }
    return false
    }
    - Builtins Search adalah pencarian data menggunakan built-in function atau fungsi bawaan yang sudah disediakan oleh Golang, seperti fungsi len() untuk menghitung panjang slice, array, map, atau string, atau fungsi cap() untuk menghitung kapasitas slice atau array.
4. Sorting atau pengurutan dalam bahasa pemrograman Golang adalah proses mengurutkan elemen-elemen pada suatu koleksi data seperti array, slice, atau map menjadi urutan tertentu berdasarkan kriteria tertentu, misalnya dari yang terkecil ke terbesar, dari yang terbesar ke terkecil, atau berdasarkan kriteria lain yang ditentukan. Sorting dibagi menjadi 4 yaitu Selection Sort, Counting Sort, Merge Sort, Builtins Sort in Java.
