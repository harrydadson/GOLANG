package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go keyword only infront of func calls
	}
	for l := range c { 
		go func(link1 string) { // function literal equivalent to Lambda in python
			time.Sleep(5 * time.Second)
			checkLink(link1, c)
		}(l)
	}
}

/* or
for i := 0; i < len(links); i++ {
	go(fmt.checklink(<- c) 

for l := range c { 
		go checkLink(l, c)
*/

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err!=nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "yup it's up"

}

/*

Sending Date with Channels

channel <- 5 send the value "5" into this channel

myNumber <- channel wait for a value to be sent into the channel. when
we get one, assign the value to "myNumber"

fmt.Println(<- channel) wait for a value to be sent into the channel
when we get one, log it out immediately

*/