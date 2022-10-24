package dups

import (
	"fmt"
	"log"

	"github.com/msolimans/dups/pkg/core"
	"github.com/msolimans/dups/pkg/file"
)

func Start(source, destination string, concurrent, pretty, out bool) {
	
	res := core.DBSchema{}
	
	//read source file and marshal/store in res
	if err := file.UnMarshalFile(source, &res); err != nil {
		log.Fatal("Error reading/unmarshalling source file, \n details:", err)
	}

	//record current counts 
	fieldCount := res.FieldCount
	viewCount := res.ViewCount
	
	//remove duplicates
	
	res.RemoveDuplicates(concurrent)

	//write to a new file
	if err := file.MarshalToFile(destination, pretty, res); err != nil {
		log.Fatal("Error saving to destination file \n details: ", err)
	}

	//output summary
	if out {
		fmt.Println("\n ---------- Summary ------------")
		fmt.Printf("\nField count: \n\t before: %d \n\t after: %d", fieldCount, res.FieldCount)
		fmt.Printf("\nView count: \n\t before: %d \n\t after: %d", viewCount, res.ViewCount)
	}

	fmt.Println("\nDone!")

}