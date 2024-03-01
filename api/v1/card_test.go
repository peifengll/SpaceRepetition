package v1

import (
	"fmt"
	"github.com/open-spaced-repetition/go-fsrs"
	"testing"
)

func TestCardReviewRsp(t *testing.T) {
	c := CardReviewOptReq{
		Card: fsrs.Card{},
		ID:   0,
	}

	fmt.Printf("%#v", c)
}
