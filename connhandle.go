package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func handleConnection(conn net.Conn, conStat servStat) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println("client left..")
		conStat.disconnect()

		conn.Close()
		return
	}

	message := string(bufferBytes[:len(bufferBytes)-1])
	clientAddr := conn.RemoteAddr().String()
	response := fmt.Sprintf("message send: " + message + " from " + clientAddr)
	if strings.Compare(message, "quit") == 0 {
		conn.Write([]byte("Disconnect "))
		conn.Close()
	} else {
		conn.Write(handlehttp(message))
		handleServError(err)
	}

	log.Println(response)
	conn.Write([]byte("you sent: " + response))

	handleConnection(conn, conStat)
}

func handlehttp(keyword string) []byte {
	res, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=" + apiKeyT + "&query=" + keyword)
	handleServError(err)
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	handleServError(err)
	return sitemap
}
