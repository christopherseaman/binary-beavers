package main

import (
    "math/big"
    "net/http"
    "bytes"
    "encoding/json"
)

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
