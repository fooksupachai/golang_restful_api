package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/fooksupachai/golang_restful_api/database"
	m "github.com/fooksupachai/golang_restful_api/model"
)

// Risk for risk handler
type Risk struct {
	Manufactory float64 `json:"manufactory"`
	Person      float64 `json:"person"`
	Government  float64 `json:"government"`
}

// RiskManagement to provide risk value
func RiskManagement(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	M, P, G := m.ProvideCalcPrice(r.Body)

	fmt.Println("Contoller do")
	fmt.Println(M, P, G)

	database.InsertRiskData(M, P, G)

	resp := struct {
		Status int `json:"status"`
	}{
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)

}

// GetRiskInfomation to provide data of risk
func GetRiskInfomation(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// keyRisk, okRisk := r.URL.Query()["risk"]
	// keyAmount, okAmount := r.URL.Query()["amount"]
	// keyPerson, okPerson := r.URL.Query()["person"]
	// keyDepartment, okDepartment := r.URL.Query()["department"]

	// if !okRisk || len(keyRisk[0]) < 0 {
	// 	u.ValiateQueryURL(w, "risk")
	// 	return
	// }

	// if !okAmount || len(keyAmount[0]) < 0 {
	// 	u.ValiateQueryURL(w, "amount")
	// 	return
	// }

	// if !okPerson || len(keyPerson[0]) < 0 {
	// 	u.ValiateQueryURL(w, "person")
	// 	return
	// }

	// if !okDepartment || len(keyDepartment[0]) < 0 {
	// 	u.ValiateQueryURL(w, "department")
	// 	return
	// }

	result := database.GetAllRiskData()

	var risks []Risk

	for result.Next() {

		var risk Risk

		result.Scan(
			&risk.Manufactory,
			&risk.Person,
			&risk.Government,
		)

		risks = append(risks, risk)
	}

	resp := struct {
		Risk   []Risk `json:"risk"`
		Status int    `json:"status"`
	}{
		Risk:   risks,
		Status: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(resp)

}
