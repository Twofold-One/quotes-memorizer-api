package main

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/", app.getHelloWorld).Methods("GET")
	// app.Router.HandleFunc("/getfour", app.getFour).Methods("GET")
}