package main

import (
	"bufio"
	"fmt"
	"os"
)

var myVar map[int]int = map[int]int{1: 2, 2: 6, 3: 3, 4: 5, 5: 1, 6: 4}

func main(){
	// scan user input
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	// remove \n symbol
	line = line[:len(line) -1]
	// code user input
	coded := code(line, myVar)
	// decode user input
	decoded := decode(coded, myVar)
	fmt.Printf("RESULT\n------\ninput:  '%s'\ncode:   '%s'\ndecode: '%s'", line, coded, decoded)
}

//Code string with input permutation group
func code(line string, mapa map[int]int) string{
	// tirm last line's spaces
	end := TrimLastSpaces(line)
	// init slice of runes for better optimization
	lineRunes := []rune(line[:end+1])
	// add spaces for len(line) % 6 == 0
	for len(lineRunes) % 6 != 0{
		lineRunes = append(lineRunes, ' ')
	}
	// slice of runes for optimization
	coded := make([]rune, 0)
	for i := 0; i < end; i+=len(mapa){
		// get line's segment with mapa's len
		segment := lineRunes[i:i+len(mapa)]
		// slice contains coded segment
		codedPart := make([]rune, len(mapa))
		for i, r := range segment{
			// set runes according to the table
			codedPart[mapa[i+1] - 1] = r
		}
		coded = append(coded, codedPart...)
	}
	// trim last spaces
	res := string(coded)
	end = TrimLastSpaces(res)
	return res[:end + 1]
}

//Decode string with input permutation group
func decode(line string, mapa map[int]int) string{
	// reverse mapa
	mapa = reverseMap(mapa)
	// and invoke code function with reverse mapa
	return code(line, mapa)
}

//Trim last spaces
func TrimLastSpaces(s string) int{
	end := len(s) - 1
	for s[end] == ' '{
		end--
	}
	return end
}

//Reverse map[int]int
func reverseMap(mapa map[int]int) map[int]int{
	res := make(map[int]int)
	for k, v := range mapa{
		res[v] = k
	}
	return res
}