package main

import "log"

type servStat struct {
	connCount int
	connNow   int
}

func (connStat *servStat) connect() {
	connStat.connCount++
	connStat.connNow++
	log.Printf("connect count : %d connect now: %d", connStat.connCount, connStat.connNow)
}

func (connStat *servStat) disconnect() {
	connStat.connNow--
	log.Printf("connect count : %d connect now: %d", connStat.connCount, connStat.connNow)
}
