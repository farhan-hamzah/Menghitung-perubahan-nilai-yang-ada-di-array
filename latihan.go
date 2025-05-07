package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, count int
	count = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}

	for i = 0; i < n; i++{
		if A[i] != A[i+1]{
			count++
		}
	}
	fmt.Print(count)
}