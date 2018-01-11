package main

import (
	"fmt"
	"net"
	"sort"
)

func listInterfaces() {
	ifs, _ := net.Interfaces()
	sort.SliceStable(ifs, func(i, j int) bool { return ifs[i].Index < ifs[j].Index })
	if len(ifs) == 0 {
		println("\nThere is no network interfaces")
	} else {
		println("\n Index Name                                Address")
		println(" ----- ----------------------------------- ----------------------------")
		for _, v := range ifs {
			fmt.Printf(" %-5d %-35s ", v.Index, v.Name)
			addrs, _ := v.Addrs()
			for _, a := range addrs {
				print(a.String())
			}
			println("")
		}
	}
}
