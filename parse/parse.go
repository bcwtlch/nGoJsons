package parse

import (
	"github.com/bcwtlch/nGoJsons/ijsoner"
	"github.com/bcwtlch/nGoJsons/parse/pFastJson"

	"github.com/bcwtlch/nGoJsons/parse/pGeyJson"
	"github.com/bcwtlch/nGoJsons/parse/pJsonIter"
	"github.com/bcwtlch/nGoJsons/parse/pJsonParser"
	"github.com/bcwtlch/nGoJsons/parse/pSimpleJson"
	"github.com/bcwtlch/nGoJsons/parse/pSonic"
)

type config struct {
	t ParseFrame
}

type ParseFrame int

const (
	SimpleJsonFrame ParseFrame = 0
	FastJsonFrame   ParseFrame = 1
	GeyJsonFrame    ParseFrame = 2
	JsonIterFrame   ParseFrame = 3
	JsonParserFrame ParseFrame = 4
	SonicFrame      ParseFrame = 5
)

type ParseOption func(cfg *config)

func SetParseFrame(v ParseFrame) ParseOption {
	return func(cfg *config) {
		cfg.t = v
	}
}

//default JsonIterFrame

func Parse(data []byte, opts ...ParseOption) (ijsoner.IJsonParseRet, error) {
	t := JsonIterFrame
	if len(opts) > 0 {
		cfg := &config{}
		opt := opts[0]
		opt(cfg)
		t = cfg.t
	}
	return parse(data, t)
}

func ReleaseCache(jsonparser ijsoner.IJsonParseRet) {
	if parser, ok := jsonparser.(interface{ ReleaseParseCache() }); ok {
		parser.ReleaseParseCache()
	}
}

func parse(data []byte, t ParseFrame) (ijsoner.IJsonParseRet, error) {
	switch t {
	case SimpleJsonFrame:
		return pSimpleJson.Parse(data)
	case FastJsonFrame:
		return pFastJson.Parse(data)
	case GeyJsonFrame:
		return pGeyJson.Parse(data)
	case JsonIterFrame:
		return pJsonIter.Parse(data)
	case JsonParserFrame:
		return pJsonParser.Parse(data)
	case SonicFrame:
		return pSonic.Parse(data)
	default:
		return pJsonIter.Parse(data)
	}
}
