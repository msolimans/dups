package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func UnMarshalFile(fileName string, result interface{}) error {

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error reading file named", fileName, err.Error())
	}

	err = json.Unmarshal(raw, result)

	return err
}

func MarshalToFile(fileName string, pretty bool, any interface{}) (error) {
	var bte []byte 
	var err error
	
	if pretty {
		bte, err = json.MarshalIndent(any, "", "  ")
	} else {
		bte, err = json.Marshal(any)
	}
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, bte, 0644)
}


