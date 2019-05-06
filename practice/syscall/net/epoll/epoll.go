package epoll

import (
	"log"
	"syscall"
)

const (
	EPOLLIN       = 0x1
	EPOLLOUT      = 0x4
	EPOLLERR      = 0x8
	EPOLLHUP      = 0x10
	EPOLLRDHUP    = 0x2000
	EPOLLET       = 0x80000000
	EPOLL_CLOEXEC = 0x80000
	EPOLL_CTL_ADD = 0x1
	EPOLL_CTL_DEL = 0x2
	EPOLL_CTL_MOD = 0x3

	MaxEpollEvents = 32
)

type Poller struct {
	FD     int
	Event  syscall.EpollEvent
	Events [MaxEpollEvents]syscall.EpollEvent
}

func New() *Poller {
	fd, err := syscall.EpollCreate1(0)
	if err != nil {
		log.Fatalf("syscall.EpollCreate1 error(%v)", err)
	}
	return &Poller{
		FD: fd,
	}
}

func (p *Poller) Add(fd int) {
	p.Event.Fd = int32(fd)
	p.Event.Events = EPOLLIN | EPOLLOUT | EPOLLRDHUP | EPOLLET
	if err := syscall.EpollCtl(p.FD, EPOLL_CTL_ADD, fd, &p.Event); err != nil {
		log.Fatalf("Add syscall.EpollCtl error(%v)", err)
	}
}

func (p *Poller) Modify(fd int) {
	p.Event.Fd = int32(fd)
	p.Event.Events = EPOLLIN | EPOLLOUT | EPOLLRDHUP | EPOLLET
	if err := syscall.EpollCtl(p.FD, EPOLL_CTL_MOD, fd, &p.Event); err != nil {
		log.Fatalf("Modify syscall.EpollCtl error(%v)", err)
	}
}

func (p *Poller) Wait(fd int) {
	for {
		n, err := syscall.EpollWait(p.FD, p.Events[:], -1)
		if err != nil {
			log.Fatalf("syscall.EpollWait error(%v)", err)
			return
		}

		for i := 0; i < n; i++ {
			ev := p.Events[i]
			if ev.Events == 0 {
				continue
			}

			if (ev.Events&EPOLLERR) != 0 ||
				(ev.Events&EPOLLHUP) != 0 {
				log.Fatalf("syscall.EpollWait ev.Events(%v) error(%v)", ev.Events, err)
				continue
			}

			if int(ev.Fd) == fd { //accept
				// go accept(epfd, ev)
			} else if ev.Events&(EPOLLIN) != 0 { //read
				// go read(epfd, ev)
			} else if ev.Events&(EPOLLOUT) != 0 { //write
				// go write(epfd, ev)
			}
		}
	}
}

func (p *Poller) Close() {
	syscall.Close(p.FD)
}
