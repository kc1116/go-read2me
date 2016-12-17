package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kc1116/go-read2me/parse"
	"github.com/kc1116/go-read2me/read"
)

func main() {
	/*text, err := translate.TransText("es", "Hello World")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(text)*/
	txtParser := new(parse.TxtParser)
	//txtParser.Parse("./parse/test.txt")
	//read.GetVoices("en-US")
	read.Read(txtParser.Parse("./parse/test.txt"), "Kendra")

	//router := routing.NewRouter()
	//log.Fatal(http.ListenAndServe("3000", &MyServer{router}))

}

//MyServer type is a struct containing a pointer to a mux.router object
type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, AllowCredentials")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
