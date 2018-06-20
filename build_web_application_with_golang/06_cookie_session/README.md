# cookie and session

目的是为了识别用户的身份，而不需要访问每一个页面的时候都要输入一次用户名和密码。

cookie是一种客户端机制，而session是服务端机制。session可以存储在内存里，数据库里（memcache/redis）。

以下会着重介绍session。
### cookie
1. cookie是本地计算机保存一些用户操作的历史信息（包括登录信息），
并在用户再次访问该站点时浏览器通过HTTP协议将本地cookie内容发送给服务器，从而完成验证。
2. cookie一般是在第一次请求页面的时候，从服务器端获取，然后以后的请求会自动带上。
3. cookie是由浏览器维持的，存储在客户端的一小段文本信息，伴随着用户请求和页面在Web服务器和浏览器之间传递。
（打开浏览器里的cookie设置，即可查看已访问过的网站的cookie）
4. cookie有时间限制，根据生命周期分为：会话cookie，持久cookie。
    * 不设置过期时间就是会话cookie，保存于内存，关闭浏览器就没了；
    * 设置了过期时间（setMaxAge(606024)），浏览器就会把cookie存到硬盘上，关闭后再打开浏览器，cookie依然有效（除非过期了）
    * 存储在硬盘上的cookie可以在不同的浏览器进程间共享，而保存在内存的cookie，不同的浏览器处理方式就不一样了。


### session
1. session是服务器上保存用户操作的历史信息。服务器用 session id 来标识session。session id是由服务器产生，保证随机、唯一。
2. 但该方式下，仍然需要将发送请求的客户端与session进行对应，所以可以借助cookie机制来获取客户端的标识（即session id），
也可以通过GET方式将 id 提交给服务器。
3. session一般翻译为会话，与网络协议相关联时，又往往隐含了"面向连接"和/或"保持状态"这两个含义。
4. 服务器端一般采用类似散列表的结构来保存信息
5. 当程序要为某个客户端的请求创建一个session时，会先检查请求里是否带上了 session id，
    * 有的话，校验合法性，
    * 没有的话，或者不合法，就会创建一个新的session，然后把session id返回给客户端保存。

session机制并不复杂，但是实现和配置上的灵活性却使得具体情况复杂多变。

### 小结
session和cookie的目的相同，都是要解决http协议无状态的缺陷。
session通过cookie，在客户端保存session id，而将用户的其他会话信息保存在服务端的session对象里。
与此相对的，cookie要将所有信息都保存在客户端。因此cookie存在一定的安全隐患。

例如本地cookie中保存的用户名和密码被破译，或cookie被其他网站收集。


### session 创建过程
1. generate session id
2. store session info in memory or in database
3. send session id back to client

最关键的是如何发送session id这一步。
考虑到HTTP协议的定义，数据无非可以放到 `请求行 request line`、`头域 header`、`body`里，所以一般会有两种常用的方式：cookie和URL重写。
1. Cookie。服务端通过设置 set-cookie 头就可以将session id 传送给客户端，此后，客户端每次请求都会带上这个id。
另外，一般包含session信息的cookie会将失效时间设置为0（会话cookie），即浏览器进程有效时间。
2. URL重写。就是返回给用户的页面里所有的URL后面追加session id，这样用户在收到响应后，无论惦记响应页面里的哪一个链接或提交表单，都会自动带上session id，
从而实现了会话的保持。这种方式比较麻烦，但是如果客户端禁用了cookie的话，就首选这种方式。


