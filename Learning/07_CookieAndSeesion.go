package Learning

//- http协议，有 3 个版本：
//  - http/1.0 版：无状态，短连接。
//  - http/1.1 版：可以记录状态。—— 默认支持。
//  - http/2.0 版：可以支持长连接。 协议头：Connection: keep-alive 。
//
//###
//### Cookie
//
//- 最早的 http/1.0 版，提供 Cookie 机制， 但是没有 Session。
//- Cookie 作用：一定时间内， 存储用户的连接信息。如：用户名、登录时间 ... 不敏感信息。
//- Cookie 出身：http自带机制。Session不是！
//- Cookie 存储：Cookie 存储在 客户端 (浏览器) 中。—— 浏览器可以存储数据。少
//  - 存储形式：key - value
//  - 可以在浏览器中查看。
//  - Cookie 不安全。直接将数据存储在浏览器上。

////### Session
////
////- ”会话“：在一次会话交流中，产生的数据。不是http、浏览器自带。
////- Session 作用：一定时间内， 存储用户的连接信息。
////- Session 存储：在服务器中。一般为 临时 Session。—— 会话结束 (浏览器关闭) ， Session被干掉！

//### 对比 Cookie 和 Session
//
//1.  Cookie 存储在 浏览器， 在哪生成呢？
//		在web服务生成
//2.  Session 存储在 服务器，在哪生成呢？
//		session是以cookie加密为key,在web服务生成的
//3.  什么时候生成Cookie ， 什么时候生成 Session？

//	a.浏览器发送请求-不携带数据,到web服务
//	b.web服务产生cookie,携带cookie返回到浏览器;浏览器存储cookie
//	c.cookie加密作为key,生产session作为value,存入容器中
//	d.浏览器携带上次的cookie发送到web服务;web服务以cookie加密为key查session

//###  Session操作
//- gin 框架， 默认不支持Session功能。要想在 gin 中使用 Session，需要添加插件！
//
//- gin 框架中的 “插件”  —— 中间件 —— gin MiddleWare
//
//- 去 github 搜索，gin Session 可以得到：https://github.com/gin-contrib/sessions
//
//- 安装 Session 插件。
//go get github.com/gin-contrib/sessions

//context.SetCookie()参数
//name 名称;
//value 值;
//maxAge：最大生命周期。
//	 = 0 : 表示没指定该属性。
//	 < 0 ：表示删除。 ---- // 删除Cookie 的操作， 可以使用 该属性实现。
//	 > 0 ：指定生命周期。 单位：s
//path：路径。—— 通常传""
//domain：域名。 IP地址。
//secure：设置是否安全保护。true：不能在 地址栏前，点击查看。 可以使用 F12 查看。
//false：能在 地址栏前，点击查看。
//httpOnly：是否只针对http协议。
