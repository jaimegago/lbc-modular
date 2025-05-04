package fizzbuzz

import (
	"log"
	"strconv"
)

type ReqData struct {
	Int1     int    `json:"int1"`
	Int2     int    `json:"int2"`
	Limit    int    `json:"limit"`
	Str1     string `json:"str1"`
	Str2     string `json:"str2"`
	Results  []string
	HitCount int
}

func (req *ReqData) Get() error {

	results := []string{}
	for i := 0; i <= req.Limit; i++ {
		if i%req.Int1 == 0 && i%req.Int2 == 0 {
			results = append(results, req.Str1+req.Str2)
		} else if i%req.Int1 == 0 {
			results = append(results, req.Str1)
		} else if i%req.Int2 == 0 {
			results = append(results, req.Str2)
		} else {
			results = append(results, strconv.Itoa(i))
		}
	}
	req.Results = results
	return nil
}

func (fb *ReqData) Validate() error {
	// Here we would validate fb params according to the fizzbuzz game logic rules
	// e.g. str1 and str2 != ""
	// (done in previous "non modular" version)
	log.Println("INFO: params we got that we would validate:", fb)
	return nil
}
