package main

import (
	"fmt"
	"net"
	"sort"
)

func listInterfaces() {
	ifs, _ := net.Interfaces()
	sort.SliceStable(ifs, func(i, j int) bool { return ifs[i].Index < ifs[j].Index })
	for _, v := range ifs {
		fmt.Printf("%3d %-35s ( ", v.Index, v.Name)
		addrs, _ := v.Addrs()
		for _, a := range addrs {
			print(a.String())
		}
		println(" )")
	}
}
