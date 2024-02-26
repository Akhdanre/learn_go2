package learngowebsite

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/if_else.gohtml"))
	t.ExecuteTemplate(writer, "if_else.gohtml", Page{
		Title: "Template Data action",
		Name: "akhdan",
	})
}

func TestTemplateDataAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataAction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
