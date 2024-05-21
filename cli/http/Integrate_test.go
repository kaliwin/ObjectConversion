package http

import "testing"

func TestDiversion(t *testing.T) {

	Diversion("/root/tmp/data", "", "/root/tmp/dir")

}
