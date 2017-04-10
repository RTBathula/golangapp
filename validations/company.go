package validations

import (
	// Standard library packages
	"io"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"

	// Third party packages
	_"github.com/gorilla/mux"
	_"gopkg.in/mgo.v2/bson"

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
	if newCompany.Name == "" {
		response.Message = "Company name is required"
		respByt,_ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}	

	next(w, r)
}
