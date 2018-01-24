package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func errcheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func fileWriter(filename string, content string) {
	file, err := os.Create(filename)
	errcheck(err)
	file.WriteString(content)
}

func resolve(hostname string) *net.TCPAddr {
	ip, err := net.ResolveTCPAddr("tcp4", hostname+":80")
	errcheck(err)
	fmt.Println("Resolved TCP name : " + ip.String())
	return ip
}

func pingBase(version string, hostname string, route string, filename string) string {
	connection, err := net.DialTCP("tcp", nil, resolve(hostname))
	errcheck(err)
	_, err = connection.Write([]byte("GET " + route + " HTTP/" + version + "\r\n\r\n"))
	errcheck(err)
	result, err := ioutil.ReadAll(connection)
	errcheck(err)
	if filename != "" {
		fileWriter(filename, string(result))
	}
	return string(result)
}

func main() {
	arguments := os.Args
	//print(len(arguments))
	if arguments[1] == "help" {
		fmt.Println("HTTP ping program can be used to send HTTP header pings to HTTP servers\n" +
			"Usage:\n httpping version hostname route [dump-file]")
	} else {
		//Change default port if necessary, 80 is the default HTTP port
		var data string
		if len(arguments) < 5 {
			data = pingBase(arguments[1], arguments[2], arguments[3], "")
		} else {
			data = pingBase(arguments[1], arguments[2], arguments[3], arguments[4])
		}
		fmt.Println("Response:\n" + data)
	}
}
