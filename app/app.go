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

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

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

func (a *App) setRouters() {
	a.Get("/api/isConnect", a.IsConnect)
	a.Get("/api/getVotes", a.GetVotes)
	a.Post("/api/createVotes", a.CreateVotes)
	a.Post("/api/createVote", a.CreateVote)
	a.Get("/api/voters/{voter}", a.GetVote)
	a.Put("/api/voters/{voter}", a.UpdateVote)
	a.Delete("/api/voters/{voter}", a.DeleteVote)
	//a.Put("/employees/{title}/enable", a.EnableEmployee)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) IsConnect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server started")
}

func (a *App) GetVotes(w http.ResponseWriter, r *http.Request) {
	dao.GetVotes(a.DB, w, r)
}

func (a *App) CreateVote(w http.ResponseWriter, r *http.Request) {
	dao.CreateVote(a.DB, w, r)
}

func (a *App) CreateVotes(w http.ResponseWriter, r *http.Request) {
	dao.CreateVotes(a.DB, w, r)
}

func (a *App) GetVote(w http.ResponseWriter, r *http.Request) {
	dao.GetVote(a.DB, w, r)
}

func (a *App) UpdateVote(w http.ResponseWriter, r *http.Request) {
	dao.UpdateVote(a.DB, w, r)
}

func (a *App) DeleteVote(w http.ResponseWriter, r *http.Request) {
	dao.DeleteVote(a.DB, w, r)
}

//func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
//	dao.EnableEmployee(a.DB, w, r)
//}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
