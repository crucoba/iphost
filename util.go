package main

import (
	"fmt"
	"net"
	"strconv"
)

func getAddrByIdx(idx string) (net.Addr, error) {
	ix, err1 := strconv.Atoi(idx)
	if err1 != nil {
		return nil, fmt.Errorf("Número de interfaz incorrecto: %s", err1)
	}
	iface, err2 := net.InterfaceByIndex(ix)
	if err2 != nil {
		return nil, fmt.Errorf("No existe el interfaz: %s", err2)
	}
	addr, err3 := iface.Addrs()
	if err3 != nil {
		return nil, fmt.Errorf("No hay direcciones: %s", err3)
	}
	return addr[0], nil
}
