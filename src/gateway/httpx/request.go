package httpx

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/text/language"
	"net"
	"net/http"
)

type Request struct {
	*http.Request

	UserId    int64
	SessionId string
	Lang      language.Tag
	IP        net.IP
}

func ParseJSON(r *Request,
	pb proto.Message) error {
	return jsonpb.Unmarshal(r.Request.Body, pb)
}
