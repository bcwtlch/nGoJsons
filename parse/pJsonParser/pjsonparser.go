package pJsonParser

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	"github.com/buger/jsonparser"
)

type Item struct {
	data []byte
	err  error
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	retbytes, _, _, err := jsonparser.Get(data)
	if err != nil {
		return nil, err
	}
	return &Item{data: retbytes}, nil
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	if item.data == nil || item.err != nil {
		return item
	}

	retbytes, _, _, err := jsonparser.Get(item.data, key)
	return &Item{data: retbytes, err: err}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	if item.data == nil || item.err != nil {
		if item.err != nil {
			return nil, item.err
		}
		if item.data == nil {
			return nil, fmt.Errorf("jsonparser-Array item.data==nil")
		}
	}

	var items []ijsoner.IJsonParseRet

	_, err := jsonparser.ArrayEach(item.data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		items = append(items, &Item{data: value, err: err})
	})
	if err != nil {
		return nil, fmt.Errorf("jsonparser-Array ArrayEach fail %s", err.Error())
	}
	return items, nil
}

func (item *Item) Bool() (bool, error) {
	if item.err != nil {
		return false, fmt.Errorf("jsonparser-Bool item.err %s", item.err.Error())
	}
	return jsonparser.ParseBoolean(item.data)
}

func (item *Item) String() (string, error) {
	if item.err != nil {
		return "", fmt.Errorf("jsonparser-String item.err %s", item.err.Error())
	}
	return jsonparser.ParseString(item.data)
}

func (item *Item) Float64() (float64, error) {
	if item.err != nil {
		return 0, fmt.Errorf("jsonparser-Float64 item.err %s", item.err.Error())
	}
	return jsonparser.ParseFloat(item.data)
}

func (item *Item) Int64() (int64, error) {
	if item.err != nil {
		return 0, fmt.Errorf("jsonparser-Int64 item.err %s", item.err.Error())
	}
	return jsonparser.ParseInt(item.data)
}

func (item *Item) Int() (int, error) {
	if item.err != nil {
		return 0, fmt.Errorf("jsonparser-Int item.err %s", item.err.Error())
	}

	n, err := jsonparser.ParseInt(item.data)
	return int(n), err
}

func (item *Item) Uint() (uint, error) {
	if item.err != nil {
		return 0, fmt.Errorf("jsonparser-Uint item.err %s", item.err.Error())
	}

	n, err := jsonparser.ParseInt(item.data)
	return uint(n), err
}

func (item *Item) Uint64() (uint64, error) {
	if item.err != nil {
		return 0, fmt.Errorf("jsonparser-Uint64 item.err %s", item.err.Error())
	}

	n, err := jsonparser.ParseInt(item.data)
	return uint64(n), err
}
