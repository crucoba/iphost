package main

import (
	"fmt"
	"net"
	"strconv"
)

func getAddrByIdx(idx string) (net.Addr, error) {
	ix, err1 := strconv.Atoi(idx)
	if err1 != nil {
		return nil, fmt.Errorf("Invalid iterface index %s: %s", idx, err1)
	}
	iface, err2 := net.InterfaceByIndex(ix)
	if err2 != nil {
		return nil, fmt.Errorf("Unknown interface %d: %s", ix, err2)
	}
	addr, err3 := iface.Addrs()
	if err3 != nil {
		return nil, fmt.Errorf("Interface %d has no address: %s", ix, err3)
	}
	return addr[0], nil
}
