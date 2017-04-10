package api

import (
	// Standard library packages
	"encoding/json"
	"net/http"

	// Third party packages
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	//Custom packages
	"github.com/rtbathula/golangapp/services"
	"github.com/rtbathula/golangapp/validations"
)

// *****************************************************************************
// API Routes
// *****************************************************************************

func CompanyApi(r *mux.Router) {
	
	r.Handle("/company",negroni.New(
		negroni.HandlerFunc(validations.CreateNew),
		negroni.Wrap(CreateNew),
	)).Methods("POST")

	/*r.HandleFunc("/company/{id}", getDetails).Methods("GET")
	r.HandleFunc("/company", getList).Methods("GET")	
	r.HandleFunc("/company/{id}/update-company", updateCompany).Methods("PUT")	
	r.HandleFunc("/company/{id}/add-beneficial", addBeneficial).Methods("PUT")	*/
}


var CreateNew = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var newCompany services.NewCompany
	decoder := json.NewDecoder(r.Body)	
	decoder.Decode(&newCompany)

	resp, err := services.CreateNew(newCompany)
	respByt,_:= json.Marshal(resp)

	if err != nil {		
		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	w.WriteHeader(200)
	w.Write(respByt)
	return
})
