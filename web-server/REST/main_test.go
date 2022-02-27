package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestJsonHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	ast.Nil(err)
	ast.Equal(3, len(list))
	ast.Equal(10, list[1].Age)
	ast.Equal("test", list[2].Name)
}

func TestStudentJsonHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)

	var student Student
	err := json.NewDecoder(res.Body).Decode(&student)
	ast.Nil(err)
	ast.Equal(1, student.Id)
	ast.Equal("bconfiden2", student.Name)
	ast.Equal(27, student.Age)
	ast.Equal(100, student.Score)
}

func TestPostHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
			strings.NewReader(`{"Id":1234,"Name":"post","Age":123,"Score":0}`))

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusCreated, res.Code)
	

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/4", nil)

	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)

	var student Student
	err := json.NewDecoder(res.Body).Decode(&student)
	ast.Nil(err)
	ast.Equal(4, student.Id)
	ast.Equal("post", student.Name)
	ast.Equal(123, student.Age)
	ast.Equal(0, student.Score)
}