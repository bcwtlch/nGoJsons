package pGeyJson

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	geyjson "github.com/bcwtlch/nGoJsons/ngeyjson/parse"
)

type Item struct {
	*geyjson.Node
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	node, err := geyjson.Parse(data)
	return &Item{node}, err
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	if item == nil {
		nitem := &Item{Node: &geyjson.Node{}}
		nitem.Node.SetErr(fmt.Errorf("fastjson->Get item==nil,key=%s", key))
		return nitem
	}

	node := item.Node.Get(key)
	return &Item{node}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	if item == nil {
		return nil, fmt.Errorf("geyjson->Array item==nil")
	}
	nodes, err := item.Node.Array()
	if err != nil {
		return nil, err
	}
	var items []ijsoner.IJsonParseRet
	for _, v := range nodes {
		items = append(items, &Item{v})
	}
	return items, nil
}
