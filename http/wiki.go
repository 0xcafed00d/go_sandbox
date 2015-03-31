package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %v!", r.URL.Path[1:])
	fmt.Fprint(w, spew.Sdump(r))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(spew.Sdump(r))
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(spew.Sdump(r))
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(spew.Sdump(r))
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	log.Printf("%s", p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func editCommandHandler(w http.ResponseWriter, r *http.Request) {
	//log.Print(spew.Sdump(r))
	renderTemplate(w, "editcmd", nil)
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	//log.Print(spew.Sdump(r))
	//title := r.URL.Path[len("/exec/"):]
	body := r.FormValue("body")
	argv := strings.Split(body, " ")

	cmd := exec.Command(argv[0])

	cmd.Stdout = &fw
	cmd.Stderr = &fw

	if len(argv) > 1 {
		cmd.Args = append(cmd.Args, argv[1:]...)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/editcmd/", editCommandHandler)
	http.HandleFunc("/exec/", execHandler)
	http.ListenAndServe(":8080", nil)
}
