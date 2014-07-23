package main

import (
	"fmt"
	"github.com/fasterness/lxc-config"
)

func main() {
	c := lxcconfig.New()
	str := c.String()
	fmt.Println(str)
}
