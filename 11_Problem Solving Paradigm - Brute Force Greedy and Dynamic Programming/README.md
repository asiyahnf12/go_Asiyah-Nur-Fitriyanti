1. Problem Solving Paradigma dalam Golang merujuk pada cara-cara untuk memecahkan masalah dengan menggunakan bahasa pemrograman Golang.
2. Complete search dalam Golang merujuk pada teknik pencarian yang mencoba semua kemungkinan solusi secara sistematis. Teknik ini berguna untuk menyelesaikan masalah yang ukurannya kecil hingga sedang dengan jumlah kemungkinan solusi yang terbatas.
3. Divide and Conquer dalam Golang adalah sebuah paradigma pemrograman yang memecahkan masalah menjadi beberapa submasalah yang lebih kecil, menyelesaikan submasalah tersebut secara rekursif, dan kemudian menggabungkan solusi untuk mendapatkan solusi untuk masalah asli.
4. Greedy pada dasarnya adalah suatu strategi algoritma dimana pemilihan langkah dilakukan dengan memilih opsi terbaik pada setiap langkah, dengan harapan bahwa keputusan yang diambil pada setiap langkah akan menghasilkan solusi terbaik secara keseluruhan. Dalam bahasa pemrograman Golang, greedy biasanya digunakan dalam implementasi algoritma seperti algoritma Greedy Best-First Search (GBFS) pada pencarian jalur terpendek dalam grafik, algoritma Greedy Algorithm pada masalah optimasi kombinatorial, atau algoritma Greedy Strategy pada masalah pengalokasian sumber daya. Contohnya 
func GreedyAlgorithm(weights []int, values []int, capacity int) int {
    n := len(weights)
    var maxVal int
    for i := 0; i < n; i++ {
        if weights[i] <= capacity {
            capacity -= weights[i]
            maxVal += values[i]
        }
    }
    return maxVal
}
