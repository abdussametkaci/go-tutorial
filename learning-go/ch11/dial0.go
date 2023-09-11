package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://www.gutenberg.org/cache/epub/16328/pg16328.txt"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		file, err := os.Create("beowulf.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Text copied to file", file.Name())
	} else {
		fmt.Println("Received unexpected status code:", response.StatusCode)
	}
}

/* // Do Not Work
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	host, port := "www.gutenberg.org", "80"
	addr := net.JoinHostPort(host, port)
	httpRequest := "GET  /cache/epub/16328/pg16328.txt HTTP/1.1\n" +
		"Host: " + host + "\n\n"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("beowulf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(file, conn)
	fmt.Println("Text copied to file", file.Name())
}
*/
