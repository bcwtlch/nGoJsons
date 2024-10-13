package pFastJson

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	"github.com/valyala/fastjson"
)

type Item struct {
	v   *fastjson.Value
	err error
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	v, err := fastjson.Parse(string(data))
	if err != nil {
		return nil, err
	}
	return &Item{v: v, err: nil}, nil
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	if item == nil {
		return &Item{err: fmt.Errorf("fastjson->Get item==nil,key=%s", key)}
	}

	if item.v == nil || item.err != nil {
		return item
	}
	return &Item{v: item.v.Get(key)}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	if item == nil {
		return nil, fmt.Errorf("fastjson->Array item==nil")
	}
	if item.v == nil {
		return nil, fmt.Errorf("fastjson->Array() item.v==nil")
	}
	arr, err := item.v.Array()
	if err != nil {
		return nil, err
	}
	var items []ijsoner.IJsonParseRet
	for _, v := range arr {
		items = append(items, &Item{v: v, err: nil})
	}
	return items, nil
}

func (item *Item) String() (string, error) {
	if item == nil {
		return "", fmt.Errorf("fastjson->String item==nil")
	}

	if item.v == nil {
		return "", fmt.Errorf("fastjson->String() item.v==nil")
	}
	str := item.v.String()
	return str, nil
}

func (item *Item) Bool() (bool, error) {
	if item == nil {
		return false, fmt.Errorf("fastjson->Bool item==nil")
	}

	if item.v == nil {
		return false, fmt.Errorf("fastjson->Bool() item.v==nil")
	}
	return item.v.Bool()
}

func (item *Item) Float64() (float64, error) {
	if item == nil {
		return 0, fmt.Errorf("fastjson->Float64 item==nil")
	}
	if item.v == nil {
		return 0, fmt.Errorf("fastjson->Float64() item.v==nil")
	}
	return item.v.Float64()
}

func (item *Item) Int() (int, error) {
	if item == nil {
		return 0, fmt.Errorf("fastjson->Int item==nil")
	}
	if item.v == nil {
		return 0, fmt.Errorf("fastjson->Int() item.v==nil")
	}
	return item.v.Int()
}

func (item *Item) Uint() (uint, error) {
	if item == nil {
		return 0, fmt.Errorf("fastjson->Uint item==nil")
	}
	if item.v == nil {
		return 0, fmt.Errorf("fastjson->Uint() item.v==nil")
	}
	return item.v.Uint()
}

func (item *Item) Int64() (int64, error) {
	if item == nil {
		return 0, fmt.Errorf("fastjson->Int64 item==nil")
	}
	if item.v == nil {
		return 0, fmt.Errorf("fastjson->Int64() item.v==nil")
	}
	return item.v.Int64()
}

func (item *Item) Uint64() (uint64, error) {
	if item == nil {
		return 0, fmt.Errorf("fastjson->Uint64 item==nil")
	}
	if item.v == nil {
		return 0, fmt.Errorf("fastjson->Uint64() item.v==nil")
	}
	return item.v.Uint64()
}
