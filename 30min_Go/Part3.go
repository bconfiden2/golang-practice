package main

import "fmt"

func switchtest(value int) (retu string) {
	switch {
	case value >= 90:
		retu = "over 90"
	case value >= 30:
		retu = "over 30"
		fmt.Print("over 30 but Fall Through! ")
		fallthrough
	case value >= 20:
		retu = "over 20"
	default:
		fmt.Print("Default!")
		retu = "zzz"
	}
	return retu
}

func main() {
	sum := 0
	for i := 0 ; i < 10 ; i++ {
		sum += i
		fmt.Println("Index:", i, "Sum:", sum)
	}

	fmt.Println()

	i := 1
	sum = 0
	for sum + i < 100 {
		i++
		sum += i
	}
	if v := sum*2 ; v < 200 {
		fmt.Println(sum)
	}
	if i < 10 {
		fmt.Println("under 10")
	} else if i == 10 {
		fmt.Println("same 10")
	} else {
		fmt.Println("over 10")
	}

	fmt.Println()
	for sum > 10 {
		fmt.Println(sum, switchtest(sum))
		sum /= 2
	}
}