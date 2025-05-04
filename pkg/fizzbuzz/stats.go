package fizzbuzz

import (
	"sort"
)

func (req *ReqData) CountReqParamsHit(stats []ReqData) []ReqData {
	for i, reqStat := range stats {
		if req.Int1 == reqStat.Int1 &&
			req.Int2 == reqStat.Int2 &&
			req.Limit == reqStat.Limit &&
			req.Str1 == reqStat.Str1 &&
			req.Str2 == reqStat.Str2 {
			stats[i].HitCount++
			return stats
		}
	}
	stats = append(stats,
		ReqData{
			Int1:     req.Int1,
			Int2:     req.Int2,
			Limit:    req.Limit,
			Str1:     req.Str1,
			Str2:     req.Str2,
			HitCount: 1,
		})
	return stats
}

func GetHighestHitCount(stats []ReqData) ReqData {
	if len(stats) < 1 {
		return ReqData{}
	}
	sort.SliceStable(stats, func(i, j int) bool {
		return stats[i].HitCount > stats[j].HitCount
	})
	return stats[0]
}
