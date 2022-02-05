package main

import (
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Join("dir", "dir", "file"))
	fmt.Println(filepath.Join("dir//", "file"))
	fmt.Println(filepath.Join("../dir/..", "file"))
	
	p := filepath.Join("dir", "../../", "dir2", "dir3", "file.txt")
	fmt.Println("전체경로:", p)
	fmt.Println("디렉토리:", filepath.Dir(p))
	fmt.Println("파일명:", filepath.Base(p))
	fmt.Println("절대경로?", filepath.IsAbs(p))
	fmt.Println("확장자:", filepath.Ext(p))
	fmt.Println("확장자 제거:", strings.TrimSuffix(p, filepath.Ext(p)))

	rel, err := filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}