package app

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"../dao"
	"../model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/api/voters", a.GetAllVoters)
	a.Post("/api/voter", a.CreateVoter)
	a.Post("/api/voters", a.CreateVoters)
	a.Get("/api/voters/{voter}", a.GetVoter)
	a.Put("/api/voters/{voter}", a.UpdateVoter)
	a.Delete("/api/voters/{voter}", a.DeleteVoter)
	//a.Put("/employees/{title}/enable", a.EnableEmployee)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a *App) GetAllVoters(w http.ResponseWriter, r *http.Request) {
	dao.GetAllVoters(a.DB, w, r)
}

func (a *App) CreateVoter(w http.ResponseWriter, r *http.Request) {
	dao.CreateVoter(a.DB, w, r)
}

func (a *App) CreateVoters(w http.ResponseWriter, r *http.Request) {
	dao.CreateVoters(a.DB, w, r)
}

func (a *App) GetVoter(w http.ResponseWriter, r *http.Request) {
	dao.GetVoter(a.DB, w, r)
}

func (a *App) UpdateVoter(w http.ResponseWriter, r *http.Request) {
	dao.UpdateVoter(a.DB, w, r)
}

func (a *App) DeleteVoter(w http.ResponseWriter, r *http.Request) {
	dao.DeleteVoter(a.DB, w, r)
}

//func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
//	dao.EnableEmployee(a.DB, w, r)
//}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
