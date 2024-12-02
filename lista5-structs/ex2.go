package main

import (
    "fmt"
)

// Estrutura Person
type Person struct {
    ID          int
    Name        string
    DateOfBirth string
}

// Estrutura Employee com composição de Person
type Employee struct {
    ID       int
    Position string
    Person
}

func (e Employee) PrintEmployee() {
    fmt.Printf("ID: %d\n", e.ID)
    fmt.Printf("Name: %s\n", e.Name)
    fmt.Printf("Date of Birth: %s\n", e.DateOfBirth)
    fmt.Printf("Position: %s\n", e.Position)
}

func main() {
    person := Person{
        ID:          1,
        Name:        "Alice Smith",
        DateOfBirth: "1985-01-15",
    }

    employee := Employee{
        ID:       1001,
        Position: "Software Engineer",
        Person:   person,
    }

    employee.PrintEmployee()
}

