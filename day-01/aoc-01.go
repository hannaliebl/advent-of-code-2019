package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("file", "", "Fuel input file")

func main() {
	flag.Parse()

	f := *inputFile
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += calcFuel(mass)
	}

	fmt.Println(total)
}

func calcFuel(mass int) int {
	return (mass / 3) - 2
}
