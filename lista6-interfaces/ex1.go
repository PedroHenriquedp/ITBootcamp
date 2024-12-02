package main

import "fmt"

type Student struct {
    Name  string
    Surname string
    DNI   int
    Date  string
}

func (s Student) PrintDetails() {
    fmt.Printf("Nome: %s\n", s.Name)
    fmt.Printf("Sobrenome: %s\n", s.Surname)
    fmt.Printf("ID: %d\n", s.DNI)
    fmt.Printf("Data: %s\n", s.Date)
}

func main() {
    student := Student{
        Name:    "João",
        Surname: "Silva",
        DNI:     123456,
        Date:    "2023-01-15",
    }

    student.PrintDetails()
}

