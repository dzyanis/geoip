package main

import (
	"log"
	"net"
	"flag"
	"github.com/dzyanis/geoip/drive"
)

var (
	httpAddress = flag.String("http.addr", ":3000", "HTTP listen address")
)

func main() {
	log.Fatal(Listen(*httpAddress))
}


func Write(c *net.Conn, d drive.GoIpInterface) {
	cc, err := d.GetCountryCode()
	if err != nil {
		log.Println(err)
		(*c).Write([]byte("error fetching ip country"))
	}

	(*c).Write([]byte(cc))
}

func Listen(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func(c net.Conn) {
			defer c.Close()

			//ip := c.RemoteAddr().(*net.TCPAddr).IP

			//Write(&c, &drive.DriveFreeGeoIp{Ip: "8.8.8.8"})
			Write(&c, &drive.DriveNekudo{Ip: "8.8.8.8"})

		}(conn)
	}

	return nil
}
