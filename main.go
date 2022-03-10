package main
import (
	"fmt"
)
func main() {
	var num int
	fmt.Scanf("%d", &num)
	if num % 2 == 0 {
		yesfunc(num)
	} else {
		fmt.Println("NO")
	}
}
func yesfunc (num int) {
	if num == 2 {
		fmt.Println("NO")
		return
	}
	for i := 2; i <= (num/2);  i+=2 {
		if (num - i) % 2 == 0 {
			fmt.Println("YES")
			break
		}
	}
}