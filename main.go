package main

import (
	facebook "Packages/Facebook"
	interfaces "Packages/Interfaces"
	twitter "Packages/Twitter"
	"fmt"
	"os"
)

// func main() {
// 	fb := new(facebook.Facebook)
// 	twtr := new(twitter.Twitter)
// 	lnkdin := new(linkedln.Linkedln)

// 	ScrollFeeds(twtr, fb, lnkdin)
// }

func main() {
	fb := new(facebook.Facebook)
	err := export(fb, "fbdata.txt")
	twtr := new(twitter.Twitter)
	err = export(twtr, "twtrdata.txt")

	if err != nil {
		panic(err)
	}
}

func ScrollFeeds(platforms ...interfaces.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("************************************")
	}
}

func export(u interfaces.SocialMedia, filename string) error {
	f, err := os.OpenFile("fbdata.txt", os.O_CREATE | os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd+"\n"))
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes\n", n)
	}
	return nil
}