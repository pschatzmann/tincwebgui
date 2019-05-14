package utils

/**
* Utility function to print the IP address and port on which is used by the webservice
 */
import (
	"log"
	"net"
	"strings"
)

// Count - provides the number of displayed ip addresses
var IPCount uint64

// PrintIP - prints the IP address or multiple addresses if we listen on 0.0.0.0
func PrintIP(address *string) {
	if strings.HasPrefix(*address, "0.0.0.0:8000") {
		printIPs()
	} else {
		IPCount++
		log.Println("Server started: The tinc gui is available at http://" + *address)
	}
}

func printIPs() {
	ifaces, err := net.Interfaces()
	if err == nil {
		log.Println("Server started: The tinc gui is available at")
		// handle err
		for _, i := range ifaces {
			addrs, _ := i.Addrs()
			// handle err
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				IPCount++
				log.Println(" http://" + ip.String())
			}
		}
	}
}
