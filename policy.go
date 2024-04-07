package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Policy struct {
	PolicyDocument struct {
		Statement []struct {
			Resource interface{} `json:"Resource"`
		} `json:"Statement"`
	} `json:"PolicyDocument"`
}

func verifyPolicy(jsonData []byte) bool {
	var policy *Policy
	err := json.Unmarshal(jsonData, &policy)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return false
	}

	for _, statement := range policy.PolicyDocument.Statement {
		//Check if the Resource is a string or a slice of strings
		switch resource := statement.Resource.(type) {

		case string:
			//When the Resource is a string
			if strings.Contains(resource, "*") {
				return false
			}

		case []interface{}:
			// When the Resource is an array of interface{}
			for _, res := range resource {
				if rStr, ok := res.(string); ok {
					// Check if the element is a string
					if strings.Contains(rStr, "*") {
						return false
					}
				} else {
					// Handle if the element is not a string
					fmt.Println("Error: Unexpected type inside array")
					return false
				}
			}
		default:
			fmt.Println("Error: Unexpected type for Resource")
			return true
		}
	}

	return true
}
