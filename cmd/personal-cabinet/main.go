package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Shalom!")           ///TEST
	fmt.Println("kak dela?")         ///TEST
	fmt.Println("ya toje normal'no") /*TTTTTTTTTTTTTTTTTTTTTTTTTTTEEEEEEEEEEEEEEEEEEEEEEEEEEE
	SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSTTTTTTTTTTTTTTTT*/

	var (
		seew          string
		a, b, c, d, e int

		bytes []byte
	)

	seew = "Shalomssssss"

	/*
		Formatting ( import "fmt"; use fmt.Printf(); use {varName} = fmt.Sprintf())

		%v - value interface{}(any)
		%s - strings ""
		%d - decimal(0-9)
		%x - bytes
		%t - bool
		%T - type

		string to bytes([]byte)
	*/
	bytes = []byte(seew + "_e_" + "_e")
	fmt.Printf("as struct:\t%v\n", struct {
		name string
		age  int
	}{name: "name_value"})
	fmt.Printf("as struct:\t%+v\n", struct {
		name string
		age  int
	}{name: "name"})
	fmt.Printf("as string: \t%s\n", bytes)
	fmt.Printf("as decimals: \t%d\n", 8)
	fmt.Printf("as bytes:\t%x\n", bytes)
	fmt.Printf("as bool:\t%t\n", true)
	fmt.Printf("as type:\t%T\n", 8)

	scanName(os.Stdin)
	scanAndPrintName(os.Stdin)

	a = 1
	b = 2
	c = 3
	d = 4
	e = a + b*c - d
	fmt.Println(seew, e)
}

// получить имя и напечатать его в консоль
func scanAndPrintName(reader io.Reader) { printName(scanName(reader)) }

// получить имя и вернуть результат работы функции (return)
func scanName(reader io.Reader) string { return bufio.NewScanner(reader).Text() }

// печатаем введённое имя
func printName(name string) { fmt.Printf("Your name:  %s", name) }
