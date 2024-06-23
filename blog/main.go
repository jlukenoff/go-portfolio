package blog

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/russross/blackfriday"
)

func MarkdownToHtml(md string) string {
	return string(blackfriday.MarkdownCommon([]byte(md)))
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	postName := path.Base(r.URL.Path)
	pathToPostMarkdownFile := path.Join("static/posts", postName+".md")

	postMarkdownString, err := os.ReadFile(pathToPostMarkdownFile)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	postHtml := MarkdownToHtml(string(postMarkdownString))

	tpl, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	templateData := struct {
		PostHtml template.HTML
	}{
		PostHtml: template.HTML(postHtml),
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, templateData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the captured output to the response writer
	w.Write(buf.Bytes())
}
