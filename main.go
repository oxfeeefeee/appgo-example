package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	//"github.com/codegangsta/negroni"
	//_ "github.com/jinzhu/configor"
	//_ "github.com/jinzhu/gorm"
	"github.com/oxfeeefeee/appgo"
	"github.com/oxfeeefeee/appgo/server"
	//_ "github.com/parnurzeal/gorequest"
	"net/http"
)

func main() {
	log.Infoln("starting server...")

	s := server.NewServer(auth{})

	s.AddHandler("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to appgo")
	})

	s.AddRest("/api", []interface{}{
		&Hello{},
	})

	s.Serve()
}

type auth struct{}

func (_ auth) Validate(token string) (userId appgo.Id, role appgo.Role) {
	return 1, appgo.RoleAppUser
}
