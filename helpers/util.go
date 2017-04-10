package helpers
import (
    // Standard library packages    
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"    

    // Third party packages
    "github.com/jmoiron/jsonq"    
)

func GetConfigKeys()  *jsonq.JsonQuery{
    file, err := ioutil.ReadFile("./config/keys.json")
    if err != nil {
        fmt.Println(err)
    }     
    var dat map[string] interface{}
    json.Unmarshal(file, &dat)

    jq :=jsonq.NewQuery(dat)
    return jq
}

func IsProduction() bool {

    port := os.Getenv("PORT")
    if len(port) == 0 {
        return  false
    } 

    return true  
}

func GetEnvVariable() string {

    isProd:=IsProduction()
    if isProd {
      return "production"
    }

    return "development" 
}

func GetPortAddress() string {

    var PORT string = ":3000" // If not found in env
    
    envport := os.Getenv("PORT")
    if envport != "" {
        PORT = ":" +os.Getenv("PORT")
    }

    return PORT
}
