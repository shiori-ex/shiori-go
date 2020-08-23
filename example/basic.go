package main

import (
	"fmt"

	"github.com/shiori-ex/shiori-go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c := shiori.NewClient("shiori_gateway_token", "https://shiori.zekro.de/api")

	newLink, err := c.CreateLink(&shiori.Link{
		Url:         "https://zekro.de",
		Description: "Web page from a cool developer!",
		Tags:        []string{"dev", "code", "profile"},
	})
	check(err)
	fmt.Printf("New link: %+v\n", newLink)

	links, err := c.Links(10, 0)
	check(err)
	for _, l := range links {
		fmt.Printf("Link in list: %+v\n", l)
	}

	newLink.Description = "Delete me!"
	updatedLink, err := c.UpdateLink(newLink)
	check(err)
	fmt.Printf("Updated link: %+v\n", updatedLink)

	fmt.Println(c.RemoveLink(updatedLink.Id))
}
