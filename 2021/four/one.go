package main

func main() {

	numbers := []string{"1", "2", "3", "4"}
	c := newCage(numbers)

	// This code will close the channel, alow for looping...
	//	close(c)
	//	for ball := range c {
	//		fmt.Println(ball)
	//	}

	// If left open we can just read entries..

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	// Is a slice better? here, Kind like the stack queue

}

func newCage(numbers []string) chan string {
	buffer := len(numbers)
	c := make(chan string, buffer)
	for _, num := range numbers {
		c <- num
	}

	return c
}
