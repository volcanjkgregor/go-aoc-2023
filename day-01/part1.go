package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//var firstNumber, secondNumber int
	input, err := os.Open("input1.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var sum rune
	var leftnumstr, rightnumstr string
	for scanner.Scan() {
	subloop:
		for _, elem := range scanner.Text() {
			if elem >= '0' && elem <= '9' {
				leftnumstr = string(elem)
				break subloop
			}
		}
	secondsubloop:
		for i := len(scanner.Text()) - 1; i >= 0; i-- {
			elem := scanner.Text()[i]
			if elem >= '0' && elem <= '9' {
				rightnumstr = string(elem)
				break secondsubloop
			}
		}
		lineresult, err := strconv.Atoi(leftnumstr + rightnumstr)
		if err != nil {
			fmt.Println("Strconv Error:", err)
			continue
		}

		sum += rune(lineresult)

	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
