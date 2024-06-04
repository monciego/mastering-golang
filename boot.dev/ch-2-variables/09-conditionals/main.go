package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}


	/* An if conditional can have an "initial" statement.  */
	/* 
	syntax: if INITIAL_STATEMENT; CONDITION {
		...code here
	}
	*/
	if age := 10;  age < 18 {
		fmt.Println("Grow up!")
	}

}
