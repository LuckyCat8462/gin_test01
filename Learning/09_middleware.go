package Learning

//# 中间件 (middleWare)
//
//- 中间件，对以后的路由全部生效。
//  - 设置好中间件以后，所有的路由都会使用这个中间件。
//  - 设置以前的路由，不生效。

//什么是"中间件":

//- 早期：
//  - 用于 系统 和 应用之间。
//  - 中间件： 内核 —— 中间件 ——  用户应用
//
//- 现在：
//  - 用于 两个模块之间的 功能 软件(模块)
//  - 中间件：—— 承上启下。  前后台开发： 路由 ——> 中间件 (起过滤作用) ——> 控制器
//  - 特性：对 “中间件”指定位置 ， 以下的路由起作用！以上的，作用不到。

//- 中间件使用的 3 个知识：
//#### Next()
//
//- 表示，跳过当前中间件剩余内容， 去执行下一个中间件。 当所有操作执行完之后，以出栈的执行顺序返回，执行剩余代码。

//#### return
//
//- 终止执行当前中间件剩余内容，执行下一个中间件。 当所有的函数执行结束后，以出栈的顺序执行返回，但，不执行return后的代码！

//#### Abort()
//
//-  只执行当前中间件， 操作完成后，以出栈的顺序，依次返回上一级中间件。
