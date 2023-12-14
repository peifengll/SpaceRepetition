package test

import (
	"fmt"
	"testing"
)

func TestSprintf(t *testing.T) {
	con := fmt.Sprintf("%%%s%%", "content")
	fmt.Println(con)
}
