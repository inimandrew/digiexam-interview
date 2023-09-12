package services

import (
	"digiexam-interview/src/responses"
	"math"
)

func GetPaginationMeta(total int64, resultCount int64,limit int64,offset int64) responses.PaginateMeta {
	var divisor int64
	var firstPage int64
	from := offset

	if resultCount > 0 {
		firstPage = 1;
	}

	if offset == 0 {
		from = 1
	}

	if total == limit {
		divisor = int64(resultCount)
	}else {
		divisor = int64(limit)
	}

    return responses.PaginateMeta{
      FirstPage: firstPage,
      LastPage: int64(math.Ceil(float64(total) / float64(divisor))),
      Total: total,
      From: from,
      To: offset + resultCount,
    };
}