package main

import (  
    //Standard library packages
    "fmt" 
    "net/http"

    //Third party packages  
    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/rs/cors"

    //Custom packages      
    "github.com/rtbathula/golangapp/api"   
    "github.com/rtbathula/golangapp/helpers"
    "github.com/rtbathula/golangapp/databases"    
)

func main() {     

    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Clearhaus API golangapp up and running")
    })

    n := negroni.Classic() // Includes some default middlewares
    n.Use(cors.New(cors.Options{
        AllowedOrigins   : []string{"*"},
        AllowedMethods   : []string{"GET","POST","DELETE","PUT", "PATCH"},
        AllowedHeaders   : []string{"Origin","Authorization","X-Requested-With","Content-Type","Accept"},
        ExposedHeaders   : []string{"Content-Length"},
        AllowCredentials : true,
    }))
    n.UseHandler(r)

    //Connect MongoDB
    connectMongoDB();    

    //Routes
    routes(r)

    //Run server
    http.Handle("/", r)   
    http.ListenAndServe(helpers.GetPortAddress(), n)
}


//Private Fuctions
func connectMongoDB(){      
    databases.ConnectDB()    
}

func routes(r *mux.Router){
    api.CompanyApi(r)           
}
