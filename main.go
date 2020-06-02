package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//-------------------------------------Soal 0--------------------------------------

//function to check string is palindrome
func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

//function to check string is number
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func polindromes(w http.ResponseWriter, r *http.Request) {

	//make validate just method POST to send data input
	message := ""
	if r.Method == "GET" {
		message = "Gunakan method post untuk mengakses soal 1"
	}
	//method POST
	if r.Method == "POST" {
		var data = r.FormValue("input")
		ranges := strings.Fields(data)

		//validation input just number
		if isNumeric(ranges[0]) && isNumeric(ranges[1]) {
			from, _ := strconv.Atoi(ranges[0])
			to, _ := strconv.Atoi(ranges[1])

			//check is number one must smaller from number two
			if from > to {
				message = "Inputan angka pertama harus lebih kecil, Contoh input 100 150"
			} else {
				res := 0
				for x := from; x < to; x++ {
					var word string = strconv.Itoa(x)
					if isPalindrome(word) {
						res++
					}
				}
				message = strconv.Itoa(res)
			}
		} else {
			message = "Data input bukan number, Contoh input 1 10"
		}
	}

	fmt.Println(message)
	fmt.Fprintf(w, message)
}

//-------------------------------------Soal 1--------------------------------------
func pustakawan(w http.ResponseWriter, r *http.Request) {

	//make validate just method POST to send data input
	if r.Method == "GET" {
		fmt.Fprintf(w, "Gunakan method post untuk mengakses soal 1")
	}
	//method POST
	if r.Method == "POST" {
		var data = r.FormValue("input")
		// data := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
		listBook := strings.Fields(data)

		category := [10]string{"6", "7", "0", "9", "4", "8", "1", "2", "5", "3"}

		var resBook []string
		for i := 0; i < len(category); i++ {

			//sorting category books
			var bookCategory []string
			for l := 0; l < len(listBook); l++ {
				s := string([]byte{listBook[l][0]})
				if s == category[i] {
					bookCategory = append(bookCategory, listBook[l])
				}
			}

			//sorting height books
			for x := len(bookCategory); x > 0; x-- {
				for j := 1; j < x; j++ {
					left, _ := strconv.Atoi(bookCategory[j-1][2:])
					right, _ := strconv.Atoi(bookCategory[j][2:])
					if left < right {
						intermediate := bookCategory[j]
						bookCategory[j] = bookCategory[j-1]
						bookCategory[j-1] = intermediate
					}
				}
			}

			//delete more than 2 books
			tmp := bookCategory[:0]
			dict := make(map[string]int)
			for _, v := range bookCategory {
				dict[v[0:2]] = dict[v[0:2]] + 1
				if dict[v[0:2]] <= 2 {
					tmp = append(tmp, v)
				}
			}

			//append to respon book
			for l := 0; l < len(tmp); l++ {
				resBook = append(resBook, tmp[l])
			}
		}
		justString := strings.Join(resBook, " ")
		fmt.Fprintf(w, justString)
	}
}

//-------------------------------------Soal 2--------------------------------------
func missingNumbers(w http.ResponseWriter, r *http.Request) {
	//make validate just method POST to send data input
	if r.Method == "GET" {
		fmt.Fprintf(w, "Gunakan method post untuk mengakses soal 2")
	}
	//method POST
	if r.Method == "POST" {
		// s := "12346789"
		// s := "23242526272830"
		// s := "101102103104105106107108109111112113"
		// s := "100001100002100003100004100006"
		var s = r.FormValue("input")
		var arr []int
		res := ""
		digit := getDigit(s)
		for i, r := range s {
			res = res + string(r)

			if digit == 1 {
				tmp, _ := strconv.Atoi(res)
				arr = append(arr, tmp)
				res = ""
			} else {
				if i > 0 && (i+1)%digit == 0 {
					tmp, _ := strconv.Atoi(res)
					arr = append(arr, tmp)
					res = ""
				}
			}
		}
		// fmt.Println(getMissingNo(arr, len(arr)))
		fmt.Fprintf(w, "%d", getMissingNo(arr, len(arr)))
	}
}

//get digits from string
func getDigit(s string) int {

	n := len(s)
	for i := 1; i < 1000001; i++ {
		from, _ := strconv.Atoi(s[0:i])
		to, _ := strconv.Atoi(s[n-i:])
		if (from + (n / i)) == to {
			return i
		}
	}

	return n
}

//function to find missing number on array
func getMissingNo(a []int, n int) int {
	ideal := 0
	for i := a[0]; i <= a[n-1]; i++ {
		ideal += i
	}

	sum := 0
	for _, v := range a {
		sum += v
	}

	return ideal - sum
}

// -------------------------------------main method------------------------------------

func main() {

	http.HandleFunc("/soal-0", polindromes)
	http.HandleFunc("/soal-1", pustakawan)
	http.HandleFunc("/soal-2", missingNumbers)

	fmt.Println("Listen :8080")
	http.ListenAndServe(":8080", nil)

}
