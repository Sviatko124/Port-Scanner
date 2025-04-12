# Port-Scanner
A simple tool written in Go to scan a target for all potentially open ports using connect scanning. I primarily wrote the tool to experiment and learn Go's "net" library. The tool launches multiple threads which are given evenly distributed work loads, and the port scan range is from 1 to the given number of ports to scan argument. 

## Build/Usage
Build with:
`go build port_scanner.go`

To run the tool:
`port_scanner.exe <target-IP> <# of threads> <number of ports to scan>`
