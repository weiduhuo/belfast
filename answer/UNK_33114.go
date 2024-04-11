package answer

import (
	"github.com/bettercallmolly/belfast/connection"

	"github.com/bettercallmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

var validSC33114 protobuf.SC_33114

func UNK_33114(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_33114
	response.IsWorldOpen = proto.Uint32(0)
	response.Progress = proto.Uint32(0)
	return client.SendMessage(33114, &response)
}

func init() {
	data := []byte{0x08, 0x01, 0x10, 0x86, 0x02, 0x10, 0x85, 0x01, 0x10, 0xe7, 0x05, 0x10, 0x01, 0x10, 0x95, 0x04, 0x10, 0xb5, 0x03, 0x10, 0xc3, 0x1a, 0x10, 0x64, 0x10, 0xe6, 0x0c, 0x10, 0xb1, 0x36, 0x10, 0x8d, 0x03, 0x10, 0xba, 0x31, 0x10, 0x8c, 0x11, 0x10, 0x26, 0x10, 0xf8, 0x27, 0x10, 0x82, 0x0e, 0x10, 0xc6, 0x04, 0x10, 0xec, 0x07, 0x10, 0xcc, 0x15, 0x10, 0xca, 0x04, 0x10, 0x02, 0x10, 0xda, 0x03, 0x10, 0x04, 0x10, 0x0f, 0x18, 0x14, 0x18, 0x15, 0x18, 0x1b, 0x18, 0x07, 0x20, 0x1e}
	proto.Unmarshal(data, &validSC33114)
}