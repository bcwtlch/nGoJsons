package pSimpleJson

import (
	"encoding/json"
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	"github.com/bitly/go-simplejson"
	"reflect"
)

type Item struct {
	sjson *simplejson.Json
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	jsn, err := simplejson.NewJson(data)
	return &Item{sjson: jsn}, err
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	if item.sjson == nil {
		return item
	}
	return &Item{sjson: item.sjson.Get(key)}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	if item.sjson == nil {
		return nil, fmt.Errorf("simplejson-Array() item.sjson==nil")
	}
	jsnsintf, err := item.sjson.Array()
	if err != nil {
		return nil, fmt.Errorf("simple-Array() sjon.Array err %s", err.Error())
	}

	var items []ijsoner.IJsonParseRet
	for _, v := range jsnsintf {
		newjson := simplejson.New()
		newjson.SetPath(nil, v)
		items = append(items, &Item{
			sjson: newjson,
		})
	}
	return items, nil
}

func (item *Item) String() (string, error) {
	if item.sjson == nil {
		return "", fmt.Errorf("simplejson-String item.sjson==nil")
	}

	data := item.sjson.Interface()
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		x := v.Interface()
		bytes, err := json.Marshal(x)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return item.sjson.String()
}

func (item *Item) Bool() (bool, error) {
	if item.sjson == nil {
		return false, fmt.Errorf("simplejson-Bool item.sjson==nil")
	}
	return item.sjson.Bool()
}

func (item *Item) Float64() (float64, error) {
	if item.sjson == nil {
		return 0, fmt.Errorf("simplejson-Float64 item.sjson==nil")
	}
	return item.sjson.Float64()
}

func (item *Item) Int() (int, error) {
	if item.sjson == nil {
		return 0, fmt.Errorf("simplejson-Int item.sjson==nil")
	}
	return item.sjson.Int()
}

func (item *Item) Uint() (uint, error) {
	if item.sjson == nil {
		return 0, fmt.Errorf("simplejson-Uint item.sjson==nil")
	}
	m, err := item.sjson.Uint64()
	return uint(m), err
}

func (item *Item) Int64() (int64, error) {
	if item.sjson == nil {
		return 0, fmt.Errorf("simplejson-Int64 item.sjson==nil")
	}
	return item.sjson.Int64()
}

func (item *Item) Uint64() (uint64, error) {
	if item.sjson == nil {
		return 0, fmt.Errorf("simplejson-Uint64 item.sjson==nil")
	}
	return item.sjson.Uint64()
}
