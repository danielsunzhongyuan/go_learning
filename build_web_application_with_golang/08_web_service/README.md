# Web服务

目前主流的Web服务有：REST、SOAP。REST就不说了，基于HTTP协议的一个补充，
>SOAP（Simple Object Access Protocol ）简单对象访问协议是在分散或分布式的环境中交换信息的简单的协议，是一个基于XML的协议，它包括四个部分：
1. SOAP封装(envelop)，封装定义了一个描述消息中的内容是什么，是谁发送的，谁应当接受并处理它以及如何处理它们的框架；
2. SOAP编码规则（encoding rules），用于表示应用程序需要使用的数据类型的实例;
3. SOAP RPC表示(RPC representation)，表示远程过程调用和应答的协定;
4. SOAP绑定（binding），使用底层协议交换信息。


>某知乎回答：
简单来说：SOAP = HTTP + XML + RPC
而 RPC 是可以走 HTTP、TCP 等不同等协议的。

但是 SOAP 的标准规范太复杂，这里不会介绍SOAP而是会介绍：
1. Socket编程
2. Websocket
3. REST
4. RPC

很多游戏服务采用 Socket 来编写服务端，而很多页游公司采用 Websocket 进行开发。


# Socket 编程
Socket起源于Unix的哲学："一切皆文件"，打开->读写->关闭。
Socket数据传输是一种特殊的I/O，它本身也是一种文件描述符。

Socket的打开是 `Socket()`


Socket分类：
1. 流式Socket（SOCK_STREAM），面向连接，针对于TCP服务应用；
2. 数据报式Socket（SOCK_DGRAM），无连接，对应于 UDP 服务应用。


Socket的唯一标识：ip地址 + 协议 + 端口。

> parallel_programming_in_practice 里的第三章，也介绍了一些 Socket 编程。

### TCP Socket

```
type TCPAddr struct {
    IP IP
    Port int
    Zone string // IPv6 scoped addressing zone
}

// net => tcp4(IPv4 only),tcp6(IPv6 only),tcp(两个都行)
// addr 表示域名或IP地址，例如 "www.google.com:80" / "127.0.0.1:22"
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)

func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error)


// TCP client
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)


// TCP server
func ListenTCP(net string, laddr *TCPAddr) (I *TCPListener, err os.Error)
func (I *TCPListener) Accept() (c Conn, err os.Error)
```

其他一些控制TCP连接的函数：
```
func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)

func (c *TCPConn) SetReadDeadline(t time.Time) error
func (c *TCPConn) SetWriteDeadline(t time.Time) error

func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```


### UDP Socket
>Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,
UDP缺少了对客户端连接请求的Accept函数。
其他基本几乎一模一样，只有TCP换成了UDP而已。
UDP的几个主要函数如下所示：
```
func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error)
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
```


# WebSocket
>WebSocket是HTML5的重要特性，它实现了基于浏览器的远程socket，
它使浏览器和服务器可以进行全双工通信，
许多浏览器（Firefox、Google Chrome和Safari）都已对此做了支持。
它解决了Web实时化的问题，相比传统HTTP有如下好处：
* 一个Web客户端只建立一个TCP连接
* Websocket服务端可以推送(push)数据到web客户端.
* 有更加轻量级的头，减少数据传送量


### WebSocket原理
>WebSocket的协议颇为简单，在第一次handshake通过以后，连接便建立成功，
其后的通讯数据都是以”\x00″开头，以”\xFF”结尾。
在客户端，这个是透明的，WebSocket组件会自动将原始数据“掐头去尾”。
浏览器发出WebSocket连接请求，然后服务器发出回应，然后连接建立成功，这个过程通常称为“握手” (handshaking)。


请求中的"Sec-WebSocket-Key"是随机字符串，加上一个固定的字符串：`258EAFA5-E914-47DA-95CA-C5AB0DC85B11`，
然后sha1安全散列算法计算出二进制的值，再用base64编码，即可以得到握手后的字符串：`rE91AJhfC+6JdVcVXOGJEADEJdQ=`，
作为响应头 `Sec-WebSocket-Accept` 的值反馈给客户端。

客户端一共绑定了四个事件：
1. onopen 建立连接后触发
2. onmessage 收到消息后触发
3. onerror 发生错误时触发
4. onclose 关闭连接时触发

>通过上面的例子我们看到客户端和服务器端实现WebSocket非常的方便，
Go的源码net分支中已经实现了这个的协议，我们可以直接拿来用。
目前随着HTML5的发展，我想未来WebSocket会是Web开发的一个重点，我们需要储备这方面的知识。


# REST

### 总结
>REST是一种架构风格，汲取了WWW的成功经验：
无状态，以资源为中心，充分利用HTTP协议和URI协议，提供统一的接口定义，使得它作为一种设计Web服务的方法而变得流行。
在某种意义上，通过强调URI和HTTP等早期Internet标准，REST是对大型应用程序服务器时代之前的Web方式的回归。
目前Go对于REST的支持还是很简单的，通过实现自定义的路由规则，
我们就可以为不同的method实现不同的handle，这样就实现了REST的架构。


# RPC
>很多独立的应用并没有采用"约定好信息格式的信息交换"的这种模式，而是采用类似常规的函数调用的方式来完成想要的功能。

>RPC就是想实现函数调用模式的网络化。客户端就像调用本地函数一样，
然后客户端把这些参数打包之后通过网络传递到服务端，服务端解包到处理过程中执行，然后执行的结果反馈给客户端。

>RPC（Remote Procedure Call Protocol）——远程过程调用协议，
是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。
它假定某些传输协议的存在，如TCP或UDP，以便为通信程序之间携带信息数据。
通过它可以使函数调用模式网络化。在OSI网络通信模型中，RPC跨越了传输层和应用层。
RPC使得开发包括网络分布式多程序在内的应用程序更加容易。

### RPC工作原理
运行时,一次客户机对服务器的RPC调用,其内部操作大致有如下十步：
1. 调用客户端句柄；执行传送参数
2. 调用本地系统内核发送网络消息
3. 消息传送到远程主机
4. 服务器句柄得到消息并取得参数
5. 执行远程过程
6. 执行的过程将结果返回服务器句柄
7. 服务器句柄返回结果，调用远程系统内核
8. 消息传回本地主机
9. 客户句柄由内核接收消息
10. 客户接收句柄返回的数据

### Go RPC
Go标准包中提供了对RPC的支持，3个级别：TCP、HTTP、JSONRPC。

Go的RPC包和传统的RPC不一样，只支持Go开发的服务器与客户端之间的交互，因为在内部，他们采用了Gob来编码。
Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
* 函数必须是导出的(首字母大写)
* 必须有两个导出类型的参数，
* 第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
* 函数还要有一个返回值error

举个例子，正确的RPC函数格式如下：

```
func (t *T) MethodName(argType T1, replyType *T2) error
```
T、T1和T2类型必须能被encoding/gob包编解码。

任何的RPC都需要通过网络来传递数据，Go RPC可以利用HTTP和TCP来传递数据，
利用HTTP的好处是可以直接复用net/http里面的一些函数。详细的例子请看下面的实现

### HTTP RPC


# 小结
这一章我们介绍了目前流行的几种主要的网络应用开发方式：

第一小节介绍了网络编程中的基础:Socket编程，因为现在网络正在朝云的方向快速进化，
作为这一技术演进的基石的的socket知识，作为开发者的你，是必须要掌握的。

第二小节介绍了正愈发流行的HTML5中一个重要的特性WebSocket，
通过它服务器可以实现主动的push消息，以简化以前ajax轮询的模式。

第三小节介绍了REST编写模式，这种模式特别适合来开发网络应用API，
目前移动应用的快速发展，我觉得将来会是一个潮流。

第四小节介绍了Go实现的RPC相关知识，对于上面四种开发方式，
Go都已经提供了良好的支持，net包及其子包，是所有涉及到网络编程的工具的所在地。
如果你想更加深入的了解相关实现细节，可以尝试阅读这个包下面的源码。