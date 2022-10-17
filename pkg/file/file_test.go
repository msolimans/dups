package file

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var workingDir string

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	load()
}

func load() {
	_, filename, _, _ := runtime.Caller(1)
	workingDir = path.Dir(filename)
}


func TestJsonMarshal(t *testing.T) {
	
	var test = &Test{Id: 10, Name: "test"}
	
	err := MarshalToFile(workingDir+"/test.json", false, test)
	assert.Nil(t, err)
	assert.True(t, FileExists(workingDir+"/test.json"))
	
	test2 := &Test{}
	err = UnMarshalFile(workingDir + "/test.json", test2)
	assert.Nil(t, err)
	assert.NotNil(t, test2)
	assert.Equal(t, 10, test2.Id)

	//cleanup (even if test fails, it will be deleted in the next run)
	assert.Equal(t, nil,os.Remove(workingDir + "/test.json"))
}
