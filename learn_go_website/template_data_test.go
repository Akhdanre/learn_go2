package learngowebsite

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Akhdan",
	})
}

type Address struct {
	Jalan  string 
}

type Page struct {
	Title string
	Name  string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Akhdan Robbani",
		Address: Address{
			Jalan: "belum ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
