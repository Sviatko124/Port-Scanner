package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// get arguments and convert thread_count and max_port into integers
	target := os.Args[1]
	thread_count_str := os.Args[2]
	thread_count, err := strconv.Atoi(thread_count_str)
	if err != nil {
		panic(err)
	}
	max_port_str := os.Args[3]
	max_port, err := strconv.Atoi(max_port_str)
	if err != nil {
		panic(err)
	}
	fmt.Println(target)
	fmt.Println(max_port)

	var wg sync.WaitGroup
	open_ports := make(map[int]bool)
	// connect to the server on all ports in a loop
	for i := 1; i <= thread_count; i++ {
		ports_per_thread := int(math.Ceil(float64(max_port) / float64(thread_count)))
		start_port := (i - 1) * ports_per_thread
		end_port := i * ports_per_thread
		if end_port > max_port {
			end_port = max_port
		}
		wg.Add(1)
		go scan_ports(&wg, target, start_port, end_port, open_ports)
	}
	wg.Wait()
	fmt.Println("Scanning finished.")
}

func scan_ports(wg *sync.WaitGroup, target string, start_port, max_port int, open_ports map[int]bool) {
	defer wg.Done()
	for i := start_port; i <= max_port; i++ {
		port := strconv.Itoa(i)
		fullHostname := target + ":" + port
		// attempt to connect to the port
		connection, err := net.DialTimeout("tcp", fullHostname, time.Second)
		if err != nil {
			// port closed
			continue
		}
		if connection != nil {
			defer connection.Close()
			port_int, err := strconv.Atoi(port)
			if err != nil {
				panic(err)
			}
			if _, ok := open_ports[port_int]; !ok {
				fmt.Printf("Port %s open\n", port)
				open_ports[port_int] = true
			}
		}
	}
}
