package pJsonIter

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	jsoniter "github.com/json-iterator/go"
)

type Item struct {
	iter  *jsoniter.Iterator
	value jsoniter.Any
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	item := &Item{}
	item.iter = jsoniter.ParseString(jsoniter.ConfigDefault, string(data))
	return item, nil
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	if item.iter == nil {
		return item
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return &Item{iter: item.iter, value: item.value.Get(key)}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	if item.iter == nil {
		return nil, fmt.Errorf("jsoniter-Array item.iter is nil")
	}

	if item.value == nil {
		item.value = item.iter.ReadAny()
	}

	var items []ijsoner.IJsonParseRet

	curiter := jsoniter.ParseString(jsoniter.ConfigDefault, item.value.ToString())

	for curiter.ReadArray() {
		items = append(items, &Item{iter: curiter, value: curiter.ReadAny()})
	}
	return items, nil
}

func (item *Item) String() (string, error) {
	if item.iter == nil {
		return "", fmt.Errorf("jsoniter-String item.iter==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToString(), nil
}

func (item *Item) Bool() (bool, error) {
	if item.iter == nil {
		return false, fmt.Errorf("jsoniter-Bool item.iter==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToBool(), nil
}

func (item *Item) Float64() (float64, error) {
	if item.iter == nil {
		return 0, fmt.Errorf("jsoniter->Float64() item.v==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToFloat64(), nil
}

func (item *Item) Int() (int, error) {
	if item.iter == nil {
		return 0, fmt.Errorf("jsoniter->Int() item.v==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToInt(), nil
}

func (item *Item) Uint() (uint, error) {
	if item.iter == nil {
		return 0, fmt.Errorf("jsoniter->Uint() item.v==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToUint(), nil
}

func (item *Item) Int64() (int64, error) {
	if item.iter == nil {
		return 0, fmt.Errorf("jsoniter->Int64() item.v==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToInt64(), nil
}

func (item *Item) Uint64() (uint64, error) {
	if item.iter == nil {
		return 0, fmt.Errorf("jsoniter->Uint64() item.v==nil")
	}
	if item.value == nil {
		item.value = item.iter.ReadAny()
	}
	return item.value.ToUint64(), nil
}
