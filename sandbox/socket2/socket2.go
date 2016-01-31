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

func (s SocketState) String() string {
	switch s {
	case SocketConnected:
		return "SocketConnected"
	case SocketTimedOut:
		return "SocketTimedOut"
	case SocketPortClosed:
		return "SocketPortClosed"
	case SocketError:
		return "SocketError"
	}
	return "SocketInvlaidState"
}

func waitWithTimeoutPrint(socket int, timeout time.Duration) {
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
	errcode, err := syscall.GetsockoptInt(socket, syscall.SOL_SOCKET, syscall.SO_ERROR)
	fmt.Println(n, err, FD_ISSET(rfdset, socket), FD_ISSET(wfdset, socket), FD_ISSET(efdset, socket), time.Since(start), errcode)
}

func waitWithTimeout(socket int, timeout time.Duration) (state SocketState, err error) {
	wfdset := &syscall.FdSet{}

	FD_ZERO(wfdset)
	FD_SET(wfdset, socket)

	timeval := syscall.NsecToTimeval(int64(timeout))

	n, err := syscall.Select(socket+1, nil, wfdset, nil, &timeval)
	if err != nil {
		state = SocketError
		return
	}
	errcode, err := syscall.GetsockoptInt(socket, syscall.SOL_SOCKET, syscall.SO_ERROR)
	if err != nil {
		state = SocketError
		return
	}

	if errcode == int(syscall.ECONNREFUSED) {
		state = SocketPortClosed
		return
	}

	if errcode != 0 {
		state = SocketError
		err = fmt.Errorf("Connect Error: %v", errcode)
		return
	}

	if n == 0 {
		state = SocketTimedOut
	} else {
		state = SocketConnected
	}
	return
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

	state, err := waitWithTimeout(sock, 500*time.Millisecond)

	fmt.Println(state, err)
	return nil
}

func main() {
	err := connect("robotmonkey.duckdns.org", 22, 500*time.Millisecond)
	err = connect("robotmonkey.duckdns.org", 80, 500*time.Millisecond)
	err = connect("www.google.com", 80, 500*time.Millisecond)
	err = connect("www.google.com", 81, 500*time.Millisecond)

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
