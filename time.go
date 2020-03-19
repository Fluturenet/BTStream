package main

import (
	"fmt"
	_ "github.com/beevik/ntp"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

var clockOffset time.Duration

func init() {
	log.Info("Syncing NTP.")
	for i := 0; i < 10; i++ {
		log.Info("NTP attempt ", i)
		_, err := syncNTP("pool.ntp.org", 123)
		if err == nil {
			log.Info(clockOffset)
			return
		}
	}
	log.Fatal("NTP irreversible error.")
}

func syncNTP(server string, port int) (*time.Time, error) {
	var second, fraction uint64

	packet := make([]byte, 48)
	packet[0] = 0x1B

	addr, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%s", server, strconv.Itoa(port)))
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.Write(packet)
	if err != nil {
		return nil, err
	}

	_, err = conn.Read(packet)
	if err != nil {
		return nil, err
	}

	//retrieve the bytes that we need for the current timestamp
	//data format is unsigned 64 bit long, big endian order
	//see: http://play.golang.org/p/6KRE-2Hq6n
	second = uint64(packet[40])<<24 | uint64(packet[41])<<16 | uint64(packet[42])<<8 | uint64(packet[43])
	fraction = uint64(packet[44])<<24 | uint64(packet[45])<<16 | uint64(packet[46])<<8 | uint64(packet[47])

	nsec := (second * 1e9) + ((fraction * 1e9) >> 32)

	now := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(nsec))
	localtime := time.Now().UTC()
	clockOffset = localtime.Sub(now)
	return &now, nil
}

func TimeUTP() int64 {
	time := time.Now().Add(clockOffset).UTC().Unix()
	return time
}
