package main

import (
	"bufio"
	"fmt"
	"net"
	//	"time"
	//"os"
	"strconv"
	"sync"
)

func handleTCPConnection(conn *net.TCPConn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)
	fmt.Println(conn.LocalAddr().String())
	fmt.Println(conn.RemoteAddr().String())
	conn.SetWriteBuffer(1024)
	conn.SetReadBuffer(1024)
	for {
		var data = make([]byte, 1024)
		_, err := rd.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
		conn.Write(data)
	}
}

func handleUDPConnection(conn *net.UDPConn, once chan int) {
	defer conn.Close()
	/*
		conn.SetDeadline(time.Now().Add(10000))
		conn.SetReadDeadline(time.Now().Add(10000))
		conn.SetWriteDeadline(time.Now().Add(10000))
	*/
	conn.SetReadBuffer(1024)
	conn.SetWriteBuffer(1024)
	for {
		buff := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Connecting from UDP port: ")
			fmt.Println(err)
			break
		}
		fmt.Println(addr.String())
		fmt.Println(string(buff))
		conn.WriteToUDP(buff, addr)
	}
	once <- 1
}

func startTCPServer(port int, done chan int) {
	//addr, err := net.ResolveTCPAddr("tcp4", "10.71.1.174:"+strconv.Itoa(port))
	addr, err := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(port))
	ln, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println(err)
		done <- 1
		return
		//os.Exit(-1)
	}

	defer ln.Close()

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			done <- 1
			return
		}
		fmt.Println("Connecting from TCP port: ", port)
		go handleTCPConnection(conn)
	}

	done <- 1
	fmt.Println("++++++++++++++")
}

func startUDPServer(port int, done chan int) {
	once := make(chan int)
	/*
		addr, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
		//addr, err := net.ResolveUDPAddr("udp4", "10.71.1.174:"+strconv.Itoa(port))
		if err != nil {
			fmt.Println(err)
			//os.Exit(-1)
			done <- 1
			return
		}
	*/

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	//conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		done <- 1
		return
		//os.Exit(-1)

	}

	fmt.Println("Connecting from UDP port: ", port)
	go handleUDPConnection(conn, once)
	<-once
	done <- 1
}

func main() {
	var wait = sync.WaitGroup{}
	var done = make(chan int)
	for i := 50200; i < 50400; i++ {
		wait.Add(1)
		go startUDPServer(i, done)
	}

	for i := 50500; i < 50600; i++ {
		wait.Add(1)
		go startTCPServer(i, done)
	}

	for {
		select {
		case <-done:
			fmt.Println("++++++++++++++")
			wait.Done()
		}
	}
	wait.Wait()
}
