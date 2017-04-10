package validations

import (
	// Standard library packages
	"io"
	"bytes"
	"strings"
	"io/ioutil"
	"net/http"
	"encoding/json"

	// Third party packages
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"

	// Custom packages
	"github.com/rtbathula/golangapp/services"
)

func CreateNew(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var response services.Response
	response.Status  = "error"

	b       := bytes.NewBuffer(make([]byte, 0))
	reader  := io.TeeReader(r.Body, b)
	decoder := json.NewDecoder(reader)
	r.Body  = ioutil.NopCloser(b)

	var newCompany services.NewCompany
	err := decoder.Decode(&newCompany)

	if err != nil {
		response.Message = "Invalid params"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	//Validate fields
	newCompany.Name = strings.TrimSpace(newCompany.Name)
	if newCompany.Name == "" {
		response.Message = "Company name is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (newCompany.Name != "" && len(newCompany.Name)<2){
		response.Message = "Company name should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	newCompany.Address = strings.TrimSpace(newCompany.Address)
	if newCompany.Address == "" {
		response.Message = "Company address is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (newCompany.Address != "" && len(newCompany.Address)<2){
		response.Message = "Company address should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	newCompany.City = strings.TrimSpace(newCompany.City)
	if newCompany.City == "" {
		response.Message = "Company city is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	if (newCompany.City != "" && len(newCompany.City)<2){
		response.Message = "Company city should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	newCompany.Country = strings.TrimSpace(newCompany.Country)
	if newCompany.Country == "" {
		response.Message = "Company country is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (newCompany.Country != "" && len(newCompany.Country)<2){
		response.Message = "Company country should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	newCompany.Email = strings.TrimSpace(newCompany.Email)
	if(newCompany.Email!= "" && !govalidator.IsEmail(newCompany.Email)){
		response.Message = "Invalid company email"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	newCompany.Phone = strings.TrimSpace(newCompany.Phone)
	if(newCompany.Phone!= "" && len(newCompany.Phone)<9){
		response.Message = "Company phone should atleast of 9 digits"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if(len(newCompany.Directors)==0){     
        response.Message = "Atleast one company director is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
    }	

    var uniqueDirectors []string
    for _,director:=range newCompany.Directors {    

    	for _,uniqueEmail:=range uniqueDirectors {
    		if(uniqueEmail == director.Email){
    			response.Message = "Duplicate director email:"+uniqueEmail
				respByt,_ := json.Marshal(response)

				w.WriteHeader(400)
				w.Write(respByt)
				return
    		}
    	}    			
    	
    	director.Name = strings.TrimSpace(director.Name)
		if director.Name == "" {
			response.Message = "Director name is required"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}

		if (director.Name != "" && len(director.Name)<2){
			response.Message = "Director name should contain atleast of 2 letters"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}

		director.Email = strings.TrimSpace(director.Email)
		if(!govalidator.IsEmail(director.Email)){
			response.Message = "Invalid director email"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}    	

	    uniqueDirectors = append(uniqueDirectors,director.Email)
    }

    if(len(newCompany.Beneficials)==0){     
        response.Message = "Atleast one company beneficial is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
    }	

    var uniqueBeneficials []string
    for _,beneficial:=range newCompany.Beneficials {    

    	for _,uniqueEmail:=range uniqueBeneficials {
    		if(uniqueEmail == beneficial.Email){
    			response.Message = "Duplicate beneficial email:"+uniqueEmail
				respByt,_ := json.Marshal(response)

				w.WriteHeader(400)
				w.Write(respByt)
				return
    		}
    	}    			
    	
    	beneficial.Name = strings.TrimSpace(beneficial.Name)
		if beneficial.Name == "" {
			response.Message = "Beneficial name is required"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}

		if (beneficial.Name != "" && len(beneficial.Name)<2){
			response.Message = "Beneficial name should contain atleast of 2 letters"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}

		beneficial.Email = strings.TrimSpace(beneficial.Email)
		if(!govalidator.IsEmail(beneficial.Email)){
			response.Message = "Invalid beneficial email"
			respByt,_ := json.Marshal(response)

			w.WriteHeader(400)
			w.Write(respByt)
			return
		}    	

	    uniqueBeneficials = append(uniqueBeneficials,beneficial.Email)
    }

	next(w, r)
}

func GetDetails(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	params := mux.Vars(r)
	id := params["id"]

	var response services.Response
	response.Status  = "error"

	if id == "" {
		response.Message = "Company id required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if bson.IsObjectIdHex(id) ==false {
		response.Message = "Invalid company id"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	next(w, r)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var response services.Response
	response.Status  = "error"

	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		response.Message = "Company id required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if bson.IsObjectIdHex(id) ==false {
		response.Message = "Invalid company id"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	type Request struct {
		Address           string        `json:"address"`
        City              string        `json:"city"`
        Country           string        `json:"country"`
        Email             string        `json:"email"`
        Phone             string        `json:"phone"`
	}

	b       := bytes.NewBuffer(make([]byte, 0))
	reader  := io.TeeReader(r.Body, b)
	decoder := json.NewDecoder(reader)
	r.Body  = ioutil.NopCloser(b)

	req := Request{}
	err := decoder.Decode(&req)

	if err != nil {
		response.Message = "Invalid update company object"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	req.Address = strings.TrimSpace(req.Address)
	req.City    = strings.TrimSpace(req.City)
	req.Country = strings.TrimSpace(req.Country)
	req.Email   = strings.TrimSpace(req.Email)
	req.Phone   = strings.TrimSpace(req.Phone)

	//Validate fields
	if (req.Address == "" && req.City == "" && req.Country == "" && req.Email == "" && req.Phone == "") {
		response.Message = "Invalid update company object"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (req.Address != "" && len(req.Address)<2){
		response.Message = "Company address should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	if (req.City != "" && len(req.City)<2){
		response.Message = "Company city should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (req.Country != "" && len(req.Country)<2){
		response.Message = "Company country should contain atleast of 2 letters"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	if(req.Email!= "" && !govalidator.IsEmail(req.Email)){
		response.Message = "Invalid company email"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	if(req.Phone!= "" && len(req.Phone)<9){
		response.Message = "Company phone should atleast of 9 digits"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}	

	next(w, r)
}