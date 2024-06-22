package blog

import (
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/russross/blackfriday"
)

func MarkdownToHtml(md string) string {
	return string(blackfriday.MarkdownCommon([]byte(md)))
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	postName := path.Base(r.URL.Path)
	pathToPostMarkdownFile := path.Join("posts", postName+".md")

	postMarkdownString, err := os.ReadFile(pathToPostMarkdownFile)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	postHtml := MarkdownToHtml(string(postMarkdownString))

	tpl, err := template.ParseFiles("client/blog.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, postHtml)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
