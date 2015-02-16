package main

import (

	//        "github.com/packetbeat/"
	//"bytes"
	"encoding/hex"
	"strconv"
	//"bytes"
)

// tds tokens
const (
	TDS_ALTCONTROL = 0xAF 
	TDS_ALTFMT = 0xA8
	TDS_ALTNAME = 0xA7
	TDS_ALTROW = 0xD3
	TDS_CAPABILITY = 0xE2
	TDS_COLFMT = 0xA1
	TDS_COLFMTOLD = 0x2A
	TDS_COLINFO = 0xA5
	TDS_COLINFO2 = 0x20
	TDS_COLNAME = 0xA0
	TDS_CONTROL = 0xAE
	TDS_CURCLOSE = 0x80
	TDS_CURDECLARE = 0x86
	TDS_CURDECLARE2 = 0x23
	TDS_CURDECLARE3 = 0x10
	TDS_CURDELETE = 0x81
	TDS_CURFETCH = 0x82
	TDS_CURINFO = 0x83
	TDS_CURINFO2 = 0x87
	TDS_CURINFO3 = 0x88
	TDS_CUROPEN = 0x84
	TDS_CURUPDATE = 0x85
	TDS_DBRPC = 0xE6
	TDS_DBRPC2 = 0xE8
	TDS_DEBUGCMD = 0x60
	TDS_DONE = 0xFD
	TDS_DONEINPROC = 0xFF
	TDS_DONEPROC = 0xFE
	TDS_DYNAMIC = 0xE7
	TDS_DYNAMIC2 = 0x62
	TDS_EED = 0xE5
	TDS_ENVCHANGE = 0xE3
	TDS_ERROR = 0xAA
	TDS_EVENTNOTICE = 0xA2
	TDS_INFO = 0xAB
	TDS_KEY = 0xCA
	TDS_LANGUAGE = 0x21
	TDS_LOGINACK = 0xAD
	TDS_LOGOUT = 0x71
	TDS_MSG = 0x65
	TDS_OFFSET = 0x78
	TDS_OPTIONCMD = 0xA6
	TDS_OPTIONCMD2 = 0x63
	TDS_ORDERBY = 0xA9
	TDS_ORDERBY2 = 0x22
	TDS_PARAMFMT = 0xEC
	TDS_PARAMFMT2 = 0x20
	TDS_PARAMS = 0xD7
	TDS_PROCID = 0x7C 
	TDS_RETURNSTATUS = 0x79
	TDS_RETURNVALUE = 0xAC 
	TDS_RPC = 0xE0 
	TDS_ROW = 0xD1
	TDS_ROWFMT = 0xEE
	TDS_ROWFMT2 = 0x61
	TDS_TABNAME = 0xA4
)

type Tds5Stream struct {
	tcpStream *TcpStream

	message *PgsqlMessage
}

type Tds5Message {
        Query          string
        Size           uint64
        Fields         []string
        FieldsFormat   []byte
        Rows           [][]string
        NumberOfRows   int
        NumberOfFields int
        IsOK           bool
        IsError        bool
        ErrorInfo      string
        ErrorCode      string
        ErrorSeverity  string

        Direction    uint8
        Incomplete   bool
        TcpTuple     TcpTuple
        CmdlineTuple *CmdlineTuple	
}

func ParseTds5(pkt *Packet, tcp *TcpStream, dir uint8) {

        defer RECOVER("ParseTds5 exception")

        DEBUG("tds5", "### Start Parsing tds5 data stream")
        //bts := pkt.payload
        //DEBUG("tds5Test", string(bts[:]))
        //s := string(pkt.payload[:])
        //DEBUG("tds5Test", s)
        //s = hex.EncodeToString(pkt.payload)
        //DEBUG("tds5Test", s)
        //s = hex.EncodeToString(pkt.payload[0:1])
        //DEBUG("tds5Test", s)
        //s = hex.EncodeToString(pkt.payload[1:2])
        //DEBUG("tds5Test", s)
        //s = hex.EncodeToString(pkt.payload[2:4])

        DEBUG("tds5Test", hex.EncodeToString(pkt.payload[0:]))
        DEBUG("tds5Test", "type: " + hex.EncodeToString(pkt.payload[8:9]) + " command: " + string(pkt.payload[14:]) + " length: " + strconv.Itoa(int(Bytes_Ntohl(pkt.payload[9:13]))))

        if ( pkt.payload[8] == TDS_LANGUAGE ) {
                DEBUG("tds5Test", "!!!!!!!!!!!!!")
        }
        DEBUG("tds5", "### End Parsing tds5 data stream")
}
