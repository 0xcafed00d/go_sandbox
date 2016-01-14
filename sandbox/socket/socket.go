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

func connect(host string, port, ttl int, timeout time.Duration) error {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		return err
	}
	defer syscall.Close(sock)

	err = syscall.SetsockoptInt(sock, 0x0, syscall.IP_TTL, ttl)
	if err != nil {
		return err
	}

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
	_ = syscall.Connect(sock, &syscall.SockaddrInet4{Port: 80, Addr: addr})

	fdset := &syscall.FdSet{}
	timeoutVal := &syscall.Timeval{}
	timeoutVal.Sec = int64(timeout / time.Second)
	timeoutVal.Usec = int64(timeout-time.Duration(timeoutVal.Sec)*time.Second) / 1000

	fmt.Println(timeoutVal)

	FD_ZERO(fdset)
	FD_SET(fdset, sock)

	start := time.Now()
	x, err := syscall.Select(sock+1, nil, fdset, nil, timeoutVal)
	elapsed := time.Since(start)

	fmt.Println(x, elapsed)
	if err != nil {
		return err
	}

	if FD_ISSET(fdset, sock) {
		fmt.Println("conencted")
	} else {
		fmt.Println("timedout")
	}

	return nil
}

func main() {
	fmt.Println(connect("www.google.com", 80, 80, 10*time.Millisecond))
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
