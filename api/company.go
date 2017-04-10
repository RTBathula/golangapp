package api

import (
	// Standard library packages
	"encoding/json"
	"net/http"
	"strconv"

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

	r.Handle("/company/{id}",negroni.New(
		negroni.HandlerFunc(validations.GetDetails),
		negroni.Wrap(GetDetails),
	)).Methods("GET")

	r.Handle("/company",negroni.New(
		negroni.Wrap(GetList),
	)).Methods("GET")

	r.Handle("/company/{id}/update-company",negroni.New(
		negroni.HandlerFunc(validations.UpdateCompany),
		negroni.Wrap(UpdateCompany),
	)).Methods("PUT")

	/*	
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

var GetDetails = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	resp, err := services.GetDetails(id)
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

var GetList = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
	var (
		skip int   = 0 //default
		limit int  = 999 //default	
	)

	qs := r.URL.Query()	
	reqSkip, reqLimit:= qs.Get("skip"), qs.Get("limit")	

	if reqSkip != "" {
		skip, _ = strconv.Atoi(reqSkip)
	}

	if reqLimit != "" {
		limit, _ = strconv.Atoi(reqLimit)
	}	

	resp, err := services.GetList(skip,limit)
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


var UpdateCompany = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	type Request struct {
		Address           string        `json:"address"`
        City              string        `json:"city"`
        Country           string        `json:"country"`
        Email             string        `json:"email"`
        Phone             string        `json:"phone"`
	}

	var req Request
	decoder := json.NewDecoder(r.Body)	
	decoder.Decode(&req)

	resp, err := services.UpdateCompany(id,req.Address,req.City,req.Country,req.Email,req.Phone)
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