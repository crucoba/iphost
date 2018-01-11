package main

import "net"

func listInterfaces() {
	ifs, _ := net.Interfaces()
	for _, v := range ifs {
		print(v.Index, " ", v.Name, " ( ")
		addrs, _ := v.Addrs()
		for _, a := range addrs {
			print(a.String())
		}
		println(" )")
	}
}
