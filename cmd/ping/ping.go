package main

import (
	"fmt"
	"net"
	"time"
)

func PingHost(host string) string{
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "80"), 2*time.Second)
	if err != nil{
		return fmt.Sprintf("%s: unreachable", host)
	}
	defer conn.Close()

	duration := time.Since(start)
	return fmt.Sprintf("%s: %s", host, duration)
}
