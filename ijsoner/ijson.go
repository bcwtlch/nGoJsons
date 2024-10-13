package ijsoner

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
