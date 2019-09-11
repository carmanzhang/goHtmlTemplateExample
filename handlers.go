package main

import (
	"fmt"
	"github.com/mattetti/goHtmlTemplateExample/pb/csv"
	"net/http"
	"strings"
)

type Msg struct {
	FileName string
	Content  [][]string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	s := r.RequestURI
	split := strings.Split(s, "=")
	file := csv.FileBase + split[1]
	//file := "C:/Users/Zhang_Li/idea-workspace/near_synonym_extraction/src/test/0_62_word_tag.csv"
	content := csv.ReadCsv(file, 1)

	passedObj := Msg{
		FileName: file,
		Content:  content,
	}

	templates.ExecuteTemplate(w, "homePage", passedObj)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
	}

	fmt.Println(r.PostForm)
}
