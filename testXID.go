package main

import (
	"fmt"

	"github.com/rs/xid"
)

func XID() {
	fmt.Println(xid.New().String())
	fmt.Println(xid.New().Machine())
	fmt.Println(xid.New().Pid())
	fmt.Println(xid.New().Time())
	fmt.Println(xid.New().Counter())

}
