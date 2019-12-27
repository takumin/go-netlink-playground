package main

import (
	"log"

	"github.com/vishvananda/netlink"
)

func main() {
	la := netlink.NewLinkAttrs()
	la.Name = "test"
	br := &netlink.Bridge{LinkAttrs: la}
	if err := netlink.LinkAdd(br); err != nil {
		log.Fatal(la.Name, err)
	}
}
