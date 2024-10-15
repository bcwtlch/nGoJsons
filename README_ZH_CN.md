# nGoJsons
  一个全新集成的开源库。包括2个层面功能。
  - 兼容官方库的序列反序列化接口
  - 支持Get的部分获取json数据.

##  特点：
- 官方的接口完全统一了对外的调用方式。
- Get的功能统一了对外接口

## 接口：
- 官方接口
```go
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

HTMLEscape
Marshal
MarshalIndent
Indent
Valid
Unmarshal
NewDecoder
NewEncoder
```

- Get接口
```go
type IJsonParseRet interface {
	Get(key string) IJsonParseRet
	Array() ([]IJsonParseRet, error)

	Bool() (bool, error)

	String() (string, error)
	Float64() (float64, error)
	Int() (int, error)
	Uint() (uint, error)
	Int64() (int64, error)
	Uint64() (uint64, error)
}
```

## 集成的Json开源库
- 官方接口
- - stdlib
- - go-json
- - json-iterator
- - sonic



- Get接口
- - FastJson
- - json-iterator
- - sonic
- - SimpleJson
- - jsonparser
- - ngeyjson

其中ngeyjson是我独立开发的有自己特色的开源json库，包括 ngeyjson-parse（Get接口）以及对官方接口序列化和反序列化的
接口的支持。由于序列化和反序列化还不成熟，暂不开源。 关于ngeyjson的设计见 About ngeyjson.


## 用法

### 序列化和反序列化
简单使用见下面的demo用例：具体可参考example部分
```go
func testMarshal() {
	var s = struct {
		Name string
		Age  int
	}{
		"json",
		30,
	}

	fmt.Println("-----  Marshal Default Test -------")
	retbytes, err := nGoJsons.Marshal(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  Marshal StdlibJsonFrame Test -------")
	retbytes, err = nGoJsons.Marshal(&s, nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  Marshal GoJsonFrame Test -------")
	retbytes, err = nGoJsons.Marshal(&s, nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  Marshal SonicJsonFrame Test -------")
	retbytes, err = nGoJsons.Marshal(&s, nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  Marshal JsonIterJsonFrame Test -------")
	retbytes, err = nGoJsons.Marshal(&s, nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))
}
```

### Get用法
简单示例：具体见example部分
```go
func TestJson1(t *testing.T) {
	s := []byte(`{"name":{"first":"Janet","last":"Prichard"},"age":47}`)

	var fn = func(v ijsoner.IJsonParseRet) {
		defer func() { JsonParse.ReleaseCache(v) }()

		vstr, err := v.String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vstr)

		vstr, err = v.Get("name").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vstr)

		str, err := v.Get("name").Get("first").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(str)

		str, err = v.Get("name").Get("last").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(str)

		age, err := v.Get("age").Int()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(age)
	}

	t.Run("testDefaultParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testSimpleJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.SimpleJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testFastJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.FastJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testGeyJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.GeyJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testJsonIterParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.JsonIterFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testJsonParser", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.JsonParserFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testSonicJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse(s, JsonParse.SetParseFrame(JsonParse.SonicFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})
}
```
















