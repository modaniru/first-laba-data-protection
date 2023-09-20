package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapa map[int]int = map[int]int{1: 2, 2: 6, 3: 3, 4:5, 5:1, 6:4}
var decodeMapa map[int]int = map[int]int{1: 5, 2: 1, 3: 3, 4: 6, 5: 4, 6: 2}

func main(){
	// scan user input
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	// remove \n symbol
	line = line[:len(line) -1]
	// code user input
	coded := code(line)
	// decode user input
	decoded := decode(coded)
	fmt.Printf("RESULT\n------\ninput: '%s'\ncode: '%s'\ndecode: '%s'", line, coded, decoded)
}

func code(line string) string{
	i := 0
	coded := make([]rune, 0)
	end := len(line) - 1
	for line[end] == ' '{
		end--
	}
	for i = 0; i < end - len(mapa) - 1; i+=len(mapa){
		segment := line[i:i+len(mapa)]
		codedPart := make([]rune, len(mapa))
		for i, r := range segment{
			codedPart[mapa[i+1] - 1] = r
		}
		coded = append(coded, codedPart...)
	}
	if i <= end{
		segment := line[i:end + 1]
		codedPart := make([]rune, len(mapa))
		for i := range codedPart{
			codedPart[i] = ' '
		}
		for i, r := range segment{
			codedPart[mapa[i+1] - 1] = r
		}
		coded = append(coded, codedPart...)
	}
	return string(coded)
}


func decode(line string) string{
	i := 0
	coded := make([]rune, 0)
	end := len(line) - 1
	for i = 0; i < end - len(mapa) - 1; i+=len(mapa){
		segment := line[i:i+6]
		codedPart := make([]rune, len(decodeMapa))
		for i, r := range segment{
			codedPart[decodeMapa[i+1] - 1] = r
		}
		coded = append(coded, codedPart...)
	}
	if i <= end{
		segment := line[i:end + 1]
		codedPart := make([]rune, len(decodeMapa))
		for i := range codedPart{
			codedPart[i] = ' '
		}
		for i, r := range segment{
			codedPart[decodeMapa[i+1] - 1] = r
		}
		coded = append(coded, codedPart...)
	}
	res := strings.TrimSpace(string(coded))
	return res
}