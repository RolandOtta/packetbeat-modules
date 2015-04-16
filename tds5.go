package main

import (

	//        "github.com/packetbeat/"
	//"bytes"
//	"encoding/hex"
//	"strconv"
        "time"
	//"bytes"
	"labix.org/v2/mgo/bson"
)


// tds tokens
const (
	TDS_ALTCONTROL = 0xAF //not yet supported
	TDS_ALTFMT = 0xA8 //not yet supported
	TDS_ALTNAME = 0xA7 //not yet supported
	TDS_ALTROW = 0xD3 //not yet supported
	TDS_CAPABILITY = 0xE2 //not yet supported
	TDS_COLFMT = 0xA1 //not yet supported
	TDS_COLFMTOLD = 0x2A //not yet supported
	TDS_COLINFO = 0xA5 //not yet supported
	TDS_COLINFO2 = 0x20 //not yet supported
	TDS_COLNAME = 0xA0 //not yet supported
	TDS_CONTROL = 0xAE //not yet supported
	TDS_CURCLOSE = 0x80 //not yet supported
	TDS_CURDECLARE = 0x86 //not yet supported
	TDS_CURDECLARE2 = 0x23 //not yet supported
	TDS_CURDECLARE3 = 0x10 //not yet supported
	TDS_CURDELETE = 0x81 //not yet supported
	TDS_CURFETCH = 0x82 //not yet supported
	TDS_CURINFO = 0x83 //not yet supported
	TDS_CURINFO2 = 0x87 //not yet supported
	TDS_CURINFO3 = 0x88 //not yet supported
	TDS_CUROPEN = 0x84 //not yet supported
	TDS_CURUPDATE = 0x85 //not yet supported
	TDS_DBRPC = 0xE6 //not yet supported
	TDS_DBRPC2 = 0xE8 //not yet supported
	TDS_DEBUGCMD = 0x60 //not yet supported
	TDS_DONE = 0xFD //not yet supported
	TDS_DONEINPROC = 0xFF //not yet supported
	TDS_DONEPROC = 0xFE //not yet supported
	TDS_DYNAMIC = 0xE7 //not yet supported
	TDS_DYNAMIC2 = 0x62 //not yet supported
	TDS_EED = 0xE5 //not yet supported
	TDS_ENVCHANGE = 0xE3 //not yet supported
	TDS_ERROR = 0xAA //not yet supported
	TDS_EVENTNOTICE = 0xA2 //not yet supported
	TDS_INFO = 0xAB //not yet supported
	TDS_KEY = 0xCA //not yet supported
	TDS_LANGUAGE = 0x21 //not yet supported
	TDS_LOGINACK = 0xAD //not yet supported
	TDS_LOGOUT = 0x71 //not yet supported
	TDS_MSG = 0x65 //not yet supported
	TDS_OFFSET = 0x78 //not yet supported
	TDS_OPTIONCMD = 0xA6 //not yet supported
	TDS_OPTIONCMD2 = 0x63 //not yet supported
	TDS_ORDERBY = 0xA9 //not yet supported
	TDS_ORDERBY2 = 0x22 //not yet supported
	TDS_PARAMFMT = 0xEC //not yet supported
	TDS_PARAMFMT2 = 0x20 //not yet supported
	TDS_PARAMS = 0xD7 //not yet supported
	TDS_PROCID = 0x7C //not yet supported
	TDS_RETURNSTATUS = 0x79 //not yet supported
	TDS_RETURNVALUE = 0xAC  //not yet supported
	TDS_RPC = 0xE0 //not yet supported
	TDS_ROW = 0xD1 //not yet supported
	TDS_ROWFMT = 0xEE //not yet supported
	TDS_ROWFMT2 = 0x61 //not yet supported
	TDS_TABNAME = 0xA4 //not yet supported
)

// rpc options for TDS_DBRPC
const (
	TDS_RPC_UNUSED = 0x0000 //no options used
	TDS_RPC_RECOMPILE = 0x0001 //recompile rpc before execution
	TDS_RPC_PARAMS = 0x0002 //there are parameters for the rpc
)

type Tds5Stream struct {
	tcpStream *TcpStream

	requestData [] byte
	responseData [] byte
	message *Tds5Message
}

type Tds5Message struct {
        reqSqlText string
	reqTokenSupported bool
	reqTdsToken byte
	respTdsToken byte
	respErrorCode int
	respErrorInfo string
	ts time.Time
	respNumOfRows int
	reqSize uint64
	respSize uint64
	expectedNextToken byte

	Publisher *PublisherType
}

func ParseTds5(pkt *Packet, tcp *TcpStream, dir uint8) {

        defer RECOVER("ParseTds5 exception")

        if tcp.tds5Data[dir] == nil {
                tcp.tds5Data[dir] = &Tds5Stream{
                        tcpStream: tcp,
                        requestData:      pkt.payload,
                        message:   &Tds5Message{ts: pkt.ts},
                }
                DEBUG("tds5", "New stream created")
        }

	stream :=  tcp.tds5Data[dir]
	
	ok, complete := tds5MessageParser(stream)	
	
	if ok && complete {
		DEBUG("tds5","publishing message")
		//TODO nicht jedesmal initalisieren
		stream.message.Publisher = &Publisher
		publish(stream.message)
		//TODO: move to the right place - after finished parsing
		tcp.tds5Data[dir] = nil
		stream.PrepareForNewMessage()
	}
	
}

// returns ok, complete
func tds5MessageParser(s *Tds5Stream) (bool, bool) {
	DEBUG ("tds5Test", "parsing message")
	
	requestData := s.requestData
	tdsToken := requestData[8]
	message := s.message

        switch tdsToken {
        case TDS_LANGUAGE:
		// return nicht complete und handle response
		//DEBUG("tds5", "Received a TDS_LANGUAGE token")
		//DEBUG("tds5", "SQLText: " + string(requestData[14:]))
		DEBUG("tds5", "sent response directly for TDS_LANG // todo")
	case TDS_DBRPC: 
		// return nicht complete und handle response
		//DEBUG("tds5", "Received a TDS_DBRPC")
		//DEBUG("tds5", "Length: " + strconv.Itoa(int(Bytes_Ntohs(requestData[9:11]))))
		//sqlTextLength := int(requestData[11]);
		//DEBUG("tds5", "SQLText Length: " + strconv.Itoa(sqlTextLength) )
		//DEBUG("tds5", "Command: " + string(requestData[12:12+sqlTextLength]))
		//DEBUG("tds5", "TDS Options: 0x" + hex.EncodeToString(requestData[12+sqlTextLength:12+sqlTextLength+2]))
		//DEBUG("tds5Test", "type: " + hex.EncodeToString(requestData[12+sqlTextLength+2:]))
		DEBUG("tds5", "sent response directly for TDS_DBRPC // todo")
	default:
		// wenn kein req token, dann setzen und return incomplete
		// sonst setzte resp token und return complete
		//DEBUG("tds5Test", hex.EncodeToString(requestData[8:9]))
		if message.reqTdsToken == 0 {
			DEBUG("tds5", "do not publish ... unknown request ... waiting for unknonw reply")
			message.reqTdsToken = tdsToken
			return true, false
		} else {
			DEBUG("tds5", "yeah .. got unknown reply ... publish")
			//DEBUG("tds5","reqtoken",  message.reqTdsToken)
			// todo wenns schon einen resp token gab fehler werfen
			message.respTdsToken = tdsToken
			return true, true
		}
	}


	// todo: das sollte nie aufgerufen werden
	// ob das complete ist, wird normalerweise im tdstoken entschieden
	return true, true
}

func publish(msg *Tds5Message) {

      	DEBUG("tds5", "....")

		event := Event{}

		event.Type = "tds5"
		
		event.Status = OK_STATUS
		event.ResponseTime = 100
		event.Tds5 = bson.M{}

			event.Tds5 = bson.M{
				"request": bson.M{
					"method": "blah",
					"params": "blah",
					"size":   "blah",
				},
				"service": "blah",
			}

        Src := Endpoint{
                Ip:   "127.0.0.1",
                Port: 5001,
                Proc: "tds2",
        }
        Dst := Endpoint{
                Ip:   "127.0.0.1",
                Port: 5000,
                Proc: "tds3",
        }
		if msg.Publisher != nil {
			msg.Publisher.PublishEvent(msg.ts, Src, Dst, &event)
		}

		DEBUG("tds5", "Published event")

}

func (stream *Tds5Stream) PrepareForNewMessage() {
        //stream.data = stream.data[stream.message.end:]
        //stream.parseState = PgsqlStartState
        //stream.parseOffset = 0
        stream.message = nil
}

