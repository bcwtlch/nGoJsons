package parse

import (
	"fmt"
	"github.com/bcwtlch/nGoJsons/ngeyjson/parse/fastfloat"
	"strconv"
	"strings"
	"unicode/utf16"
)

/**************************************************
Design planning:
  For json,the key of the operation is used as a constant comparison.
  And the value of v corresponding to k is taken as the obtained result value.
  Generally, the value of key is relatively small, while the value of v may be relatively large.
  it is considered to design the value of v as a digital range of strings without assigning values first,


eg:
 json: `{"foo":[{"bar":{"baz":123,"x":"434"},"y":[]},[null, false]],"qwe":false}`

top parse.Value:
   o    --  json subobject
   a    --  nil
   spos --  1  //start pos
   epos --  71 //end pos
   t    --  TypeObject

****************************************************/

type kv struct {
	k string
	v *Value
}

type Object struct {
	kvs           []kv
	keysUnescaped bool
}

type Value struct {
	o    Object
	a    []*Value
	spos int32
	epos int32
	t    Type
}

type Parser struct {
	// b contains working  the []byte to be parsed.
	b []byte

	s string
	// c is a cache for json values.
	c cache
}

type Node struct {
	p     *Parser
	value *Value
	err   error
}

type cache struct {
	vs []Value
}

// Type represents JSON type.
type Type int

const (
	// TypeNull is JSON null.
	TypeNull Type = 0

	// TypeObject is JSON object type.
	TypeObject Type = 1

	// TypeArray is JSON array type.
	TypeArray Type = 2

	// TypeString is JSON string type.
	TypeString Type = 3

	// TypeNumber is JSON number type.
	TypeNumber Type = 4

	// TypeTrue is JSON true.
	TypeTrue Type = 5

	// TypeFalse is JSON false.
	TypeFalse Type = 6

	typeRawString Type = 7

	TypeNil Type = 8
)

const MaxDepth = 1000

var (
	valueTrue  = &Value{t: TypeTrue}
	valueFalse = &Value{t: TypeFalse}
	valueNull  = &Value{t: TypeNull}
)

var mParserPool ParserPool

func Parse(b []byte) (*Node, error) {
	p := mParserPool.Get() // &Parser{}
	v, err := p.Parse(b)
	return &Node{p: p, value: v, err: err}, err
}

func ParseFn(b []byte, fn func(*Node) error) error {
	p := mParserPool.Get()
	defer func() {
		mParserPool.Put(p)
	}()
	v, err := p.Parse(b)
	node := &Node{p: p, value: v, err: err}
	if fn != nil {
		err = fn(node)
	}
	return err
}

/*  Node implements external interface */

func (node *Node) Type() Type {
	if node == nil || node.value == nil {
		return TypeNil
	}
	if node.value.t == typeRawString {
		return TypeString
	}
	return node.value.t
}

func (node *Node) SetErr(err error) {
	node.err = err
}

func (node *Node) Err() error {
	return node.err
}

func (node *Node) ReleaseParseCache() {
	if node.p != nil {
		mParserPool.Put(node.p)
	}
}

func (node *Node) Get(key string) *Node {
	if node.value == nil || node.err != nil {
		return node
	}
	v, err := node.value.Get(key)
	return &Node{p: node.p, value: v, err: err}
}

func (node *Node) Array() ([]*Node, error) {
	if node.value == nil || node.err != nil {
		if node.err != nil {
			return nil, fmt.Errorf("node->Array() node err %s", node.err.Error())
		}
		return nil, fmt.Errorf("node->Array() value  is nil")
	}
	vs, err := node.value.Array()
	if err != nil {
		return nil, err
	}
	var nodes []*Node
	for _, v := range vs {
		nodes = append(nodes, &Node{p: node.p, value: v, err: nil})
	}
	return nodes, nil
}

func (node *Node) Bool() (bool, error) {
	if node.value == nil || node.err != nil {
		if node.err != nil {
			return false, fmt.Errorf("node->Bool() node err %s", node.err.Error())
		}
		return false, fmt.Errorf("node->Bool() value  is nil")
	}
	if node.value.t == TypeTrue {
		return true, nil
	} else if node.value.t == TypeFalse {
		return false, nil
	}
	return false, fmt.Errorf("node->Bool value Type err %s", node.value.t.String())
}

func (node *Node) Float64() (float64, error) {
	if err := node.checkNumber(); err != nil {
		return 0, fmt.Errorf("node->Float64 checkNumber fail %s", err.Error())
	}
	s, err := node.data()
	if err != nil {
		return 0, fmt.Errorf("node->Float64 data() fail %s", err.Error())
	}
	return fastfloat.Parse(s)
}

func (node *Node) Int() (int, error) {
	if err := node.checkNumber(); err != nil {
		return 0, fmt.Errorf("node->Int checkNumber fail %s", err.Error())
	}
	s, err := node.data()
	if err != nil {
		return 0, fmt.Errorf("node->Int data() fail %s", err.Error())
	}
	n, err := fastfloat.ParseInt64(s)
	if err != nil {
		return 0, err
	}

	nn := int(n)
	if int64(nn) != n {
		return 0, fmt.Errorf("number %q doesn't fit int", s)
	}
	return nn, nil
}

func (node *Node) Uint() (uint, error) {
	if err := node.checkNumber(); err != nil {
		return 0, fmt.Errorf("node->Uint checkNumber fail %s", err.Error())
	}
	s, err := node.data()
	if err != nil {
		return 0, fmt.Errorf("node->Uint data() fail %s", err.Error())
	}
	n, err := fastfloat.ParseUint64(s)
	if err != nil {
		return 0, err
	}

	nn := uint(n)
	if uint64(nn) != n {
		return 0, fmt.Errorf("number %q doesn't fit int", s)
	}
	return nn, nil
}

func (node *Node) Int64() (int64, error) {
	if err := node.checkNumber(); err != nil {
		return 0, fmt.Errorf("node->Int64 checkNumber fail %s", err.Error())
	}
	s, err := node.data()
	if err != nil {
		return 0, fmt.Errorf("node->Int64 data() fail %s", err.Error())
	}
	return fastfloat.ParseInt64(s)
}

func (node *Node) Uint64() (uint64, error) {
	if err := node.checkNumber(); err != nil {
		return 0, fmt.Errorf("node->Uint64 checkNumber fail %s", err.Error())
	}
	s, err := node.data()
	if err != nil {
		return 0, fmt.Errorf("node->Uint64 data() fail %s", err.Error())
	}
	return fastfloat.ParseUint64(s)
}

func (node *Node) String() (string, error) {
	return node.string()
}

func (node *Node) data() (string, error) {
	start, end, err := node.value.Data()
	if err != nil {
		return "", fmt.Errorf("node.value.Data ERR %s", err.Error())
	}
	return node.p.s[start:end], nil
}

func (node *Node) checkNumber() error {
	if node.value == nil || node.err != nil {
		if node.err != nil {
			return fmt.Errorf("node.err %s", node.err.Error())
		}
		return fmt.Errorf("node.value  is nil")
	}
	if node.value.t != TypeNumber {
		return fmt.Errorf("node.value Type err %s", node.value.t.String())
	}
	return nil
}

func (node *Node) parsenewstring(start, end int32, t Type) (string, error) {
	var newstart, newend int = -1, -1
	var cs, ce uint8

	cs = '{'
	ce = '}'

	if t == TypeArray {
		cs = '['
		ce = ']'
	} else if t == typeRawString || t == TypeString {
		cs = '"'
		ce = '"'
		if start == end && node.p.s[start] == '"' {
			return `""`, nil
		}
	}

	for i := int(start); i >= 0; i-- {
		if node.p.s[i] == cs {
			newstart = i
			break
		}
	}
	for i := int(end); i < len(node.p.s); i++ {
		if node.p.s[i] == ce {
			newend = i
			break
		}
	}
	if newstart < 0 || newend < 0 {
		return "", fmt.Errorf("node.string %s parse err", t.String())
	}
	return node.p.s[newstart : newend+1], nil
}

func (node *Node) string() (string, error) {
	if node.value == nil || node.err != nil {
		if node.err != nil {
			return "", fmt.Errorf("node err %s", node.err.Error())
		}
		return "", fmt.Errorf("node value  is nil")
	}

	switch node.value.t {
	case TypeTrue:
		return "true", nil
	case TypeFalse:
		return "false", nil
		//case TypeNull:
		//	return "null", nil
	}

	start, end, err := node.value.Data()
	if err != nil {
		return "", fmt.Errorf("node->String ERR %s", err.Error())
	}

	s := node.p.s[start:end]
	var dst []byte

	switch node.value.t {
	case TypeNull:
		return s, nil
	case TypeNumber:
		return s, nil
	case TypeString: //
		if !hasSpecialChars(s) {
			return node.parsenewstring(start, end, node.value.t)
		}
		dst = escapeString(dst, s)
		return b2s(dst), nil
	case TypeArray:
		return node.parsenewstring(start, end, node.value.t)
	case TypeObject:
		return node.parsenewstring(start, end, node.value.t)
	case typeRawString:
		return node.parsenewstring(start, end, node.value.t)
	}
	return s, nil
}

/***********   value methods *s*t*a*r*t*********/

func (v *Value) Get(key string) (*Value, error) {
	if v == nil {
		return nil, fmt.Errorf("value->Get value is nil")
	}
	if v.t == TypeObject {
		v = v.o.Get(key)
		if v == nil {
			return nil, fmt.Errorf("value not find by key %s", key)
		}
		return v, nil
	}
	return nil, fmt.Errorf("value type not corrent %s", v.t.String())
}

func (v *Value) Array() ([]*Value, error) {
	if v == nil {
		return nil, fmt.Errorf("value->Array value is nil")
	}
	if v.t == TypeArray {
		return v.a, nil
	}
	return nil, fmt.Errorf("value ->Array Type not corrent %s", v.t.String())
}

func (v *Value) Data() (int32, int32, error) {
	if v.spos < 0 || v.spos > v.epos {
		return -1, -1, fmt.Errorf("value->string except spos=%d,epos=%d", v.spos, v.epos)
	}
	return v.spos, v.epos, nil
}

//******************  value  method end **************************

func (p *Parser) Parse(b []byte) (*Value, error) {
	p.b = b
	p.s = b2s(b)
	p.c.reset()

	var soffset int = 0

	s := skipWS(p.s, &soffset)
	v, tail, err := parseValue(s, &p.c, 0, &soffset)
	if err != nil {
		return nil, fmt.Errorf("cannot parse JSON: %s; unparsed tail: %q", err, startEndString(tail))
	}
	tail = skipWS(tail, &soffset)
	if len(tail) > 0 {
		return nil, fmt.Errorf("unexpected tail: %q", startEndString(tail))
	}
	return v, nil
}

func (t Type) String() string {
	switch t {
	case TypeObject:
		return "object"
	case TypeArray:
		return "array"
	case TypeString:
		return "string"
	case TypeNumber:
		return "number"
	case TypeTrue:
		return "true"
	case TypeFalse:
		return "false"
	case TypeNull:
		return "null"

	// typeRawString is skipped intentionally,
	// since it shouldn't be visible to user.
	default:
		return fmt.Sprintf("BUG: unknown Value type: %d", t)
	}
}

func (c *cache) reset() {
	c.vs = c.vs[:0]
}

func (c *cache) getValue() *Value {
	if cap(c.vs) > len(c.vs) {
		c.vs = c.vs[:len(c.vs)+1]
	} else {
		c.vs = append(c.vs, Value{})
	}
	// Do not reset the value, since the caller must properly init it.
	return &c.vs[len(c.vs)-1]
}

func (o *Object) reset() {
	o.kvs = o.kvs[:0]
	o.keysUnescaped = false
}

func (o *Object) getKV() *kv {
	if cap(o.kvs) > len(o.kvs) {
		o.kvs = o.kvs[:len(o.kvs)+1]
	} else {
		o.kvs = append(o.kvs, kv{})
	}
	return &o.kvs[len(o.kvs)-1]
}

func (o *Object) Get(key string) *Value {
	if !o.keysUnescaped && strings.IndexByte(key, '\\') < 0 {
		// Fast path - try searching for the key without object keys unescaping.
		for _, kv := range o.kvs {
			if kv.k == key {
				return kv.v
			}
		}
	}
	// Slow path - unescape object keys.
	o.unescapeKeys()

	for _, kv := range o.kvs {
		if kv.k == key {
			return kv.v
		}
	}
	return nil
}

func (o *Object) unescapeKeys() {
	if o.keysUnescaped {
		return
	}
	kvs := o.kvs
	for i := range kvs {
		kv := &kvs[i]
		kv.k = unescapeStringBestEffort(kv.k)
	}
	o.keysUnescaped = true
}

/********************  local function  **************************/

func parseValue(s string, c *cache, depth int, soffset *int) (*Value, string, error) {
	if len(s) == 0 {
		return nil, s, fmt.Errorf("cannot parse empty string")
	}

	depth++
	if depth > MaxDepth {
		return nil, s, fmt.Errorf("too big depth for the nested JSON; it exceeds %d", MaxDepth)
	}

	if s[0] == '{' {
		*soffset++
		v, tail, err := parseObject(s[1:], c, depth, soffset)
		if err != nil {
			return nil, tail, fmt.Errorf("cannot parse object: %s", err)
		}
		return v, tail, nil
	}

	if s[0] == '[' {
		*soffset++
		v, tail, err := parseArray(s[1:], c, depth, soffset)
		if err != nil {
			return nil, tail, fmt.Errorf("cannot parse array: %s", err)
		}
		return v, tail, nil
	}

	if s[0] == '"' {
		*soffset++
		oldstr := s[1:]
		start, end, _, tail, err := parseRawString(s[1:])
		if err != nil {
			return nil, tail, fmt.Errorf("cannot parse string: %s", err)
		}
		v := c.getValue()
		v.t = typeRawString

		v.spos = int32(*soffset) + start
		v.epos = int32(*soffset) + end

		*soffset += len(oldstr) - len(tail)
		//if start >= 0 && end >= start {
		//	*soffset += int(end - start)
		//}
		return v, tail, nil
	}

	if s[0] == 't' {
		if len(s) < len("true") || s[:len("true")] != "true" {
			return nil, s, fmt.Errorf("unexpected value found: %q", s)
		}
		*soffset += 4
		return valueTrue, s[len("true"):], nil
	}

	if s[0] == 'f' {
		if len(s) < len("false") || s[:len("false")] != "false" {
			return nil, s, fmt.Errorf("unexpected value found: %q", s)
		}
		*soffset += 5
		return valueFalse, s[len("false"):], nil
	}

	if s[0] == 'n' || s[0] == 'N' {
		if len(s) < len("null") || strings.ToLower(s[:len("null")]) != "null" {
			// Try parsing NaN
			if len(s) >= 3 && strings.EqualFold(s[:3], "nan") {
				v := c.getValue()
				v.t = TypeNumber
				v.spos = int32(*soffset) + 0
				v.epos = int32(*soffset) + 3
				*soffset += 3
				return v, s[3:], nil
			}
			return nil, s, fmt.Errorf("unexpected value found: %q", s)
		}

		v := c.getValue()
		v.t = TypeNull
		v.spos = int32(*soffset) + 0
		v.epos = int32(*soffset) + 4
		*soffset += 4
		return v, s[4:], nil
		//return valueNull, s[len("null"):], nil
	}

	start, end, tail, err := parseRawNumber(s)
	if err != nil {
		return nil, tail, fmt.Errorf("cannot parse number: %s", err)
	}
	v := c.getValue()
	v.t = TypeNumber
	v.spos = int32(*soffset) + start
	v.epos = int32(*soffset) + end
	if start >= 0 && end >= start {
		*soffset += int(end - start)
	}
	return v, tail, nil
}

func parseArray(s string, c *cache, depth int, soffset *int) (*Value, string, error) {
	curoffset := *soffset
	s = skipWS(s, soffset)
	if len(s) == 0 {
		return nil, s, fmt.Errorf("missing ']'")
	}

	if s[0] == ']' {
		v := c.getValue()
		v.t = TypeArray
		v.spos = int32(curoffset)
		v.epos = int32(*soffset)
		v.a = v.a[:0]
		*soffset++
		return v, s[1:], nil
	}

	a := c.getValue()
	a.t = TypeArray
	a.spos = int32(*soffset)
	a.a = a.a[:0]
	for {
		var v *Value
		var err error

		s = skipWS(s, soffset)
		v, s, err = parseValue(s, c, depth, soffset)
		if err != nil {
			return nil, s, fmt.Errorf("cannot parse array value: %s", err)
		}
		a.a = append(a.a, v)

		s = skipWS(s, soffset)
		if len(s) == 0 {
			return nil, s, fmt.Errorf("unexpected end of array")
		}
		if s[0] == ',' {
			s = s[1:]
			*soffset++
			continue
		}
		if s[0] == ']' {
			a.epos = int32(*soffset)
			*soffset++
			s = s[1:]
			return a, s, nil
		}
		return nil, s, fmt.Errorf("missing ',' after array value")
	}
}

func parseObject(s string, c *cache, depth int, soffset *int) (*Value, string, error) {
	curoffset := *soffset
	s = skipWS(s, soffset)
	if len(s) == 0 {
		return nil, s, fmt.Errorf("missing '}'")
	}

	if s[0] == '}' {
		v := c.getValue()
		v.t = TypeObject
		v.o.reset()
		v.spos = int32(curoffset)
		v.epos = int32(*soffset)
		*soffset++
		return v, s[1:], nil
	}

	o := c.getValue()
	o.t = TypeObject
	o.o.reset()
	o.spos = int32(*soffset)
	for {
		var err error
		kv := o.o.getKV()

		// Parse key.
		s = skipWS(s, soffset)
		if len(s) == 0 || s[0] != '"' {
			return nil, s, fmt.Errorf(`cannot find opening '"" for object key`)
		}
		*soffset++
		kv.k, s, err = parseRawKey(s[1:], soffset)
		if err != nil {
			return nil, s, fmt.Errorf("cannot parse object key: %s", err)
		}
		s = skipWS(s, soffset)
		if len(s) == 0 || s[0] != ':' {
			return nil, s, fmt.Errorf("missing ':' after object key")
		}
		s = s[1:]
		*soffset++

		// Parse value
		s = skipWS(s, soffset)
		kv.v, s, err = parseValue(s, c, depth, soffset)
		if err != nil {
			return nil, s, fmt.Errorf("cannot parse object value: %s", err)
		}
		s = skipWS(s, soffset)
		if len(s) == 0 {
			return nil, s, fmt.Errorf("unexpected end of object")
		}
		if s[0] == ',' {
			s = s[1:]
			*soffset++
			continue
		}
		if s[0] == '}' {
			o.epos = int32(*soffset)
			*soffset++
			return o, s[1:], nil
		}
		return nil, s, fmt.Errorf("missing ',' after object value")
	}
}

func parseRawKey(s string, soffset *int) (string, string, error) {
	for i := 0; i < len(s); i++ {
		*soffset++
		if s[i] == '"' {
			// Fast path.
			return s[:i], s[i+1:], nil
		}
		if s[i] == '\\' {
			// Slow path.
			oldstr := s
			_, _, s, tail, err := parseRawString(s)
			if err == nil {
				*soffset += len(oldstr) - len(tail)
			}
			//if start >= 0 && end > start {
			//	*soffset += int(end - start)
			//}
			return s, tail, err
		}
	}
	return s, "", fmt.Errorf(`missing closing '"'`)
}

func parseRawNumber(s string) (int32, int32, string, error) {
	// The caller must ensure len(s) > 0

	// Find the end of the number.
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if (ch >= '0' && ch <= '9') || ch == '.' || ch == '-' || ch == 'e' || ch == 'E' || ch == '+' {
			continue
		}
		if i == 0 || i == 1 && (s[0] == '-' || s[0] == '+') {
			if len(s[i:]) >= 3 {
				xs := s[i : i+3]
				if strings.EqualFold(xs, "inf") || strings.EqualFold(xs, "nan") {
					return 0, int32(i + 3), s[i+3:], nil
				}
			}
			return -1, -1, s, fmt.Errorf("unexpected char: %q", s[:1])
		}

		s = s[i:]
		return 0, int32(i), s, nil
	}
	return 0, int32(len(s)), "", nil
}

func parseRawString(s string) (int32, int32, string, string, error) {
	n := strings.IndexByte(s, '"')
	if n < 0 {
		return -1, -1, s, "", fmt.Errorf(`missing closing '"'`)
	}
	if n == 0 || s[n-1] != '\\' {
		// Fast path. No escaped ".
		return 0, int32(n), s[:n], s[n+1:], nil
	}

	// Slow path - possible escaped " found.
	ss := s
	for {
		i := n - 1
		for i > 0 && s[i-1] == '\\' {
			i--
		}
		if uint(n-i)%2 == 0 {
			return 0, int32(len(ss) - len(s) + n), ss[:len(ss)-len(s)+n], s[n+1:], nil
		}
		s = s[n+1:]

		n = strings.IndexByte(s, '"')
		if n < 0 {
			return -1, -1, ss, "", fmt.Errorf(`missing closing '"'`)
		}
		if n == 0 || s[n-1] != '\\' {
			return 0, int32(len(ss) - len(s) + n), ss[:len(ss)-len(s)+n], s[n+1:], nil
		}
	}
}

func skipWS(s string, soffset *int) string {
	if len(s) == 0 || s[0] > 0x20 {
		// Fast path.
		return s
	}
	return skipWSSlow(s, soffset)
}

func unescapeStringBestEffort(s string) string {
	n := strings.IndexByte(s, '\\')
	if n < 0 {
		// Fast path - nothing to unescape.
		return s
	}

	// Slow path - unescape string.
	b := s2b(s) // It is safe to do, since s points to a byte slice in Parser.b.
	b = b[:n]
	s = s[n+1:]
	for len(s) > 0 {
		ch := s[0]
		s = s[1:]
		switch ch {
		case '"':
			b = append(b, '"')
		case '\\':
			b = append(b, '\\')
		case '/':
			b = append(b, '/')
		case 'b':
			b = append(b, '\b')
		case 'f':
			b = append(b, '\f')
		case 'n':
			b = append(b, '\n')
		case 'r':
			b = append(b, '\r')
		case 't':
			b = append(b, '\t')
		case 'u':
			if len(s) < 4 {
				// Too short escape sequence. Just store it unchanged.
				b = append(b, "\\u"...)
				break
			}
			xs := s[:4]
			x, err := strconv.ParseUint(xs, 16, 16)
			if err != nil {
				// Invalid escape sequence. Just store it unchanged.
				b = append(b, "\\u"...)
				break
			}
			s = s[4:]
			if !utf16.IsSurrogate(rune(x)) {
				b = append(b, string(rune(x))...)
				break
			}

			// Surrogate.
			// See https://en.wikipedia.org/wiki/Universal_Character_Set_characters#Surrogates
			if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
				b = append(b, "\\u"...)
				b = append(b, xs...)
				break
			}
			x1, err := strconv.ParseUint(s[2:6], 16, 16)
			if err != nil {
				b = append(b, "\\u"...)
				b = append(b, xs...)
				break
			}
			r := utf16.DecodeRune(rune(x), rune(x1))
			b = append(b, string(r)...)
			s = s[6:]
		default:
			// Unknown escape sequence. Just store it unchanged.
			b = append(b, '\\', ch)
		}
		n = strings.IndexByte(s, '\\')
		if n < 0 {
			b = append(b, s...)
			break
		}
		b = append(b, s[:n]...)
		s = s[n+1:]
	}
	return b2s(b)
}

func skipWSSlow(s string, soffset *int) string {
	if len(s) == 0 || s[0] != 0x20 && s[0] != 0x0A && s[0] != 0x09 && s[0] != 0x0D {
		return s
	}
	for i := 1; i < len(s); i++ {
		*soffset++
		if s[i] != 0x20 && s[i] != 0x0A && s[i] != 0x09 && s[i] != 0x0D {
			return s[i:]
		}
	}
	return ""
}

func hasSpecialChars(s string) bool {
	if strings.IndexByte(s, '"') >= 0 || strings.IndexByte(s, '\\') >= 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 0x20 {
			return true
		}
	}
	return false
}

func escapeString(dst []byte, s string) []byte {
	if !hasSpecialChars(s) {
		// Fast path - nothing to escape.
		dst = append(dst, '"')
		dst = append(dst, s...)
		dst = append(dst, '"')
		return dst
	}

	// Slow path.
	return strconv.AppendQuote(dst, s)
}
