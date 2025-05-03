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

type Stats struct {
	ReqParamsHits []ReqData
}

func (req *ReqData) Create() error {

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
	// validate fb params according to the fizzbuzz game logic rules
	// e.g. str1 and str2 != ""
	log.Println("INFO: params we got that we would validate:", fb)
	return nil
}

func (st Stats) CountReqParamsHit(req ReqData) {
	log.Println("DEBUG: Enter stats count")
	if len(st.ReqParamsHits) == 0 {
		st.ReqParamsHits = []ReqData{
			{
				Int1:     req.Int1,
				Int2:     req.Int2,
				Limit:    req.Limit,
				Str1:     req.Str1,
				Str2:     req.Str2,
				HitCount: 1,
			},
		}
	}
	for i, reqStat := range st.ReqParamsHits {
		log.Println("DEBUG: enter stat range")
		if req.Int1 == reqStat.Int1 &&
			req.Int2 == reqStat.Int2 &&
			req.Limit == reqStat.Limit &&
			req.Str1 == reqStat.Str1 &&
			req.Str2 == reqStat.Str2 {
			st.ReqParamsHits[i].HitCount++
			log.Println("DEBUG: fb stats:", req)
			return
		}
		log.Println("DEBUG: stats not found")
		st.ReqParamsHits = append(st.ReqParamsHits,
			ReqData{
				Int1:     req.Int1,
				Int2:     req.Int2,
				Limit:    req.Limit,
				Str1:     req.Str1,
				Str2:     req.Str2,
				HitCount: 1,
			})
		log.Println("DEBUG: fb stats:", req)
	}
}
