# ngeyjson Design

    ngeyjson is a Json library based on no reflection

## Marshal/Unmarshal
Json of go sdk provides an application layer interface.。
```
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

```
The application layer can realize Marshaler and Unmarshaler of json without reflection.
For complex Json, one is that the application layer processing is more troublesome, and the other is that it is not universal.
Perhaps this project realize marshal/unmarshal of the specific structure, and subsequent projects may need to be re-implemented.

By referring to the implementation of easyjson and ffjson,
We can further develop it and achieve such a way：

Enable the program to automatically generate marshal/unmarshal code during runtime. The benefits of this are:

1）If the structure changes, it can adapt and have scalability.

2）Not as invasive and troublesome as easyjson and ffjson.

3）It has universality and can fully adapt to both custom types and struct types. It is dynamically generated.

The downside of this is that the code generated dynamically at the beginning is relatively time-consuming, 
but it only takes time once, and there is also the possibility of debugging, tracking, and locating problems, which can be troublesome.
But as this framework gradually matures, debugging, tracking, and positioning will also be easily solved.

## Parse parts Json
This section has developed and referenced the implementation of fastjson. 
But it didn't achieve the extreme of Sonic, so it was further improved and optimized on the basis of FastJSON.
There may be time to consider further efficient implementation in the future.

The difference between ngeyjson and fastjson is 
- 1、Different structures

  fastjson - Value Struct：
   ```
   type Value struct {
	o Object
	a []*Value
	s string
	t Type
   } 
   ```

ngeyjson - Value Struct：
   ```
   type Value struct {
	o    Object
	a    []*Value
	spos int32
	epos int32
	t    Type
	}   
   ```
- 2、Support for null

  ngeyjson supports both uppercase and lowercase mixed case for null, such as Null.


- 3、Support for String()

  The handling of String() in fastjson is a flaw, and the author also stated in the comments that it is only used for testing purposes.
  The processing of String() by ngeyjson has no performance efficiency issue at all

- 4、Handling of escape characters

Due to poor handling of String() in fastjson, the value of the string type with escape characters corresponding to the key is affected
Not handled properly. And ngeyjson fully supports it.


- 5、Subsequent iterations

  Fastjson hasn't been updated for a long time......











