package main

import (
	"bytes"
	"fmt"
	"github.com/bcwtlch/nGoJsons"
	"reflect"
)

/*
HTMLEscape
Marshal
MarshalIndent
Indent
Unmarshal
Valid
NewDecoder
NewEncoder
*/

func testHTMLEscape() {
	var b bytes.Buffer
	m := `{"M":"<html>foo &` + "\xe2\x80\xa8 \xe2\x80\xa9" + `</html>"}`
	//result: `{"M":"\u003chtml\u003efoo \u0026\u2028 \u2029\u003c/html\u003e"}`

	fmt.Println("-----  HTMLEscape default Test -------")
	nGoJsons.HTMLEscape(&b, []byte(m))
	fmt.Println(b.String())

	fmt.Println("-----  HTMLEscape GoJsonFrame Test -------")
	b.Reset()
	nGoJsons.HTMLEscape(&b, []byte(m), nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	fmt.Println(b.String())

	fmt.Println("-----  HTMLEscape StdlibJson Test -------")
	b.Reset()
	nGoJsons.HTMLEscape(&b, []byte(m), nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	fmt.Println(b.String())
}

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

func testMarshalIndent() {
	var s = struct {
		Name string
		Age  int
	}{
		"json",
		30,
	}

	fmt.Println("-----  MarshalIndent Default Test -------")
	retbytes, err := nGoJsons.MarshalIndent(&s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  MarshalIndent StdlibJsonFrame Test -------")
	retbytes, err = nGoJsons.MarshalIndent(&s, "", "\t", nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  MarshalIndent GoJsonFrame Test -------")
	retbytes, err = nGoJsons.MarshalIndent(&s, "", "\t", nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  MarshalIndent SonicJsonFrame Test -------")
	retbytes, err = nGoJsons.MarshalIndent(&s, "", "\t", nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))

	fmt.Println("-----  MarshalIndent JsonIterJsonFrame Test -------")
	retbytes, err = nGoJsons.MarshalIndent(&s, "", " ", nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(retbytes))
}

func testIndent() {
	var s = struct {
		Name string
		Age  int
	}{
		"json",
		30,
	}

	retbytes, err := nGoJsons.Marshal(&s)
	if err != nil {
		fmt.Println(err)
		return
	}

	var b bytes.Buffer
	fmt.Println("-----  Indent Default Test -------")

	err = nGoJsons.Indent(&b, retbytes, "=", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b.String()))
	b.Reset()

	fmt.Println("-----  Indent StdlibJsonFrame Test -------")
	err = nGoJsons.Indent(&b, retbytes, "=", "\t", nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b.String()))
	b.Reset()

	fmt.Println("-----  Indent GoJsonFrame Test -------")
	err = nGoJsons.Indent(&b, retbytes, "=", "\t", nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b.String()))
	b.Reset()
}

func testValid() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println("-----  Valid default Test -------")
	bret := nGoJsons.Valid([]byte(goodJSON))
	fmt.Println("goodJSON:", bret)

	bret = nGoJsons.Valid([]byte(badJSON))
	fmt.Println("badJSON:", bret)

	fmt.Println("-----  Valid StdlibJsonFrame Test -------")
	bret = nGoJsons.Valid([]byte(goodJSON), nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	fmt.Println("goodJSON:", bret)

	bret = nGoJsons.Valid([]byte(badJSON), nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	fmt.Println("badJSON:", bret)

	fmt.Println("-----  Valid GoJsonFrame Test -------")
	bret = nGoJsons.Valid([]byte(goodJSON), nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	fmt.Println("goodJSON:", bret)

	bret = nGoJsons.Valid([]byte(badJSON), nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	fmt.Println("badJSON:", bret)

	fmt.Println("-----  Valid SonicJsonFrame Test -------")
	bret = nGoJsons.Valid([]byte(goodJSON), nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame))
	fmt.Println("goodJSON:", bret)

	bret = nGoJsons.Valid([]byte(badJSON), nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame))
	fmt.Println("badJSON:", bret)

	fmt.Println("-----  Valid JsonIterJsonFrame Test -------")
	bret = nGoJsons.Valid([]byte(goodJSON), nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame))
	fmt.Println("goodJSON:", bret)

	bret = nGoJsons.Valid([]byte(badJSON), nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame))
	fmt.Println("badJSON:", bret)

}

func testUnmarshal() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal

	fmt.Println("-----  Unmarshal Default Test -------")
	err := nGoJsons.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Println("-----  Unmarshal StdlibJsonFrame Test -------")
	err = nGoJsons.Unmarshal(jsonBlob, &animals, nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Println("-----  Unmarshal GoJsonFrame Test -------")
	err = nGoJsons.Unmarshal(jsonBlob, &animals, nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Println("-----  Unmarshal SonicJsonFrame Test -------")
	err = nGoJsons.Unmarshal(jsonBlob, &animals, nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Println("-----  Unmarshal JsonIterJsonFrame Test -------")
	err = nGoJsons.Unmarshal(jsonBlob, &animals, nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
}

func testEncoder() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)

	fmt.Println("-----  Encoder default Test -------")
	if err := nGoJsons.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("w-json:  %s\n", w.Bytes())

	w.Reset()
	fmt.Println("-----  Encoder StdlibJsonFrame Test -------")
	if err := nGoJsons.NewEncoder(w, nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame)).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("w-json:  %s\n", w.Bytes())

	w.Reset()
	fmt.Println("-----  Encoder GoJsonFrame Test -------")
	if err := nGoJsons.NewEncoder(w, nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame)).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("w-json:  %s\n", w.Bytes())

	w.Reset()
	fmt.Println("-----  Encoder SonicJsonFrame Test -------")
	if err := nGoJsons.NewEncoder(w, nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame)).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("w-json:  %s\n", w.Bytes())

	w.Reset()
	fmt.Println("-----  Encoder JsonIterJsonFrame Test -------")
	if err := nGoJsons.NewEncoder(w, nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame)).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("w-json:  %s\n", w.Bytes())
}

func testDecoder() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))

	fmt.Println("-----  Decoder default Test -------")
	if err := nGoJsons.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	fmt.Println("-----  Decoder StdlibJsonFrame Test -------")
	r.Reset([]byte(`{"height":1.5,"age":10}`))
	if err := nGoJsons.NewDecoder(r, nGoJsons.SetJsonFrame(nGoJsons.StdlibJsonFrame)).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	fmt.Println("-----  Decoder GoJsonFrame Test -------")
	r.Reset([]byte(`{"height":1.5,"age":10}`))
	if err := nGoJsons.NewDecoder(r, nGoJsons.SetJsonFrame(nGoJsons.GoJsonFrame)).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	fmt.Println("-----  Decoder SonicJsonFrame Test -------")
	r.Reset([]byte(`{"height":1.5,"age":10}`))
	if err := nGoJsons.NewDecoder(r, nGoJsons.SetJsonFrame(nGoJsons.SonicJsonFrame)).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	fmt.Println("-----  Decoder JsonIterJsonFrame Test -------")
	r.Reset([]byte(`{"height":1.5,"age":10}`))
	if err := nGoJsons.NewDecoder(r, nGoJsons.SetJsonFrame(nGoJsons.JsonIterJsonFrame)).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
}

func main() {
	//1.HTMLEscape
	//testHTMLEscape()
	//2.Marshal
	//testMarshal()

	//testMarshalIndent()
	//testIndent()
	//testValid()

	//testUnmarshal()
	//testEncoder()
	testDecoder()

}
