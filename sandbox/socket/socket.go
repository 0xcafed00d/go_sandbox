package main

import (
	"net"
	"syscall"
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

func main() {

	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(sock)

	err = syscall.SetsockoptInt(sock, 0x0, syscall.IP_TTL, 6)
	if err != nil {
		panic(err)
	}

	addr, err := getAddress("www.ebay.com")
	if err != nil {
		panic(err)
	}

	err = syscall.Connect(sock, &syscall.SockaddrInet4{Port: 80, Addr: addr})
	if err != nil {
		panic(err)
	}

}
