package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	// 맵의 수정을 위해 락(lock)을 시킨다.
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0
	for {
		count++
		conns.mu.Lock()
		// 알려진(확인한) 주소에 대해 루프로 반복 처리한다.
		for _, addr := range conns.addrs {
			// 모두에게 메세지를 전달한다.
			msg := fmt.Sprintf("message %d", count)
			_, err := conn.WriteToUDP([]byte(msg), addr)
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
				return
			}
		}
		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

const addr = "localhost:8888"

func main() {
	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}
	fmt.Printf("serving on %s\n", addr)

	// udp addr을 생성한다.
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	// 지정된 addr(주소)로 요청을 대기한다.
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go broadcast(conn, conns)

	msg := make([]byte, 1024)
	for {
		// 메세지를 다시 보내기 위한 주소와 포트 정보의 수집을 위해
		// 메세지를 수신한다.
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		// 수신한 메세지를 맵에 저장한다.
		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		fmt.Printf("%s connected\n", retAddr)
	}
}
