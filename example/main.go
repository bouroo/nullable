package main

import (
	"encoding/json"
	"fmt"
	"time"
	
	"nullable"
)

type Person struct {
	Name     string
	Birthday nullable.Time
	Age      nullable.Value[int]
}

func main() {
	p1 := Person{
		Name:     "John Doe",
		Birthday: nullable.TimeOf(time.Now()),
		Age:      nullable.ValueOf(30),
	}

	p2 := Person{
		Name:     "Jane Smith",
		Birthday: nullable.Time{},
		Age:      nullable.Value[int]{},
	}

	p3 := Person{
		Name:     "Bob Johnson",
		Birthday: nullable.TimeOf(time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)),
		Age:      nullable.ValueOf(40),
	}

	people := []Person{p1, p2, p3}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
