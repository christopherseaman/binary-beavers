package main

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
)

// Parameters for g, g^a, and p
var (
	g     = int64(5)
	g_a   = big.NewInt(27)
	p     = big.NewInt(443)
	bound = int64(31)
	// g_a  = big.NewInt(244241057144443665472449725715508406620552440771362355600491)
	// p   = big.NewInt(3217014639198601771090467299986349868436393574029172456674199)
	// bound = int64(3569817792505)9223372036854775807
)

// FactorBase returns a slice of prime numbers up to a given bound
// FactorBase returns a slice of prime numbers up to a given bound
func FactorBase() []int64 {
	var factorBase []int64

	// Iterate from 2 to the bound and check if each number is prime
	for i := int64(2); i <= bound; i++ {
		// Create a big.Int from the current number to use the ProbablyPrime method
		x := big.NewInt(i)

		// Check if the number is probably prime
		if x.ProbablyPrime(1) {
			factorBase = append(factorBase, i)
		}
	}

	return factorBase
}

type Relation map[int64]int

// generateRelation generates a relation using the quadratic sieve
func generateRelation(g int64, g_a, p *big.Int, factorBase []int64) Relation {
	// IDEA: min k st g^k > to bound and (g^k)^2 - h^2 mod p to generate relations
	// Placeholder for actual relation generation logic
	// This function should:
	// 1. Choose random h < bound and compute (g+h)(g-h) mod p
	// 2. Attempt to factor (g^j+h)(g^k-h) using the factor base sv1
	// 3. Factor (g^k+h) and (g^k-h) over the factor base as sv2 and sv3
	// 4. Save the relation sv1 - sv2 - sv3 in a map
	// Note: -1 == p-2 because exponents are mod (p-1)

	// Placeholder return value
	return Relation{2: 3, 3: 1} // Example with factored 2^3 * 3^1
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

func main() {
	// FIXME: Implement generateRelation using the quadratic sieve
	relation := generateRelation(g, g_a, p, FactorBase())
	if err := sendRelation(relation); err != nil {
		// Handle error
		panic(err)
	}
}
