package main

import "fmt"
// import "os"

type Point struct {
	x, y int
}

func main() {
	p := Point{3, 4}

	// 구조체의 값만 출력
	fmt.Printf("%v\n", p)
	// 구조체의 필드명까지 같이 출력
	fmt.Printf("%+v\n", p)
	// 해당 구조체 타입까지
	fmt.Printf("%#v\n", p)

	ptr := &p
	// 포인터값 출력
	fmt.Printf("포인터의 주소: %p\n", &ptr)
	fmt.Printf("포인터의 값: %p\n", ptr)

	// 타입 출력
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", ptr)
	
	fmt.Printf("bool: %T\n", true)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("bool: %t\n", false)

	// 10진수
	fmt.Printf("int: %d\n", 123)
	// 16진수
	fmt.Printf("int: %x\n", 123)
	// 8진수
	fmt.Printf("int: %o\n", 123)
	// 2진수
	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("float: %f\n", 78.9)
	fmt.Printf("float: %e\n", 123400000.0)
	fmt.Printf("float: %E\n", 123400000.0)

	// char 로 변환해서 출력
	fmt.Printf("char: %c\n", 97)
	fmt.Printf("char: %c\n", 'a')
	fmt.Printf("char: %c\n", "abcd")

	// 숫자, 문자열 너비 맞추기
	fmt.Printf("%6d,%6d\n", 12, 12345)
	fmt.Printf("%8.3f,%8.3f\n", 0.12, 123.45)
	fmt.Printf("%-8.3f,%8.3f\n", 0.12, 0.12)
	fmt.Printf("%6s,%6s\n", "hi", "bcon")
	fmt.Printf("%-6s,%-6s\n", "hi", "bcon")

	// 문자열 포매팅(f"")
	s := fmt.Sprintf("|%10s|\t|%10.5f|\t|%-10d|", "hi", 12.436, 2)
	fmt.Println(s)
}