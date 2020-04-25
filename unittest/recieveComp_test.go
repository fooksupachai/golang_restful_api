package testrecievecomp

import (
	"testing"

	m "github.com/fooksupachai/golang_restful_api/model"
)

// TestRecieveCompany for test amount of recieve
func TestRecieveCompany(t *testing.T) {

	amount := m.CalcReceiveComp(20)

	if 20 != amount {
		t.Error("Amount of recieve componany should be 20 but have", amount)
	}

}

// TestRecieveClient for test amount of recieve
func TestRecieveClient(t *testing.T) {

	amount := m.CalcReceiveComp(10)

	if 20 != amount {
		t.Error("Amount of recieve client should be 20 but have", amount)
	}

}
