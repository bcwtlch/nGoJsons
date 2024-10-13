package pSimpleJson

import "testing"

func TestString1(t *testing.T) {
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
}

func TestString2(t *testing.T) {
	str := `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`

	p, err := Parse([]byte(str))
	if err != nil {
		t.Fatal(err)
	}

	vstr, err := p.Get("Name").String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vstr)
}
