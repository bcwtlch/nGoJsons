package nGoJsons

import (
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"io"
)

/*
stdlib
gojson
sonic
json-iter

*/

type JsonFrame int

const (
	StdlibJsonFrame   JsonFrame = 0
	GoJsonFrame       JsonFrame = 1
	SonicJsonFrame    JsonFrame = 2
	JsonIterJsonFrame JsonFrame = 3
)

type config struct {
	t JsonFrame
}

type Option func(cfg *config)

type Decoder interface {
	Decode(val interface{}) error
	Buffered() io.Reader
	DisallowUnknownFields()
	More() bool
	UseNumber()
}

type Encoder interface {
	Encode(val interface{}) error
	SetEscapeHTML(on bool)
	SetIndent(prefix, indent string)
}

func SetJsonFrame(v JsonFrame) Option {
	return func(cfg *config) {
		cfg.t = v
	}
}

func HTMLEscape(dst *bytes.Buffer, src []byte, opts ...Option) {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		json.HTMLEscape(dst, src)
	case GoJsonFrame:
		gojson.HTMLEscape(dst, src)
	default:
		json.HTMLEscape(dst, src)
	}
}

func Marshal(v interface{}, opts ...Option) ([]byte, error) {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.Marshal(v)
	case GoJsonFrame:
		return gojson.Marshal(v)
	case SonicJsonFrame:
		return sonic.Marshal(v)
	case JsonIterJsonFrame:
		return jsoniter.Marshal(v)
	default:
		return json.Marshal(v)
	}
}

func MarshalIndent(v interface{}, prefix, indent string, opts ...Option) ([]byte, error) {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.MarshalIndent(v, prefix, indent)
	case GoJsonFrame:
		return gojson.MarshalIndent(v, prefix, indent)
	case SonicJsonFrame:
		return sonic.MarshalIndent(v, prefix, indent)
	case JsonIterJsonFrame:
		return jsoniter.MarshalIndent(v, prefix, indent)
	default:
		return json.MarshalIndent(v, prefix, indent)
	}
}

func Indent(dst *bytes.Buffer, src []byte, prefix, idx string, opts ...Option) error {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.Indent(dst, src, prefix, idx)
	case GoJsonFrame:
		return gojson.Indent(dst, src, prefix, idx)
	default:
		return json.Indent(dst, src, prefix, idx)
	}
}

func Valid(data []byte, opts ...Option) bool {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.Valid(data)
	case GoJsonFrame:
		return gojson.Valid(data)
	case SonicJsonFrame:
		return sonic.Valid(data)
	case JsonIterJsonFrame:
		return jsoniter.Valid(data)
	default:
		return json.Valid(data)
	}
}

func Unmarshal(data []byte, v interface{}, opts ...Option) error {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.Unmarshal(data, v)
	case GoJsonFrame:
		return gojson.Unmarshal(data, v)
	case SonicJsonFrame:
		return sonic.Unmarshal(data, v)
	case JsonIterJsonFrame:
		return jsoniter.Unmarshal(data, v)
	default:
		return json.Unmarshal(data, v)
	}
}

/**************  stream  start *********************/

func NewDecoder(r io.Reader, opts ...Option) Decoder {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.NewDecoder(r)
	case GoJsonFrame:
		return gojson.NewDecoder(r)
	case SonicJsonFrame:
		return sonic.ConfigStd.NewDecoder(r)
	case JsonIterJsonFrame:
		return jsoniter.NewDecoder(r)
	default:
		return json.NewDecoder(r)
	}
}

func NewEncoder(w io.Writer, opts ...Option) Encoder {
	mJsonFrame := StdlibJsonFrame
	if len(opts) > 0 {
		cfg := &config{}
		opts[0](cfg)
		mJsonFrame = cfg.t
	}
	switch mJsonFrame {
	case StdlibJsonFrame:
		return json.NewEncoder(w)
	case GoJsonFrame:
		return gojson.NewEncoder(w)
	case SonicJsonFrame:
		return sonic.ConfigStd.NewEncoder(w)
	case JsonIterJsonFrame:
		return jsoniter.NewEncoder(w)
	default:
		return json.NewEncoder(w)
	}
}

/**************  stream  end *********************/
