# Web Server 2
---
##### Requirement: implement a simple web server, it have
- 支持静态文件服务
- 支持简单 js 访问
- 提交表单，并输出一个表格
- 对 `/unknown` 给出开发中的提示，返回码 `5xx`

##### Develop Tool: Atom 1.21.0
---
这次主要是利用几个库——negroni、render、template、mux、router等完成一个能够和用户进行简单交互的服务器。

---
#### 效果展示

1. 支持静态文件访问
![静态文件访问][picture/静态文件访问.png]
