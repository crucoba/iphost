package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"regexp"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func setAlias(idx string, alias string) {
	addr, err1 := getAddrByIdx(idx)
	if err1 != nil {
		println(err1.Error())
		os.Exit(1)
	}
	ip, _, _ := net.ParseCIDR(addr.String())

	println("Setting", alias, "as alias for", ip.String())

	hostsPath := os.ExpandEnv("$SYSTEMROOT/System32/drivers/etc/hosts")
	hosts, err4 := ioutil.ReadFile(hostsPath)
	if err4 != nil {
		println("Unable to read hosts file: ", err4)
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
