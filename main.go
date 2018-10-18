package main

import (
	"fmt"

	"github.com/linkedin/goavro"
)

func main() {
	// Schema bellow the currency logical type should be treated as a string
	// https://avro.apache.org/docs/1.8.2/spec.html#Logical+Types
	_, err := goavro.NewCodec(`
	{
		"type": "record",
		"name": "TestRecord",
		"fields": [
			{
				"name": "currency",
				"type": {
						"type": "string",
						"avro.java.string": "String",
						"logicalType": "x-currency-iso4217"
				}
			}
		]
	}`)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
	}
}
