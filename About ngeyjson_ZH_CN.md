# ngeyjson Design

    ngeyjson是一个立足于无反射的Json库。

## 序列化和反序列化
  go sdk的 json 提供了对外接口。
```
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

```
 应用层可以实现 Marshaler 和 Unmarshaler , 达到无反射实现json的序列化和反序列化。而对于复杂的Json，
一个是应用层处理比较麻烦，一个是不具有通用性。可能这个项目序列反序列了特定结构，后续的项目可能还需要重新
实现。

  通过参考easyjson和ffjson的实现，我们将其进一步发展，完全可以实现这样一种方式：
  
让程序来运行时自动生成序列和反序列的代码。这样的好处是：
 
 1） 如果反序列的结构有变化，可以自适应，具有扩展性。
  
2） 不像 easyjson和ffjson那么麻烦，且具有侵入性。
  
3） 具有通用性，无论是自定义类型还是结构体类型都能完全适配.动态生成。
  
  这样的弊端是开始动态生成的代码比较耗时，但也仅仅是耗时一次，还有出问题可能调试跟踪定位比较麻烦。
  但随着这个框架能逐渐成熟的话，调试跟踪定位也会迎刃而解。


## 解析部分Json
  这部分发展和参考了 fastjson的实现。但是没有做到sonic那么极致，还是在fastjson基础上的进一步完善优化。
  后续有时间可能考虑再进一步的高效的实现。
  
  相比较于 fastjson, ngeyjson的区别在于
  - 1、结构不同

    fastjson的Value结构：
   ```
   type Value struct {
	o Object
	a []*Value
	s string
	t Type
   } 
   ```

   ngeyjson的Value结构：
   ```
   type Value struct {
	o    Object
	a    []*Value
	spos int32
	epos int32
	t    Type
	}   
   ```
- 2、对 null 的支持

  ngeyjson对null支持大小写，大写和小写混合同样支持.比如 Null


- 3、String的处理

  fastjson的String的处理是个败笔，作者在注释中也说仅仅让用在测试。
  ngeyjson对String的处理完全没有性能效率的问题.


- 4、转义符的处理
 
  因为fastjson的String处理的不好，导致key对应的有转义符的字符串类型的值
  并没有处理好。而ngeyjson完全支持。


- 5、后续迭代

  fastjson已经好久不更新了......











