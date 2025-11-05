/*
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
*/
package main

import "fmt"

// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// exemple d'une tirelire, dans une tirelire on aajoute de l'argent, et on se qu'a chaque ajout la tirelire sait qu'on a ajoute et on peut dire qu'elle compte...
func piggyBank() func(int) int {
	money := 0

	return func(amount int) int {
		money += amount
		return money
	}
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func counter(step int) func(int) int {
	
	return func(count int) int {
	
		return count * step
	}
}

// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
func main() {
	/*nextInt := intSeq()
	save := piggyBank()
	double := multiplier(2)
	triple := multiplier(3)*/
	c := counter(3)

	fmt.Println(c(3))
	fmt.Println(c(4))
	fmt.Println(c(5))

	/*fmt.Println(save(10))
	fmt.Println(save(20))
	fmt.Println(save(30))



	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(double(4))
	fmt.Println(triple(4))*/
}
