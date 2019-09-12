package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/mattetti/goHtmlTemplateExample/pb/csv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
	}

	defer func() {
		if e := recover(); e != nil {
			byts, _ := GetBytes(e)
			w.Write(byts)
		}
	}()

	values := r.PostForm
	parseParams(values)
	w.Write([]byte("Success"))
}

func parseParams(values url.Values) {
	for k, _ := range values {
		split := strings.Split(k, "\r\n")
		for _, item := range split {
			if strings.Trim(item, " ") == "" {
				continue
			}
			log.Println(item)
		}
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	dirList, e := ioutil.ReadDir(csv.FileBase)
	if e != nil {
		log.Println("read dir error")
		return
	}
	list := []string{}
	for _, v := range dirList {
		name := v.Name()
		ok := strings.HasSuffix(name, "csv")
		if ok {
			list = append(list, name)
		}
	}
	joined := strings.Join(list, "\r\n")
	w.Write([]byte(joined))
}
