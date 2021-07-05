package main

import (
	"fmt"
	"github.com/jerry/keyboard"
	"log"
)

/**
	常量声明格式: const name type = value 或者可以省略类型声明: const name = value,通过分配的值推断常量类型
 */
const validateGrade float64 = 60

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string

	if grade >= validateGrade {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Printf("A grade of %.2f is %s\n", grade, status)
}
