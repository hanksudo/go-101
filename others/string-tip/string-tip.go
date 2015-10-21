package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	// concat
	a := "Hello"
	b := a + " World"
	fmt.Println(b)

	// multiline string
	multiStr := `
    I am Hank.

    Who are you?
    `
	fmt.Println(multiStr)

	// string -> int
	val1, _ := strconv.Atoi("100")
	fmt.Println(val1)

	// string -> float
	val2, _ := strconv.ParseFloat("1.5", 64)
	fmt.Println(val2)

	// repeat
	fmt.Println(strings.Repeat("co", 2) + "a")

	// trim
	fmt.Println(strings.TrimSpace(" I Love Golang."))
	fmt.Println(strings.TrimRight("Go Next Line\n", "\n"))

	// Upper / Lower case
	s := "I Love Golang."
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))

	// swap case
	fmt.Println(strings.Map(SwapCase, s))

	// command result
	out, _ := exec.Command("date").Output()
	fmt.Printf("%s", out)
	// fmt.Println(string(out))

	// base 8 and base 16
	fmt.Println(strconv.ParseInt("010", 8, 64))
	fmt.Println(strconv.ParseInt("ff", 16, 64))
	fmt.Println(strconv.ParseInt("0xff", 0, 64))

	// ASCII code
	asciiStr := "ABC"
	fmt.Println(asciiStr[0]) // 65
	fmt.Println(string(82))  // R
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	for i := 65; i <= 90; i++ {
		fmt.Print(string(i))
	}
	fmt.Println()
	for i := 97; i <= 122; i++ {
		fmt.Print(string(i))
	}
	fmt.Println()

	// replace
	longStr := "Apple Banana Orange Apple"
	replacedOne := strings.Replace(longStr, "Apple", "Banna", 1)
	fmt.Println(replacedOne)
	replacedAll := strings.Replace(longStr, "Apple", "Banna", -1)
	fmt.Println(replacedAll)

	// index
	fmt.Println(strings.Index(longStr, "Apple"))     // 0
	fmt.Println(strings.Index(longStr, "Banana"))    // 6
	fmt.Println(strings.LastIndex(longStr, "Apple")) // 20

	// split
	address := "001,New Taipei,Taiwan"
	for _, x := range strings.Split(address, ",") {
		fmt.Println(x)
	}
}

// rune is an alias for int32 and is equivalent to int32 in all ways.
func SwapCase(s rune) rune {
	switch {
	case 97 <= s && s <= 122:
		return s - 32
	case 65 <= s && s <= 90:
		return s + 32
	default:
		return s
	}
}
