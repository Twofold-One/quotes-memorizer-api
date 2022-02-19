package main

import (
	"fmt"
	"net/http"
)

func (app *App) getHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	fmt.Fprint(w,"<h1>Hello World!</h1>")
}

// func (app *App) getFour(w http.ResponseWriter, r *http.Request) {
// 	n, err := app.quotes.Checker()
// 	if err != nil {
// 		fmt.Println("no DB connection")
// 	}
// 	fmt.Fprintf(w,"<h1>%d</h1>", n)
// }