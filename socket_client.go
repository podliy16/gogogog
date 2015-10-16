package main

import (
    "net"
	"fmt"
)

const PORTS_NUMBER = 4

type Connect struct {
	state bool
	net.Conn
}

func CloseConnects(connects [PORTS_NUMBER]Connect) (err error) {
	for i:=0; i < PORTS_NUMBER; i++ {
		connects[i].Close()
	}
	return nil
}

func CreateConnects() (connects [PORTS_NUMBER]Connect, err error) {
	ports := [PORTS_NUMBER]string{"52631", "52632", "52633", "52635"}
	for i:=0; i < PORTS_NUMBER; i++ {
		port := ports[i]
		fmt.Println("localhost:" + port)
		c, err := net.Dial("tcp", "localhost:" + port)
		if err != nil {
			return connects, err
		}
		connect := Connect{true, c}
		connects[i] = connect
	}
	fmt.Println(connects)
	return connects, nil
}

func SendData(data string, connects *[PORTS_NUMBER]Connect, results chan string) (success bool, err error) {
	success = false
	for i:=0; i < PORTS_NUMBER; i++ {
		connect := &connects[i]
		if connect.state == false {
			continue
		}
		go Reader(*connect, results)
		connect.state = false
		connect.Write([]byte(data))
		if err != nil {
			return false, err
		}
		success = true
		break
	}
	return success, nil
}


func Reader(r Connect, results chan string) {
    buf := make([]byte, 1024)
    n, err := r.Read(buf[:])
    if err != nil {
        panic(err)
    }
	results <- string(buf[0:n])
}


func main() {
	results := make(chan string, 100)
	connects, err := CreateConnects()
	if err != nil {
		panic(err)
	}
	var success bool
	//Test closed connections. First 4 success -- true, last one -- false
	success, err = SendData("hi", &connects, results)
	fmt.Println(connects, success)
	success, err = SendData("hi2", &connects, results)
	fmt.Println(connects, success)
	success, err = SendData("HI3", &connects, results)
	fmt.Println(connects, success)
	success, err = SendData("HI4", &connects, results)
	fmt.Println(connects, success)
	success, err = SendData("HI5", &connects, results)
	fmt.Println(connects, success)
	for result := range results {
		fmt.Println("Client got: " + result)
	}
}