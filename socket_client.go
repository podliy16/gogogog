package main

import (
    "net"
//	"strconv"
	"fmt"
	"io"
	"github.com/yasushi-saito/fifo_queue"
)

const PORTS_NUMBER = 1

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
	ports := [PORTS_NUMBER]string{"5270"}
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
	return connects, nil
}

//connect

func writeToConnect(data string, connect *Connect,results chan string,checkServ chan *Connect){
	go Reader(*connect, results,checkServ)
	connect.state = false
	data = "-u " + data
	connect.Write([]byte(data))
}

func SendData(data string, connects *[PORTS_NUMBER]Connect, results chan string,
			checkServ chan *Connect) (success bool, err error) {
	success = false
	for i:=0; i < PORTS_NUMBER; i++ {
		connect := &connects[i]
		if connect.state == false {
			continue
		}
		writeToConnect(data,connect,results,checkServ)
		if err != nil {
			return false, err
		}
		success = true
		break
	} 
	return success, nil
}

func CheckServ(c chan *Connect,results chan string,q *fifo_queue.Queue,len_q int){
	for len_q != 0{
		connect := <-c
		if q.Len()!=0{
			writeToConnect(q.PopFront().(string),connect,results,c)		 		
		}
		len_q -- 
	}
}

func Reader(r Connect, results chan string,checkServ chan *Connect) {
    buf := make([]byte, 1024)
    n, err := r.Read(buf[:])
	r.state = true
	
	if err == io.EOF{
		return
	}
	
	if err != nil {
        panic(err)
    }
	
	checkServ <- &r
	results <- string(buf[0:n])
}


func socketMain(DoneLinks []LinksType) (<-chan string, int) {
	results := make(chan string)
	checkServ := make(chan *Connect)

	q := fifo_queue.NewQueue()
	
	for _, Link := range DoneLinks{
		q.PushBack(Link.Link)
	}
	len_q := q.Len()
	
	connects, err := CreateConnects()
	
	if err != nil {
		panic(err)
	}
	
	var success bool
	
	go CheckServ(checkServ,results,q,len_q)
		
	//Test closed connections. First 4 success -- true, last one -- false
	for i:=0;i<len_q;i++{
		item := q.PopFront().(string)
		success, err = SendData(item, &connects, results,checkServ)
		if success == false{
			q.PushBack(item)
		}
	}
	
	if err!=nil{
		fmt.Println(err)
	}
	
	return results, len_q
}
