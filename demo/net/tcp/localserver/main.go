package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("start")

	handler := func(ls *localServer, ln net.Listener) {
		c, err := ln.Accept()
		if err != nil {
			fmt.Printf("ln.Accept error(%+v)", err)
			return
		}
		defer c.Close()

		var b [1]byte
		n, err := c.Read(b[:])
		if n != 0 || err != io.EOF {
			fmt.Printf("got (%d, %v); want (0, io.EOF)", n, err)
			return
		}
		switch c := c.(type) {
		case *net.TCPConn:
			err = c.CloseWrite()
		case *net.UnixConn:
			err = c.CloseWrite()
		}
		if err != nil {
			fmt.Printf("c.CloseWrite error(%+v)", err)
			return
		}
		n, err = c.Write(b[:])
		if err == nil {
			fmt.Printf("got (%d, %v); want (any, error)", n, err)
			return
		}
	}

	network := "tcp"
	ls, err := newLocalServer(network)
	if err != nil {
		fmt.Printf("newLocalServer error(%+v)", err)
	}
	defer ls.teardown()
	if err := ls.buildup(handler); err != nil {
		fmt.Printf("ls.buildup error(%+v)", err)
	}

	c, err := net.Dial(ls.Listener.Addr().Network(), ls.Listener.Addr().String())
	if err != nil {
		fmt.Printf("net.Dial error(%+v)", err)
	}
	switch network {
	case "unix", "unixpacket":
		defer os.Remove(c.LocalAddr().String())
	}
	defer c.Close()

	switch c := c.(type) {
	case *net.TCPConn:
		err = c.CloseWrite()
	case *net.UnixConn:
		err = c.CloseWrite()
	}
	if err != nil {
		fmt.Printf("c.CloseWrite error(%+v)", err)
	}
	var b [1]byte
	n, err := c.Read(b[:])
	if n != 0 || err != io.EOF {
		fmt.Printf("got (%d, %v); want (0, io.EOF)", n, err)
	}

	fmt.Println("read", n, string(b[:]))
	n, err = c.Write(b[:])
	fmt.Println("write", n, string(b[:]))
	if err == nil {
		fmt.Printf("got (%d, %v); want (any, error)", n, err)
	}
}

type localServer struct {
	lnmu sync.RWMutex
	net.Listener
	done chan bool // signal that indicates server stopped
}

func (ls *localServer) buildup(handler func(*localServer, net.Listener)) error {
	go func() {
		handler(ls, ls.Listener)
		close(ls.done)
	}()
	return nil
}

func (ls *localServer) teardown() error {
	ls.lnmu.Lock()
	if ls.Listener != nil {
		network := ls.Listener.Addr().Network()
		address := ls.Listener.Addr().String()
		ls.Listener.Close()
		<-ls.done
		ls.Listener = nil
		switch network {
		case "unix", "unixpacket":
			os.Remove(address)
		}
	}
	ls.lnmu.Unlock()
	return nil
}

func newLocalServer(network string) (*localServer, error) {
	ln, err := newLocalListener(network)
	if err != nil {
		return nil, err
	}
	return &localServer{Listener: ln, done: make(chan bool)}, nil
}

func newLocalListener(network string) (net.Listener, error) {
	switch network {
	case "tcp":
		if supportsIPv4() {
			if ln, err := net.Listen("tcp4", "127.0.0.1:0"); err == nil {
				return ln, nil
			}
		}
		if supportsIPv6() {
			return net.Listen("tcp6", "[::1]:0")
		}
	case "tcp4":
		if supportsIPv4() {
			return net.Listen("tcp4", "127.0.0.1:0")
		}
	case "tcp6":
		if supportsIPv6() {
			return net.Listen("tcp6", "[::1]:0")
		}
	case "unix", "unixpacket":
		return net.Listen(network, testUnixAddr())
	}
	return nil, fmt.Errorf("%s is not supported", network)
}

// testUnixAddr uses ioutil.TempFile to get a name that is unique.
// It also uses /tmp directory in case it is prohibited to create UNIX
// sockets in TMPDIR.
func testUnixAddr() string {
	f, err := ioutil.TempFile("", "go-nettest")
	if err != nil {
		panic(err)
	}
	addr := f.Name()
	f.Close()
	os.Remove(addr)
	return addr
}

type ipStackCapabilities struct {
	sync.Once             // guards following
	ipv4Enabled           bool
	ipv6Enabled           bool
	ipv4MappedIPv6Enabled bool
}

var ipStackCaps ipStackCapabilities

// supportsIPv4 reports whether the platform supports IPv4 networking
// functionality.
func supportsIPv4() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv4Enabled
}

// supportsIPv6 reports whether the platform supports IPv6 networking
// functionality.
func supportsIPv6() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv6Enabled
}

// supportsIPv4map reports whether the platform supports mapping an
// IPv4 address inside an IPv6 address at transport layer
// protocols. See RFC 4291, RFC 4038 and RFC 3493.
func supportsIPv4map() bool {
	ipStackCaps.Once.Do(ipStackCaps.probe)
	return ipStackCaps.ipv4MappedIPv6Enabled
}

// CloseFunc is used to hook the close call.
var CloseFunc func(int) error = syscall.Close

// AcceptFunc is used to hook the accept call.
var AcceptFunc func(int) (int, syscall.Sockaddr, error) = syscall.Accept

// Probe probes IPv4, IPv6 and IPv4-mapped IPv6 communication
// capabilities which are controlled by the IPV6_V6ONLY socket option
// and kernel configuration.
//
// Should we try to use the IPv4 socket interface if we're only
// dealing with IPv4 sockets? As long as the host system understands
// IPv4-mapped IPv6, it's okay to pass IPv4-mapeed IPv6 addresses to
// the IPv6 interface. That simplifies our code and is most
// general. Unfortunately, we need to run on kernels built without
// IPv6 support too. So probe the kernel to figure it out.
func (p *ipStackCapabilities) probe() {
	s, err := sysSocket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	switch err {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
	case nil:
		CloseFunc(s)
		p.ipv4Enabled = true
	}
	var probes = []struct {
		laddr net.TCPAddr
		value int
	}{
		// IPv6 communication capability
		{laddr: net.TCPAddr{IP: net.ParseIP("::1")}, value: 1},
		// IPv4-mapped IPv6 address communication capability
		{laddr: net.TCPAddr{IP: net.IPv4(127, 0, 0, 1)}, value: 0},
	}
	switch runtime.GOOS {
	case "dragonfly", "openbsd":
		// The latest DragonFly BSD and OpenBSD kernels don't
		// support IPV6_V6ONLY=0. They always return an error
		// and we don't need to probe the capability.
		probes = probes[:1]
	}
	for i := range probes {
		s, err := sysSocket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		if err != nil {
			continue
		}
		defer CloseFunc(s)
		syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, probes[i].value)
		sa, err := probes[i].laddr.sockaddr(syscall.AF_INET6)
		if err != nil {
			continue
		}
		if err := syscall.Bind(s, sa); err != nil {
			continue
		}
		if i == 0 {
			p.ipv6Enabled = true
		} else {
			p.ipv4MappedIPv6Enabled = true
		}
	}
}

var (
	testHookDialChannel  = func() {} // for golang.org/issue/5349
	testHookCanceledDial = func() {} // for golang.org/issue/16523

	// Placeholders for socket system calls.
	socketFunc        func(int, int, int) (int, error)  = syscall.Socket
	connectFunc       func(int, syscall.Sockaddr) error = syscall.Connect
	listenFunc        func(int, int) error              = syscall.Listen
	getsockoptIntFunc func(int, int, int) (int, error)  = syscall.GetsockoptInt
)

// Wrapper around the socket system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
func sysSocket(family, sotype, proto int) (int, error) {
	s, err := socketFunc(family, sotype|syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC, proto)
	// On Linux the SOCK_NONBLOCK and SOCK_CLOEXEC flags were
	// introduced in 2.6.27 kernel and on FreeBSD both flags were
	// introduced in 10 kernel. If we get an EINVAL error on Linux
	// or EPROTONOSUPPORT error on FreeBSD, fall back to using
	// socket without them.
	switch err {
	case nil:
		return s, nil
	default:
		return -1, os.NewSyscallError("socket", err)
	case syscall.EPROTONOSUPPORT, syscall.EINVAL:
	}

	// See ../syscall/exec_unix.go for description of ForkLock.
	syscall.ForkLock.RLock()
	s, err = socketFunc(family, sotype, proto)
	if err == nil {
		syscall.CloseOnExec(s)
	}
	syscall.ForkLock.RUnlock()
	if err != nil {
		return -1, os.NewSyscallError("socket", err)
	}
	if err = syscall.SetNonblock(s, true); err != nil {
		CloseFunc(s)
		return -1, os.NewSyscallError("setnonblock", err)
	}
	return s, nil
}
