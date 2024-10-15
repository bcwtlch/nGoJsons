package main

import (
	"github.com/bcwtlch/nGoJsons/ijsoner"
	JsonParse "github.com/bcwtlch/nGoJsons/parse"
	"testing"
)

func TestJson1(t *testing.T) {
	s := []byte(`{"name":{"first":"Janet","last":"Prichard"},"age":47}`)

	var fn = func(v ijsoner.IJsonParseRet) {
		defer func() { JsonParse.ReleaseCache(v) }()

		vstr, err := v.String()
		if err != nil {
			t.Fatalf("v.String error: %s", err)
		}
		t.Log(vstr)

		vstr, err = v.Get("name").String()
		if err != nil {
			t.Fatalf("v.Get(\"foo\").String error: %s", err)
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

func TestJsonSpecial1(t *testing.T) {
	//s := []byte(`{"foo": [123, "      \"b\nar\""]}`)
	//s := []byte(`{"foo": [123, "      \"c:\program files\ceshibar\\""]}`)
	s := []byte(`{"foo": [123, "      \"c:\\program files\\ceshibar\\\""]}`)

	var fn = func(v ijsoner.IJsonParseRet) {
		defer func() { JsonParse.ReleaseCache(v) }()

		vstr, err := v.String()
		if err != nil {
			t.Fatalf("v.String error: %s", err)
		}
		t.Log(vstr)

		vstr, err = v.Get("foo").String()
		if err != nil {
			t.Fatalf("v.Get(\"foo\").String error: %s", err)
		}
		t.Log(vstr)

		arr, err := v.Get("foo").Array()
		if err != nil {
			t.Fatal(err)
		}

		for i, v := range arr {
			if i == 0 {
				m, err := v.Int()
				if err != nil {
					t.Fatal(err)
				}
				t.Log(m)
			} else {
				str, err := v.String()
				if err != nil {
					t.Fatal(err)
				}
				t.Log(str)
			}
		}
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

func TestArray1(t *testing.T) {

	var fn = func(v ijsoner.IJsonParseRet) {
		defer func() { JsonParse.ReleaseCache(v) }()
		vstr, err := v.String()
		if err != nil {
			t.Fatalf("v.String error: %s", err)
		}

		t.Log(vstr)

		arr, err := v.Array()
		if err != nil {
			t.Fatalf("Array error: %s", err)
		}

		for _, v1 := range arr {
			str, err := v1.String()
			if err != nil {
				t.Fatalf("str error: %s", err)
			}
			t.Log(str)
		}
	}

	t.Run("testDefaultParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testSimpleJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`), JsonParse.SetParseFrame(JsonParse.SimpleJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testFastJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`), JsonParse.SetParseFrame(JsonParse.FastJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testGeyJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,Null]`), JsonParse.SetParseFrame(JsonParse.GeyJsonFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testJsonIterJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`), JsonParse.SetParseFrame(JsonParse.JsonIterFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testJsonParser", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`), JsonParse.SetParseFrame(JsonParse.JsonParserFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

	t.Run("testSonicJsonParse", func(t *testing.T) {
		v, err := JsonParse.Parse([]byte(`[{},[],"",123.45,true,null]`), JsonParse.SetParseFrame(JsonParse.SonicFrame))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fn(v)
	})

}
