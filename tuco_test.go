package tuco_test

import (
	"github.com/mcandre/tuco"

	"testing"
)

func TestPortMarshalingSymmetric(t *testing.T) {
	portString := "linux/riscv64"
	port, err := tuco.ParsePort(portString)

	if err != nil {
		t.Error(err)
	}

	portString2 := port.String()

	if portString2 != portString {
		t.Errorf("expected symmetric marshaling for port %v", portString)
	}
}
