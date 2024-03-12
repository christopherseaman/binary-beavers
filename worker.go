package main

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/fxtlabs/primes"
)

// Parameters for g, g^a, and p
var (
	g     = big.NewInt(5)
	g_a   = big.NewInt(27)
	p     = big.NewInt(443)
	bound = big.NewInt(31)
	// g_a  = big.NewInt(244241057144443665472449725715508406620552440771362355600491)
	// p   = big.NewInt(3217014639198601771090467299986349868436393574029172456674199)
	// bound = big.NewInt(3569817792505)
)

// FactorBase returns a slice of prime numbers up to a given bound
func FactorBase() []*big.Int {
	return primes.Sieve(bound)
}

// Relation represents a relation to be sent to the collector
type Relation struct {
	Data string `json:"data"` // Adjust according to expected format
}

func main() {
	// Example: Generate and send a relation
	relation := Relation{Data: "example_relation_data"}
	if err := sendRelation(relation); err != nil {
		// Handle error
		panic(err)
	}
}

// sendRelation sends a relation to the collector
func sendRelation(relation Relation) error {
	jsonData, err := json.Marshal(relation)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://r.badmath.org", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle response, e.g., check if status code is 202 or 418
	// ...

	return nil
}
