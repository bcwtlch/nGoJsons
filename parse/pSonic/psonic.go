package pSonic

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ijsoner"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

type Item struct {
	node *ast.Node
}

func Parse(data []byte) (ijsoner.IJsonParseRet, error) {
	item := &Item{}
	node, err := sonic.Get(data)
	item.node = &ast.Node{}
	*item.node = node
	return item, err
}

func (item *Item) Get(key string) ijsoner.IJsonParseRet {
	return &Item{node: item.node.Get(key)}
}

func (item *Item) Array() ([]ijsoner.IJsonParseRet, error) {
	nodes, err := item.node.ArrayUseNode()
	if err != nil {
		return nil, err
	}

	var items []ijsoner.IJsonParseRet
	for _, v := range nodes {
		item := &Item{}
		item.node = &ast.Node{}
		*item.node = v
		items = append(items, item)
	}
	return items, nil
}

func (item *Item) String() (string, error) {
	if item.node == nil {
		return "", fmt.Errorf("sonic-String item.node==nil")
	}
	return item.node.String()
}

func (item *Item) Bool() (bool, error) {
	if item.node == nil {
		return false, fmt.Errorf("sonic-Bool item.node==nil")
	}
	return item.node.Bool()
}

func (item *Item) Float64() (float64, error) {
	if item.node == nil {
		return 0, fmt.Errorf("sonic-Float64 item.node==nil")
	}
	return item.node.Float64()
}

func (item *Item) Int() (int, error) {
	if item.node == nil {
		return 0, fmt.Errorf("sonic-Int item.node==nil")
	}
	ret, err := item.node.Int64()
	return int(ret), err
}

func (item *Item) Uint() (uint, error) {
	if item.node == nil {
		return 0, fmt.Errorf("sonic-Uint item.node==nil")
	}
	ret, err := item.node.Int64()
	return uint(ret), err
}

func (item *Item) Int64() (int64, error) {
	if item.node == nil {
		return 0, fmt.Errorf("sonic-Int64 item.node==nil")
	}
	return item.node.Int64()
}

func (item *Item) Uint64() (uint64, error) {
	if item.node == nil {
		return 0, fmt.Errorf("sonic-Uint64 item.node==nil")
	}
	ret, err := item.node.Int64()
	return uint64(ret), err
}
