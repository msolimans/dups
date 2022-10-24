package core

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/msolimans/dups/pkg/file"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
    In *DBSchema
	Concurrent bool
    Exp *DBSchema
}

func Test(t *testing.T) {
    tests := []testCase{
        {In: nil, Exp: nil},
    }

	files, err := ioutil.ReadDir("./tests")
    assert.Nil(t, err)

    for _, f := range files {
        if !f.IsDir() && strings.ToLower(f.Name()[0:2]) == "in." { 
			res := DBSchema{}
			exp := DBSchema{}
			err := file.UnMarshalFile("./tests/" + f.Name(), &res)
			assert.Nil(t, err)
			err = file.UnMarshalFile(fmt.Sprintf("./tests/OUT%s", f.Name()[2:]), &exp)
			assert.Nil(t, err)

			tests = append(tests, testCase{
				In: &res,
				Exp: &res,
			})
		}
    }

	//parse the schema and remove dups (compare with the clean one)
    for _, test := range tests {
        test.In.RemoveDuplicates(true)
		assert.Equal(t, test.Exp,test.In)

		test.In.RemoveDuplicates(false)
		assert.Equal(t, test.Exp,test.In)
    }
}
