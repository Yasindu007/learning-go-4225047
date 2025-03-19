package main

import (
	"fmt"
	"time"
)

func main() {

	t:= time.Date(2009, time.November, 10,23, 0, 0, 0, time.UTC)
	fmt.Printf("Go lauched at %s\n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n)
	fmt.Printf("This object type is time %T\n", n)

	fmt.Printf(n.Format(time.ANSIC) + "\n")

	tomorrow := n.AddDate(0,0,1)



}
