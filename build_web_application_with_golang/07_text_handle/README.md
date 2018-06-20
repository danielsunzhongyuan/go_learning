# 文本处理

### xml
import "encoding/xml"

func Unmarshal(data []byte, v interface{}) error // 把数据data解析到 v 中，一般来说，v是一个struct
func Marshal(v interface{}) ([]byte, error)
func Marshal(v interface{}, prefix, indent string) ([]byte, error)

这里要注意的是 struct tag ，也就是 `xml:"servers"`，要和 xml 文件里的tag对应（大小写敏感）。

### json

1. Unmarshal

`import "encoding/json"`

`func Unmarshal（data []byte, v interface{}) error`

这里的 v 可以是struct，也就是说在解析之前就已经知道了 json 的格式，这种方式和 xml 的解析类似。
v 也可以直接是一个 interface，也就是说对它的格式并不清楚，然后通过 type assert 的方式进一步解析。如下：
```go
var f interface{}
err := json.Unmarshal(b, &f)
m := f.(map[string]interface{}) // 这个地方就是 type assert
for k, v := range m {
    switch vv := v.(type) {
    case string:
        // print string
    case int:
        // int
    case float64:
        // float64
    case []interface{}:
        // slice
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        // unknown
    }
}
```

上面那种方式其实并不方便，尤其是在和 ES 一起使用时：
```
ES可以支持多重的聚合，即group by A，B，C。。。
那么返回的结果里，就会有多重的嵌套数组，而每个数组前都会有一个聚合"名称"，这个名称就是一开始查询的时候定义的。
所以在解析的时候，也要用到这些预先定义好的名称（不利于代码的重复利用）
```

推荐一个开源库： `https://github.com/bitly/go-simplejson`

2. Marshal
`func Marshal(v interface{})([]byte, error)`

结构体的方式：定义结构体，赋值，`json.Marshal(s)`。

3. struct tag 注意点
* tag是 - 时，这个字段不会输出
* tag带有 omitempty 时，如果该字段为空，就不输出到json里
* tag带有 ,string 时，如果字段类型是 bool，string，int，int64等，那么会转成json字符串

json字符串的意思是：
```
type Server struct {
    ...
    ServerName1 string `json:"serverName1"`
    ServerName2 string `json:"serverName2,string"`
    ...
}

s := Server {
    ID: 3,
    ServerName1: `Go "1.0"`,
    ServerName2: `Go "1.0"`,
    ...
}
b, _ := json.Marshal(s)
os.Stdout.Write(b)

// {"serverName1":"Go \"1.0\"","serverName2":"\"Go \\\"1.0\\\"\""}
```
以上就是 json字符串 和 字符串 的区别。


转换过程中：
* JSON对象只支持 string 作为 key，所以要编码一个 map，必须是 map[string]T 这种类型。
* Channel，complex和function是不能编码成json的
* 嵌套的数据不能编码，不然会让 json 陷入死循环
* 指针在编码时会输出指针指向的内容，而空指针会输出 null


### 正则处理

1. 是否匹配
```
func Match(pattern string, b []byte) (matched bool, err error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
func MatchString(pattern string, s string) (matched bool, err error)
```

2. 获取内容
```
func Compile(expr string) (*Regexp, error)
func CompilePOSIX(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
func MustCompilePOSIX(str string) *Regexp
```

3. 搜索
```
func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindAllString(s string, n int) []string
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
func (re *Regexp) FindString(s string) string
func (re *Regexp) FindStringIndex(s string) (loc []int)
func (re *Regexp) FindStringSubmatch(s string) []string
func (re *Regexp) FindStringSubmatchIndex(s string) []int
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindSubmatchIndex(b []byte) []int
```

输入源是 []byte, string, io.RuneReader，跟匹配那里是一样的。这样简化后就是：
```
func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindSubmatchIndex(b []byte) []int
```


### 模板 template
1. 获取数据
2. 获取模板
3. 渲染数据
```go
    func handler(w http.ResponseWriter, r *http.Request) {
        t := template.New("some template") //创建一个模板
        t, _ = t.ParseFiles("tmpl/welcome.html", nil)  //解析模板文件
        user := GetUser() //获取当前用户信息
        t.Execute(w, user)  //执行模板的merger操作
    }
```

自定义管道函数：
```
    var builtins = FuncMap{
        "and":      and,
        "call":     call,
        "html":     HTMLEscaper,
        "index":    index,
        "js":       JSEscaper,
        "len":      length,
        "not":      not,
        "or":       or,
        "print":    fmt.Sprint,
        "printf":   fmt.Sprintf,
        "println":  fmt.Sprintln,
        "urlquery": URLQueryEscaper,
    }
```


Must操作是用来监测模板是否正确的。

模板嵌套：
```
{{define "subordinate template name"}}内容{{end}}
```

正因为有模板嵌套，所以 template.ParseFiles 这里才是 files。


### 文件操作
目录操作的大多数函数都在 os 包里，例如：
```
func Mkdir(name string, perm FileMode) error // perm => 0777
func MkdirAll(path string, perm FileMode) error // 创建多级目录
func Remove(name string) error // 删除目录，如果里面有文件或目录，会出错
func RemoveAll(path string) error // 删目录及里面的东西
```

文件操作：
```
// 创建，打开
func Create(name string) (file *File, err Error) // 默认权限是0666
func NewFile(fd uintptr, name string) *File // 根据文件描述符创建文件
func Open(name string) (file *File, err Error) // 只读
func OpenFile(name string, flag int, perm uint32) (file *File, err Error) // flag表示 只读、读写

// 写文件
func (file *File) Write(b []byte) (n int, err Error)
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
func (file *File) WriteString(s string) (ret int, err Error)

// 删除
func Remove(name string) Error // 跟删除文件夹是一样的
```


### 字符串处理
```
import "strings"
func Contains(s, substr string) bool
func Join(a []string, sep string) string
func Index(s, sep string) int // 在s中查找sep的位置，返回位置值，找不到返回-1
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string // n表示替换的次数，小于0表示全部替换
func Split(s, sep string) []string
func Trim(s string, cutset string) string // 在s的头部和尾部去除cutset指定的字符串
func Fields(s string) []string // 去除s里的空格符，并按照空格分割返回slice
```

字符串转换
```
import "strconv"


```