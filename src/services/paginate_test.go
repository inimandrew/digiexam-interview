package services_test

import (
	"digiexam-interview/src/responses"
	"digiexam-interview/src/services"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type PaginationTestInput struct {
	total, resultCount, limit, offset int64
	result responses.PaginateMeta
}

func TestGetPaginationMeta(t *testing.T){
	testData := []PaginationTestInput{
		{
		20, 10,10,0,
		responses.PaginateMeta{
		FirstPage: 1,
		LastPage: 2,
		Total: 20,
		From: 1,
		To: 10,
	},
		},
		{
		50, 5,5,0,
		responses.PaginateMeta{
		FirstPage: 1,
		LastPage: 10,
		Total: 50,
		From: 1,
		To: 5,
	},
		},
		{
		5, 5,5,0,
		responses.PaginateMeta{
		FirstPage: 1,
		LastPage: 1,
		Total: 5,
		From: 1,
		To: 5,
	},
		},
	}

	for _, datum := range testData {
		meta := services.GetPaginationMeta(datum.total, datum.resultCount, datum.limit, datum.offset);

			if cmp.Equal(meta, datum.result) {
		t.Logf("GetPaginationMeta(%d,%d,%d,%d) passed successfully", datum.total, datum.resultCount,datum.limit,datum.offset)
	}else{
		t.Errorf("GetPaginationMeta(%d,%d,%d,%d) failed", datum.total, datum.resultCount,datum.limit,datum.offset)
	}
	}

}