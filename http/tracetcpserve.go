package main

import (
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	//	"html/template"
	"io"
	//	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	//	"strings"
)

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
	http.ServeFile(w, r, "editcmd.html")
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	host := r.FormValue("host")
	port := r.FormValue("port")

	cmd := exec.Command("tcptraceroute")
	cmd.Stdout = &fw
	cmd.Stderr = &fw

	cmd.Args = append(cmd.Args, host)
	cmd.Args = append(cmd.Args, port)

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
}

func main() {
	http.HandleFunc("/editcmd/", editCommandHandler)
	http.HandleFunc("/exec/", execHandler)
	http.ListenAndServe(":8080", nil)
}
