package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

func getAddress(host string) ([4]byte, error) {
	var addr [4]byte

	addresses, err := net.LookupHost(host)
	if err != nil {
		return addr, err
	}

	ip, err := net.ResolveIPAddr("ip", addresses[0])
	if err != nil {
		return addr, err
	}

	copy(addr[:], ip.IP.To4())
	return addr, nil
}

type SocketState int

const (
	SocketConnected SocketState = iota
	SocketTimedOut
	SocketPortClosed
	SocketError
)

func waitWithTimeout(socket int, timeout time.Duration) {
	start := time.Now()

	rfdset := &syscall.FdSet{}
	wfdset := &syscall.FdSet{}
	efdset := &syscall.FdSet{}

	FD_ZERO(rfdset)
	FD_ZERO(wfdset)
	FD_ZERO(efdset)

	FD_SET(rfdset, socket)
	FD_SET(wfdset, socket)
	FD_SET(efdset, socket)

	timeval := syscall.NsecToTimeval(int64(timeout))

	n, err := syscall.Select(socket+1, rfdset, wfdset, efdset, &timeval)

	fmt.Println(n, err, FD_ISSET(rfdset, socket), FD_ISSET(wfdset, socket), FD_ISSET(efdset, socket), time.Since(start))
}

func connect(host string, port int, timeout time.Duration) error {
	fmt.Println("\nConnecting to: ", host, "......")

	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		return err
	}
	defer syscall.Close(sock)

	err = syscall.SetNonblock(sock, true)
	if err != nil {
		return err
	}

	addr, err := getAddress(host)
	if err != nil {
		return nil
	}

	// ignore error from connect in non-blocking mode. as it will always return a
	// in progress error
	_ = syscall.Connect(sock, &syscall.SockaddrInet4{Port: port, Addr: addr})

	waitWithTimeout(sock, 500*time.Millisecond)

	return nil
}

func main() {
	err := connect("www.google.com", 80, 500*time.Millisecond)
	err = connect("www.google.com", 89, 500*time.Millisecond)

	fmt.Println(err)
}

func FD_SET(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func FD_ISSET(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func FD_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}
