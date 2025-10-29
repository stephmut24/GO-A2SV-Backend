
// For is Go's only looping construct. here are some basic types of for loops.

package main

import "fmt"

func main() {

	//The most basic type, with a single condition.
	i :=1
	for i <= 4 { 
		fmt.Println(i)
		i = i + 1
	}
	
	// A classic inital/condition/after for loop.
	for j := 0; j < 3; j++ {
		//fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this N times" iteration
	// is to use a standard for loop that counts from 0 to N-1.
	for i := 0; i < 3; i++ {
		//fmt.Println("range", i)
	}

	// For without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop

	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}