package parse

import "testing"

func TestParseNull(t *testing.T) {
	jsonstr := `{
       "138586353": {
            "description": null,
            "id": 138586353,
            "logo": "/images/UE0AAAAACEKo8QAAAAZDSVRN",
            "name": "Pittsburgh Symphony Orchestra",
            "subTopicIds": [
                337184268,
                337184283,
                337184275
            ],
            "subjectCode": null,
            "subtitle": null,
            "topicIds": [
                324846099,
                107888604,
                324846100
            ]
        }}`
	t.Run("parase null with jsoniter", func(t *testing.T) {
		parse, err := Parse([]byte(jsonstr), SetParseFrame(JsonIterFrame))
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			ReleaseParseCache(parse)
		}()
		strnull, err := parse.Get("138586353").String() // .Get("description").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(strnull)

	})

	t.Run("parase null with sonic", func(t *testing.T) {
		parse, err := Parse([]byte(jsonstr), SetParseFrame(SonicFrame))
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			ReleaseParseCache(parse)
		}()
		strnull, err := parse.Get("138586353").Get("description").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(strnull)

	})

}

func TestParseFloat(t *testing.T) {
	//"inf_float": Inf,
	//"minus_inf_float": -Inf,
	//"nan": nan

	jsonstr := `	{ "zero_float1": 0.00,
		"zero_float2": -0e123

		 }`

	t.Run("parase float with jsoniter", func(t *testing.T) {
		parse, err := Parse([]byte(jsonstr), SetParseFrame(SonicFrame))
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			ReleaseParseCache(parse)
		}()
		strnull, err := parse.Get("zero_float2").Float64() // .Get("description").String()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(strnull)

	})

}
