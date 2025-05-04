package main

import (
	"fmt"
	"os"

	fb "github.com/jaimegago/lbc-modular/pkg/fizzbuzz"
)

func main() {
	// hardcoding some input just to show modularity of fizzbuzz code
	req := fb.ReqData{
		Int1:  3,
		Int2:  5,
		Limit: 30,
		Str1:  "foo",
		Str2:  "bar",
	}
	err := req.Validate()
	if err != nil {
		fmt.Println("params failed ", err)
		os.Exit(1)
	}
	err = req.Get()
	if err != nil {
		fmt.Println("fizzbuzz failed ", err)
		os.Exit(1)
	}
	fmt.Println(req.Results)
}
