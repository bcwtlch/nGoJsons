package parse

import "sync"

type ParserPool struct {
	pool sync.Pool
}

func (pp *ParserPool) Get() *Parser {
	v := pp.pool.Get()
	if v == nil {
		return &Parser{}
	}
	return v.(*Parser)
}

func (pp *ParserPool) Put(p *Parser) {
	pp.pool.Put(p)
}
