package configuration

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "reflect"
)

type Configuration struct {
    SomethingRequired string
    SomethingNotRequired string
}

// denotes required/optional fields
var fieldStatus = map[string]bool{
    "SomethingRequired" : true,
    "SomethingNotRequired" : false,
}

// constants that should be replaced by configuration file written by cookbook!
const(
    DEFAULT_CONFIG_FILE="/opt/awesome/config/config.json"
)

var Config *Configuration

// ReadConfig reads in the configuration file
func ReadConfig(configFile string) error {
    // read the file
    var err error
    var data []byte
    log.Printf("Reading configuration file at %s", configFile)
    data, err = ioutil.ReadFile(configFile)
    if err != nil {
        log.Println(err.Error())
        return err
    }

    // construct the config from the data
    err = construct(data)
    if err != nil {
        return err
    }

    log.Println("Configuration loaded!")
    return nil
}

func construct(data []byte) error {
    // construct
    Config = &Configuration{}
    err := json.Unmarshal(data, Config)
    if err != nil {
        log.Println(err.Error())
        return err
    }

    // validate
    elem := reflect.TypeOf(Config).Elem()
    value := reflect.ValueOf(*Config)
    for i := 0; i < elem.NumField(); i++ {
        field := elem.Field(i)
        if fieldStatus[field.Name] {
            fieldValue := value.Field(i)
            if len(fieldValue.String()) == 0 {
                return errors.New(fmt.Sprintf("[ERROR] Configuration required field missing: %s", field.Name))
            }
        }
    }

    return nil
}