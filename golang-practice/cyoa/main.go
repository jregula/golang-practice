package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io"
	"net/http"
	"log"
	"html/template"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
		<div>
			<br>
			{{ . }}
			<br/>
		</div>
		{{end}}
		<br/>
		<ul>
		{{range .Options}}
				<li><a href="http://localhost:8080/{{ .Arc }}">{{ .Text }}</li>
		{{end}}
		<ul/>
	</body>
</html>`

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

 
type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

func (h *Chapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	data := Chapter{
		Title: h.Title, 
		Story: h.Story,
		Options: h.Options,
	}

	err = t.Execute(w, data)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

}

func main() {

	jsonFile, err := os.Open("gopher.json")
	
	if err != nil {
		fmt.Println(err)
	}

	story, err := JsonStory(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	for key, value := range story {
		http.Handle("/" + key ,&Chapter{
			Title: value.Title,
			Story: value.Story,
            Options: value.Options,
		})

    }
	
	log.Fatal(http.ListenAndServe(":8080", nil))
	


}