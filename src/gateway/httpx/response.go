package httpx

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"
	"net/http"
)

type responseOpts struct {
	code int
	proto proto.Message
	message string
}

type ResponseOpts func(*responseOpts)

func WithCode(code int) ResponseOpts {
	return func(o *responseOpts) {
		o.code = code
	}
}

func WithMessage(msg interface{}) ResponseOpts {
	return func(o *responseOpts) {
		switch v := msg.(type) {
		case proto.Message:
			o.proto = v
		case string:
			o.message = fmt.Sprintf("%v", msg)
		}
	}
}

func Ok(w http.ResponseWriter, opts ...ResponseOpts) {
	var o responseOpts
	for _, opt := range opts {
		opt(&o)
	}

	if o.code == 0 {
		o.code = 200
	}
	w.WriteHeader(o.code)

	if o.proto != nil {
		var m jsonpb.Marshaler
		err := m.Marshal(w, o.proto)
		if err != nil {
			Fail(w, err)
		}
		return
	}

	if o.message == "" {
		o.message = "Success"
	}

	_, err := w.Write([]byte(o.message))
	if err != nil {
		Fail(w, err)
		return
	}
}

func Fail(w http.ResponseWriter, err error, opts ...ResponseOpts) {
	var o responseOpts
	for _, opt := range opts {
		opt(&o)
	}

	if o.message == "" {
		o.message = "An unexpected error occurred. Please try again later."
	}

	if o.code == 0 {
		o.code = 500
	}

	log.Error().Err(err).Msg("api error")

	http.Error(w, o.message, o.code)
}
