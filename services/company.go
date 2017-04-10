package services

import (  
    // Standard library packages 
    "fmt" 
    "time"  
    "strings"  
    "errors"

    // Third party packages
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"

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

    var (
        response Response
        err error
    )    

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close()    

    newCompany.Name = strings.TrimSpace(newCompany.Name)

    var company Company
    col:=sessionCopy.DB(databaseName).C("company") 
   
    err = col.Find(bson.M{"name":newCompany.Name}).One(&company)
    if (err == nil && company.Name!=""){
        response.Status  = "error" 
        response.Message = "Company with given name is already exist."
        return response,errors.New("Company with given name is already exist.")
    }   
   
    company.Id        = bson.NewObjectId() 
    company.CreatedAt = time.Now().Unix()
    company.UpdatedAt = time.Now().Unix() 

    company.Name        = strings.TrimSpace(newCompany.Name)
    company.Address     = newCompany.Address
    company.City        = newCompany.City
    company.Country     = newCompany.Country
    company.Directors   = newCompany.Directors
    company.Beneficials = newCompany.Beneficials
       
	err= col.Insert(company)    
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


/*Desc   : Get company details
  Params : company id
  Returns: Promise
           Resolve->company details
           Reject->Error on find() or document not found
*/
func GetDetails(id string) (Response,error) { 

    var (
        response Response
        err error
    )    

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close()      

    var company Company
    col:=sessionCopy.DB(databaseName).C("company") 
   
    err = col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&company)
    if (err!= nil){
        response.Status  = "error" 
        response.Message = "Unable to find company with given company id. Please check company id!"
        return response,err
    } 
   
    response.Status  = "success" 
    response.Message = "Successfully fetched the company details"
    response.Data    = company
    return response,nil        
}


/*Desc   : Get company list
  Params : skip,limit
  Returns: Promise
           Resolve->company list
           Reject->Error on find()
*/
func GetList(skip int,limit int) (Response,error) { 

    var (
        response Response
        err error
    )    

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close()      

    var list []Company
    col:=sessionCopy.DB(databaseName).C("company") 
   
    err = col.Find(bson.M{}).Limit(limit).Skip(skip).All(&list)  
    if (err!= nil){
        response.Status  = "error" 
        response.Message = "Unable to get the company list."
        return response,err
    } 
   
    response.Status  = "success" 
    response.Message = "Successfully fetched the company list"
    response.Data    = list
    return response,nil        
}


/*Desc   : Update company info
  Params : companyId, {address,city,country,email,phone}
  Returns: Promise
           Resolve->new company
           Reject->Error on findOneAndUpdate()
*/
func UpdateCompany(id string,address string,city string,country string,email string,phone string) (Response,error) { 

    var (
        response Response
        err error
    )    

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close()      

    var updateObj bson.M
    updateObj = make(map[string]interface {})
    updateObj["updatedAt"] = time.Now().Unix() 

    if(address!=""){
        updateObj["address"] = address
    } 

    if(city!=""){
        updateObj["city"]    = city
    } 

    if(country!=""){
        updateObj["country"] = country
    } 

    if(email!=""){
        updateObj["email"]   = email
    } 

    if(phone!=""){
        updateObj["phone"]   = phone
    }    

    newSet := mgo.Change{
        Update: bson.M{"$set": updateObj},
        ReturnNew: true,
    }

    var company Company
    col:=sessionCopy.DB(databaseName).C("company") 
         
    _, err = col.Find(bson.M{"_id":bson.ObjectIdHex(id)}).Apply(newSet, &company)
    if (err!= nil){
        response.Status  = "error" 
        response.Message = "Unable to update the company with given company id and update object."
        return response,err
    } 
   
    response.Status  = "success" 
    response.Message = "Successfully update the company"
    response.Data    = company
    return response,nil        
}

/*Desc   : Add beneficial
  Params : companyId, {name,email}
  Returns: Promise
           Resolve->new company
           Reject->Error on findOneAndUpdate()
*/
func AddBeneficial(id string,name string,email string) (Response,error) { 

    var (
        response Response
        err error
    )    

    //Get databaseName
    keysJson      := helpers.GetConfigKeys()
    envVariable   := helpers.GetEnvVariable()
    databaseName,_:=keysJson.String(envVariable,"databaseName") 

    //Get mongoSession
    mangoSession:=databases.GetMongoSession()
    sessionCopy := mangoSession.Copy()
    defer sessionCopy.Close()      

    var company Company
    col:=sessionCopy.DB(databaseName).C("company") 

    checkBenficialExistQuery:=bson.M{"$elemMatch": bson.M{"email" : email}}
    query:=bson.M{"_id": bson.ObjectIdHex(id),"beneficials": checkBenficialExistQuery}

    err = col.Find(query).One(&company)
    if (err == nil && company.Name!=""){
        response.Status  = "error" 
        response.Message = "Beneficial already exist with email:"+ email  
        return response,errors.New("Beneficial already exist with email:"+ email)
    } 

    var updateObj bson.M
    updateObj = make(map[string]interface {})
    updateObj["updatedAt"] = time.Now().Unix() 
    
    pushObj:= bson.M{"beneficials": bson.M{"name":name,"email":email}}       

    newSet := mgo.Change{
        Update: bson.M{"$set": updateObj,"$push": pushObj},
        ReturnNew: true,
    }    
         
    _, err = col.Find(bson.M{"_id":bson.ObjectIdHex(id)}).Apply(newSet, &company)
    if (err!= nil){
        response.Status  = "error" 
        response.Message = "Unable to add the beneficial with given company id and beneficial object."
        return response,err
    } 
   
    replyData:=map[string]string{
        "name"  : name,
        "email" : email,
    }

    response.Status  = "success" 
    response.Message = "Successfully added the new beneficial"
    response.Data    = replyData
    return response,nil        
}