package v1

import (
	"fmt"
	"testing"
)

func TestCardReviewRsp(t *testing.T) {
	c := CardReviewOptReq{
		ID: 0,
	}

	fmt.Printf("%#v", c)
}
