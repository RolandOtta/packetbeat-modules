package main

import (
//        "github.com/packetbeat/"
)

func ParseTds5(pkt *Packet, tcp *TcpStream, dir uint8) {

        defer RECOVER("ParsePgsql exception")

        DEBUG("tds5", "Parsing tds5 data stream")

}

type Tds5Stream struct {
        tcpStream *TcpStream

        data []byte

        parseOffset  int
        parseState   int
        bodyReceived int

        message *HttpMessage
}

