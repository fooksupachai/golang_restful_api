package controller

import (
	"net/http"

	m "github.com/fooksupachai/golang_restful_api/model"
	u "github.com/fooksupachai/golang_restful_api/utils"
)

// RiskManagement to provide risk value
func RiskManagement(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	keyRisk, okRisk := r.URL.Query()["risk"]
	keyAmount, okAmount := r.URL.Query()["amount"]
	keyPerson, okPerson := r.URL.Query()["person"]
	keyDepartment, okDepartment := r.URL.Query()["department"]

	if !okRisk || len(keyRisk[0]) < 0 {
		u.ValiateQueryURL(w, "risk")
		return
	}

	if !okAmount || len(keyAmount[0]) < 0 {
		u.ValiateQueryURL(w, "amount")
		return
	}

	if !okPerson || len(keyPerson[0]) < 0 {
		u.ValiateQueryURL(w, "person")
		return
	}

	if !okDepartment || len(keyDepartment[0]) < 0 {
		u.ValiateQueryURL(w, "department")
		return
	}

	m.ProvideCalcPrice()

}
