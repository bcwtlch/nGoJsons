package parse

import (
	"fmt"
	"testing"
)

func TestValueInvalidConversion(t *testing.T) {

	v, err := Parse([]byte(`[{},[],"",123.45,true,null]`))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

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
		fmt.Println(str)
	}
}

func TestValueFloat64(t *testing.T) {

	v, err := Parse([]byte(`	{ "zero_float1": 1.3456,
		"zero_float2": -0e123,
		"inf_float": Inf,
		"minus_inf_float": -Inf,
		"nan": nan }`))

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	f, err := v.Get("zero_float1").Float64()
	if err != nil {
		t.Fatalf("str error: %s", err)
	}
	fmt.Println(f)

	f, err = v.Get("zero_float2").Float64()
	if err != nil {
		t.Fatalf("str error: %s", err)
	}
	fmt.Println(f)

	f, err = v.Get("inf_float").Float64()
	if err != nil {
		t.Fatalf("str error: %s", err)
	}
	fmt.Println(f)

	f, err = v.Get("minus_inf_float").Float64()
	if err != nil {
		t.Fatalf("str error: %s", err)
	}
	fmt.Println(f)

	f, err = v.Get("nan").Float64()
	if err != nil {
		t.Fatalf("str error: %s", err)
	}
	fmt.Println(f)
}

func TestValueString(t *testing.T) {

	t.Run("teststring1", func(t *testing.T) {
		str := `{"id":10,"orderNum":"100200300","money":99.99,"payTime":"2021-12-28T23:44:36.258311+08:00","extend":{"name":"zhangsan"}}`
		p, err := Parse([]byte(str))
		if err != nil {
			t.Fatal(err)
		}

		vstr, err := p.Get("extend").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vstr)
	})

	t.Run("teststring2", func(t *testing.T) {
		str := `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`
		p, err := Parse([]byte(str))
		if err != nil {
			t.Fatal(err)
		}
		vstr, err := p.Get("Colors").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vstr)
	})

	t.Run("teststring3", func(t *testing.T) {
		str := `{"foo":[{"bar":{"baz":123,"x":"434"},"y":[]},[null, false]],"qwe":false}`
		p, err := Parse([]byte(str))
		if err != nil {
			t.Fatal(err)
		}
		vstr, err := p.Get("foo").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(vstr)

	})

}

func TestValueBool(t *testing.T) {
	s := `{"foo":[{"bar":{"baz":123,"x":"434"},"y":[]},[null, false]],"qwe":false}`
	v, err := Parse([]byte(s))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	b, err := v.Get("qwe").Bool()
	if err != nil {
		t.Fatalf("qwe bool fail %s", err)
	}
	t.Log(b)

	arr, err := v.Get("foo").Array()
	if err != nil {
		t.Fatalf("foo arr fail %s", err)
	}

	for i, v := range arr {
		if i == 0 {
			m, err := v.Get("bar").Get("baz").Int()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(m)
			xstr, err := v.Get("bar").Get("x").String()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(xstr)

			ystr, err := v.Get("y").String()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(ystr)
		} else if i == 1 { //
			if v.TypeOf() == TypeArray {
				v1, err := v.Array()
				if err != nil {
					t.Fatal(err)
				}
				for _, v2 := range v1 {
					if v2.TypeOf() == TypeFalse || v2.TypeOf() == TypeTrue {
						b, _ := v2.Bool()
						t.Log(b)
					} else if v2.TypeOf() == TypeNumber {
						f, _ := v2.Float64()
						t.Log(f)
					} else {
						str, _ := v2.String()
						t.Log(str)
					}
				}
			} else {
				str, err := v.String()
				if err != nil {
					t.Fatal(err)
				}
				t.Log(str)
			}
			//v1, err := v.Array()
			//if err != nil {
			//	t.Fatal(err)
			//}
			//for _, v2 := range v1 {
			//	//str, err := v2.String()
			//	//if err != nil {
			//	//	t.Fatal(err)
			//	//}
			//	//t.Log(str)
			//
			//} //for _, v2 := range v1 {
		}

	}
}
