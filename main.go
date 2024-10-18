package main

import (
	"fmt"
	"golang-crash-course/src/api"
	"golang-crash-course/src/functions"
	"golang-crash-course/src/hello"
	"golang-crash-course/src/loops"
	"golang-crash-course/src/slices"
	"golang-crash-course/src/variables"
)

func main() {
	fmt.Println("Hello World")

	hello.Demonstrate()
	variables.Demonstrate()
	slices.Demonstrate()
	loops.Demonstrate()
	functions.Demonstrate()
	api.Demonstrate()
}