package main

import "fmt"

func main() {
	updateSlice()
	copySlice()
	appendSlice()
}

func updateSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println("original s1:", s1)

	// Change element in slice.
	s1[1] = 200
	fmt.Println("after changing s1[1] to 100:", s1)

	s2 := s1
	fmt.Println("s2 is a copy of s1:", s2)

	s2[0] = 6
	fmt.Println("after changing s2[0] to 6, s1 is also changed:", s1)

	// Reassign slice.
	s1 = []int{4, 5, 8}
	fmt.Println("reassign s1:", s1)

	s2[0] = 9
	fmt.Println("after changing s2[0], s1 no longer changes:", s1)

	fmt.Println("\nConclusion: Reassigning a slice (line 22) points to a new underlying array.")
	fmt.Println("However, changing an element in the slice (line 12) keeps the same array.")
}

func appendSlice(){
// TODO
}

func copySlice(){
// TODO
}