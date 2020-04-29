package model

import (
	"encoding/json"
	"fmt"
	"io"
)

// PriceProvider to return computation price
type PriceProvider interface {
	ComputePrice() float64
}

// PriceManufactory to provide price structure
type PriceManufactory struct {
	amount float64
	risk   float64
}

// PricePerson to provide price structure
type PricePerson struct {
	person float64
	amount float64
	risk   float64
}

// PriceGovernment to provide price structure
type PriceGovernment struct {
	department float64
	person     float64
	amount     float64
	risk       float64
}

// Price to provide model structure
type Price struct {
	Department float64 `json:"department"`
	Person     float64 `json:"person"`
	Amount     float64 `json:"amount"`
	Risk       float64 `json:"risk"`
}

// ComputePrice to calculate type manufactory
func (model PriceManufactory) ComputePrice() float64 {

	return model.amount * model.risk
}

// ComputePrice  to calculate type person
func (model PricePerson) ComputePrice() float64 {

	return model.amount * model.person * model.risk
}

// ComputePrice to calculate type government
func (model PriceGovernment) ComputePrice() float64 {

	return model.amount * model.department * model.person * model.risk
}

// CalcPriceManufactory to provide amount of client price value
func CalcPriceManufactory(provider PriceProvider, resultManufactory chan<- float64) {

	resultManufactory <- provider.ComputePrice()
}

// CalcPricePerson to provide amount of client price value
func CalcPricePerson(provider PriceProvider, resultPerson chan<- float64) {

	resultPerson <- provider.ComputePrice()
}

// CalcPriceGovernment to provide amount of client price value
func CalcPriceGovernment(provider PriceProvider, resultGovernment chan<- float64) {

	resultGovernment <- provider.ComputePrice()
}

// RecieveCalcPriceManufactory to return value in channel
func RecieveCalcPriceManufactory(ch <-chan float64) float64 {
	return <-ch
}

// RecieveCalcPricePerson to return value in channel
func RecieveCalcPricePerson(ch <-chan float64) float64 {
	return <-ch
}

// RecieveCalcPriceGovernment to return value in channel
func RecieveCalcPriceGovernment(ch <-chan float64) float64 {
	return <-ch
}

// ProvideCalcPrice to provide amount to calculate
func ProvideCalcPrice(body io.ReadCloser) (M float64, P float64, G float64) {

	var price Price

	err := json.NewDecoder(body).Decode(&price)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Body")
	fmt.Println(
		price.Department,
		price.Person,
		price.Amount,
		price.Risk,
	)

	manufactory := PriceManufactory{
		price.Amount,
		price.Risk,
	}
	person := PricePerson{
		price.Person,
		price.Amount,
		price.Risk,
	}
	government := PriceGovernment{
		price.Department,
		price.Person,
		price.Amount,
		price.Risk,
	}

	resultManufactory := make(chan float64)
	resultPerson := make(chan float64)
	resultGovernment := make(chan float64)

	go CalcPriceManufactory(manufactory, resultManufactory)
	go CalcPricePerson(person, resultPerson)
	go CalcPriceGovernment(government, resultGovernment)

	finalResultM := RecieveCalcPriceManufactory(resultManufactory)
	finalResultP := RecieveCalcPricePerson(resultPerson)
	finalResultG := RecieveCalcPriceGovernment(resultGovernment)

	return finalResultM, finalResultP, finalResultG

}
