package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	args := os.Args
	switch args[1] {
	case "list":
		listInterfaces()
	case "set":
		setAlias(args[2], args[3])
	}
}

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

func setAlias(idx string, alias string) {
	addr, err1 := getAddr(idx)
	if err1 != nil {
		println(err1.Error())
		os.Exit(1)
	}
	ip, _, _ := net.ParseCIDR(addr.String())

	println("Setting", alias, "as alias for", ip.String())

	hostsPath := os.ExpandEnv("$SYSTEMROOT/System32/drivers/etc/hosts")
	hosts, err4 := ioutil.ReadFile(hostsPath)
	if err4 != nil {
		println("No se puede leer hosts: ", err4)
		os.Exit(3)
	}

	hosts, _, _ = transform.Bytes(charmap.Windows1252.NewDecoder(), hosts)
	re := regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+\\s+" + regexp.QuoteMeta(alias))
	idxs := re.FindIndex(hosts)
	if idxs == nil {
		println("New alias")
		var buf bytes.Buffer
		buf.Write(hosts)
		buf.WriteString("\r\n")
		buf.WriteString(ip.String())
		buf.WriteString("\t\t")
		buf.WriteString(alias)
		buf.WriteString("\r\n")
		hosts = buf.Bytes()
	} else {
		println("Found alias in", idxs[0], ",", idxs[1])
		var buf bytes.Buffer
		buf.Write(hosts[:idxs[0]])
		buf.WriteString(ip.String())
		buf.WriteString("\t\t")
		buf.WriteString(alias)
		buf.Write(hosts[idxs[1]:])
		hosts = buf.Bytes()
	}

	//print(string(hosts))

	hostsOut, _, _ := transform.Bytes(charmap.Windows1252.NewEncoder(), hosts)
	ioutil.WriteFile(hostsPath, hostsOut, 0)
}

func getAddr(idx string) (net.Addr, error) {
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
