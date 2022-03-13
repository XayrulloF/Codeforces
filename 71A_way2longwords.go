package main
import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	var sliceWords = make([]string, num)

	for i:=0; i<num; i++ {
		var word string
		fmt.Scan(&word)
		if len(word) > 10 {
			wordHelper := string(word[0]) + strconv.Itoa(len(word) - 2) + string(word[len(word)-1])
			sliceWords[i] = wordHelper
		} else {
			sliceWords[i] = word
		}
	}
	for _, j := range sliceWords {
		fmt.Println(j)
	}
}