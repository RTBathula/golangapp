package services

import (  
    // Standard library packages 
    "fmt" 
    "time"    

    // Third party packages
    "gopkg.in/mgo.v2/bson"

    //Custom packages     
    "github.com/rtbathula/golangapp/databases" 
    "github.com/rtbathula/golangapp/helpers" 
)

type (  
     Company struct { 
        Id                bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
        CreatedAt         int64     `json:"createdAt" bson:"createdAt"`
        UpdatedAt         int64     `json:"updatedAt" bson:"updatedAt"`   

        Name              string        `json:"name" bson:"name"` 
        Address           string        `json:"address" bson:"address"`
        City              string        `json:"city" bson:"city"`
        Country           string        `json:"country" bson:"country"`
        Email             string        `json:"email" bson:"email"`
        Phone             string        `json:"phone" bson:"phone"`       
        
        Directors          []Director     `json:"directors" bson:"directors"`
        Beneficials        []Beneficial  `json:"beneficials" bson:"beneficials"`                    
    }

    NewCompany struct { 
        Name              string        `json:"name" bson:"name"` 
        Address           string        `json:"address" bson:"address"`
        City              string        `json:"city" bson:"city"`
        Country           string        `json:"country" bson:"country"`
        Email             string        `json:"email" bson:"email"`
        Phone             string        `json:"phone" bson:"phone"`       
        
        Directors          []Director     `json:"directors" bson:"directors"`
        Beneficials        []Beneficial  `json:"beneficials" bson:"beneficials"`                         
    }

    Director struct {
        Name              string        `json:"name" bson:"name"`
        Email             string        `json:"email" bson:"email"`
    }

    Beneficial struct {
        Name              string        `json:"name" bson:"name"`
        Email             string        `json:"email" bson:"email"`
    }

    Response struct {
        Status      string      `json:"status"`
        Message     string      `json:"message"`
        Data        interface{} `json:"data"`
    }
)


// *****************************************************************************
// Model Methods
// *****************************************************************************

/*Desc   : Create new company
  Params : {name,address,city,country,email,phone,directors[],beneficials[]}
  Returns: Promise
           Resolve->saved success message
           Reject->Error on find() or company name already exist or save()
*/
func CreateNew(newCompany NewCompany) (Response,error) { 

    var response Response

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close() 
   
    var company Company
    company.Id        = bson.NewObjectId() 
    company.CreatedAt = time.Now().Unix()
    company.UpdatedAt = time.Now().Unix() 

    company.Name        = newCompany.Name
    company.Address     = newCompany.Address
    company.City        = newCompany.City
    company.Country     = newCompany.Country
    company.Directors   = newCompany.Directors
    company.Beneficials = newCompany.Beneficials
 

    col:=sessionCopy.DB(databaseName).C("company")    
	err:= col.Insert(company)
    
    if err!= nil { 
        fmt.Println(err)        
        response.Status  = "error" 
        response.Message = "Something went wrong. Please try after sometime"
        return response,err
    }
   
    response.Status  = "success" 
    response.Message = "Successfully created a company"
    response.Data    = company
    return response,nil        
}
