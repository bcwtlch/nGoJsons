# nGoJsons
A fully integrated JSON serialization and deserialization library, especially the Get feature.

A new integrated open source library. Include two levels of function. 
- Marshal And Unmarshal interface compatible with go json library
- Get interface supported part to obtain json data


##  Features：
- The go sdk interface completely unifies the external calling method
- The Get interface unifies the external interface.

## Interface：
- go sdk Interface
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

- Get Interface
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

## Integrated Json open source library
- go sdk Interface Support
- - stdlib
- - go-json
- - json-iterator
- - sonic



- Get Interface Support
- - FastJson
- - json-iterator
- - sonic
- - SimpleJson
- - jsonparser
- - ngeyjson

Among them, ngeyjson is an open source json library with my own characteristics,
including ngeyjson-parse(Get interface) and support for Marshal and Unmarshal of go sdk library interfaces.
Because Marshal and Unmarshal are still immature, it is not open source for the time being.
Detail See About-ngeyjson for the design of ngeyjson.



## Usage

### Marshal/Unmarshal etc.
See the demo use case below for simple use. Please refer to the example  package for details.
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

### Get Usage
See the demo use case below for simple use. Please refer to the example  package for details
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



















