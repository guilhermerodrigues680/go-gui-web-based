package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guilhermerodrigues680/go-gui-web-based/apis/apiv1"
	"github.com/guilhermerodrigues680/go-gui-web-based/browserutil"
	"github.com/markbates/pkger"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	const RunInContainer bool = false
	const Port string = ":9001"
	const ServerURL string = "http://localhost" + Port

	// Static Site pkger
	dir := http.FileServer(pkger.Dir("/public"))
	http.Handle("/", dir)

	// Static Site std
	// fs := http.FileServer(http.Dir("./public"))
	// http.Handle("/", fs)

	// API
	http.HandleFunc("/api", apiv1.Index)
	http.HandleFunc("/api/about", apiv1.About)

	if RunInContainer == false {
		browserutil.Openbrowser(ServerURL)
	}

	fmt.Println("Servidor iniciado em: ", ServerURL)
	return http.ListenAndServe(Port, nil)
}
