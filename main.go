package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/markbates/pkger"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	const Port string = ":9001"
	const ServerURL string = "http://localhost" + Port

	// Static Site pkger
	dir := http.FileServer(pkger.Dir("/public"))
	http.Handle("/", dir)

	// Static Site std
	// fs := http.FileServer(http.Dir("./public"))
	// http.Handle("/", fs)

	// API
	http.HandleFunc("/api", index)
	http.HandleFunc("/api/about", about)

	// openbrowser(ServerURL)

	fmt.Println("Servidor iniciado em: ", ServerURL)
	return http.ListenAndServe(Port, nil)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
