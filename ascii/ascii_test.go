package asciiart_test

import (
	"testing"

	asciiart "asciiart/ascii"
)

func TestAsciiDrawer(t *testing.T) {
	res, _ := asciiart.AsciiDrawer("lll", "standard")
	t.Log(res)
}
