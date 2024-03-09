package main

import "fmt"

func main() {
	number := 10

    fmt.Println("Golang program to check that the number is even or odd using the modulus Relational operator.")

    // using the % operator and using the if else block accordingly
    if number%2 == 0 {
        fmt.Printf("The number %d is Even.\n", number)
    } else {
        fmt.Printf("The number %d is Odd.\n", number)
    }
}