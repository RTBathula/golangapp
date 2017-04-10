package databases

import (  
    //Standard library packages
    //Third party packages  
    "gopkg.in/mgo.v2"  

    //Custom packages  
    "github.com/rtbathula/golangapp/helpers"    
)

var mangoSession *mgo.Session

func ConnectDB(){

    keysJson    := helpers.GetConfigKeys()
    envVariable := helpers.GetEnvVariable()
        
    dbConnection,keyErr:=keysJson.String(envVariable,"mongoDBConnection")
    if keyErr != nil {
        panic(keyErr)
    }   

    //Connect to our mongo
    s, err := mgo.Dial(dbConnection)    
    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }    
    mangoSession=s       
}

func GetMongoSession() *mgo.Session{
    return mangoSession
}
