package main

import (
	"fmt"
	"net"
	"time"
)

type Domain struct {
	Url  string
	Port string
}

func (d *Domain) Check() string {
	var status string
	address := d.Url + ":" + d.Port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		status = fmt.Sprintf("[DOWN] \t %v is unreachable, \nError: %v", d.Url, err)
	} else {
		status = fmt.Sprintf("[UP] \t %v is reachable, \nFrom: %v\nTo: %v", d.Url, conn.LocalAddr().String(), conn.RemoteAddr().String())
	}

	return status
}
