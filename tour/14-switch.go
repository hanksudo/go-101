package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	x := 5
	switch x {
	case 1:
		fmt.Println("x is equal to 1")
	case 3, 4, 5:
		fmt.Println("x is equal to 3, 4 or 5")
	default:
		fmt.Println("no matched")
	}

	// fallthrough
	fmt.Println("")
	y := 6
	switch y {
	case 4:
		fmt.Println("y is equal to 4")
		fallthrough
	case 5:
		fmt.Println("y is equal to 5")
		fallthrough
	case 6:
		fmt.Println("y is equal to 6")
		fallthrough
	case 7:
		fmt.Println("y is equal to 7")
		fallthrough
	case 8:
		fmt.Println("y is equal to 8")
		fallthrough
	default:
		fmt.Println("no matched")
	}

	// detect weekday
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// swtich with no condition (this construct can be clean way to write long if-then-else chains.)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
