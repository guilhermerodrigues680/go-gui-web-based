package apiv1

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func Dev(w http.ResponseWriter, r *http.Request) {

	if _, err := os.Stat("user-data"); os.IsNotExist(err) {
		os.Mkdir("user-data", 0755)
	}

	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile("user-data/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	if _, err := file.WriteString(strconv.Itoa(rand.Int()) + "\n"); err != nil {
		log.Fatal(err)
	}

	//Print the contents of the file
	data, err := ioutil.ReadFile("user-data/temp.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	// fmt.Fprintf(w, "ok!")
	fmt.Fprintf(w, string(data))
}
