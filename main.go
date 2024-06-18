package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/Smartdevs17/go_tutor/sample"
)

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Welcome Gophers 👩‍💻")
// }

// func main() {
// 	// var x int
// 	// var y int

// 	// x = 2
// 	// y = 3

// 	// x := 4.0
// 	// y := 2.0

// 	x, y := 4.0, 2.0

// 	// var mean float64

// 	mean := x + y/2

// 	fmt.Println("The result is: ", mean)
// 	fmt.Printf("x=%v, type is %T\n", mean, mean)
// }

// func main() {
// 	x := 4

// 	if x == 4 {
// 		fmt.Println("x is 4")
// 	} else {
// 		fmt.Println("x is not 4")
// 	}

// 	n := 50
// 	switch n {
// 	case 1:
// 		fmt.Println("n is equal to 1")
// 	case 2:
// 		fmt.Println("n is equal to 2")
// 	case 3:
// 		fmt.Println("n is equal to 3")
// 	default:
// 		fmt.Println("n is not between 1 and 3")
// 	}
// }

// func main() {
// 	myNum := []int{}
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			myNum = append(myNum, i)
// 		}
// 	}
// 	fmt.Println("This is my Array: ", myNum)
// }

func main() {
	// sample.FizzBuzz()
	// sample.EvenEndedNumber()
	// fruits := []string{"Apple", "Orange", "Banana"}
	// fmt.Printf("Fruit= %v, Type is %T\n", fruits, fruits)
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i])
	// }
	// for i, name := range fruits {
	// 	fmt.Printf("%v %v\n", i, name)
	// }
	// for _, name := range fruits {
	// 	fmt.Println(name)
	// }
	// nums := []int{16, 8, 42, 4, 23, 15}
	// sample.GetMaxValue(nums)
	// myProfile := map[string]string{
	// 	"name":    "John Doe",
	// 	"age":     "30",
	// 	"address": "123 Main St, Anytown USA",
	// }
	// fmt.Println(len(myProfile))
	// fmt.Println(myProfile["name"])
	// fmt.Println(myProfile["age"])
	// fmt.Println(myProfile["address"])
	// value, ok := myProfile["date"]
	// // value, ok := myProfile["age"]
	// if ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Key not found")
	// }
	// myProfile["date"] = "2023-05-01"
	// delete(myProfile, "age")
	// fmt.Println(myProfile)
	wordCount := sample.WordCount("I want to go to the market by I Self To get icecream")
	fmt.Println(wordCount)

	//Go Pointers
	val := 10
	doublePtr(&val)
	fmt.Println(val)

	num, err := sqrt(25)
	// num, err := sqrt(-25);
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

	//Go Errors
	fmt.Println("Hello, World!")
	fmt.Println("Welcome Gophers ������")
}
func doublePtr(n *int) {
	*n *= 2
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number (%f)", n)
	}
	return math.Sqrt(n), nil
}

// write a function called contentType which will return the value of Content-Type header returned by making
// an HTTP GET request to url
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	return contentType, nil
}
