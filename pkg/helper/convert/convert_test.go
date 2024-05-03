package convert

import (
	"log"
	"testing"
)

func TestParseStringToFloatSlice(t *testing.T) {
	cc := "0.4,     0.6, 2.4, 5.8, 4.93, 0.94, 0.86, 0.01, 1.49, 0.14, 0.94, 2.18, 0.05, 0.34, 1.26, 0.29, 2.61"
	slice, err := ParseStringToFloatSlice(cc)
	if err != nil {
		return
	}
	log.Println(slice)
	log.Println(len(slice))
}
