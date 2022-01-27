package grpcutils

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	jsonPbMarshaller = &jsonpb.Marshaler{
		EmitDefaults: true,
	}
)

// MarshalPbMessageToJsonString marshals protobuf message to json string.
func MarshalPbMessageToJsonString(msg proto.Message) string {
	msgJsonStr, _ := jsonPbMarshaller.MarshalToString(msg)
	return msgJsonStr
}

//解析返回的消息内容成json
func MarshalMessageToJsonStringForTracing(value interface{}, msgType string, maxBytes int) string {
	var messageContent string
	if msg, ok := value.(proto.Message); ok {
		if proto.Size(msg) <= maxBytes {
			messageContent = MarshalPbMessageToJsonString(msg)
		} else {
			messageContent = fmt.Sprintf(
				"[%s Message Too Large For Tracing, Max: %d bytes]",
				msgType,
				maxBytes,
			)
		}
	} else {
		messageContent = fmt.Sprintf("%v", value)
	}
	return messageContent
}