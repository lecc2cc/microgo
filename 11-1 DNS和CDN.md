### DNS、CDN和多活架构



#### DNS和CDN



DNS（Domain Name System，域名系统），DNS 服务用于在网络请求时，将域名转为 IP 地址。能够使用户更方便的访问互联网，而不用去记住能够被机器直接读取的 IP 数串。

传统的基于 UDP 协议的公共 DNS 服务极易发生 DNS 劫持，从而造成安全问题。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-2021-06-28-23.png?raw=true)

#### 查询方式

**递归查询**

如果主机所询问的本地域名服务器不知道被查询域名的 *IP* 地址，那么本地域名服务器就以 *DNS* 客户的身份，向其他根域名服务器继续发出查询请求报文，而不是让该主机自己进行下一步的查询。

**迭代查询**

当根域名服务器收到本地域名服务器发出的迭代查询请求报文时，要么给出所要查询的 *IP* 地址，要么告诉本地域名服务器：你下一步应当向哪一个域名服务器进行查询。然后让本地域名服务器进行后续的查询，而不是替本地域名服务器进行后续的查询。

客户端到 *Local DNS* 服务器，*Local DNS* 与上级 *DNS* 服务器之间属于递归查询；*DNS* 服务器与根 *DNS* 服务器之前属于迭代查询。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-dns-parser-2021-06-28-23.png?raw=true)

```
1.用户在 Web 浏览器中键入 “example.com”，查询传输到 Internet 中，并被 DNS 递归解析器接收。

2.接着，解析器查询 DNS 根域名服务器（.）。

3.然后，根服务器使用存储其域信息的顶级域（TLD）DNS 服务器（例如 .com 或 .net）的地址响应该解析器。在搜索 example.com 时，我们的请求指向 .com TLD。

4.然后，解析器向 .com TLD 发出请求。

5.TLD 服务器随后使用该域的域名服务器 example.com 的 IP 地址进行响应。

6.最后，递归解析器将查询发送到域的域名服务器。

7.example.com 的 IP 地址而后从域名服务器返回解析器。

8.然后 DNS 解析器使用最初请求的域的 IP 地址响应 Web 浏览器。
```

#### DNS 问题

**Local DNS 劫持**

Local DNS 把域名劫持到其他域名，实现其不可告人的目的。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-dns-6-2021-06-28-23.png?raw=true)

**域名缓存**

域名缓存就是 LocalDNS 缓存了业务的域名的解析结果，不向权威 DNS 发起递归。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-dns7-2021-06-28-23.png?raw=true)

+ 保证用户访问流量在本网内消化：国内的各互联网接入运营商的带宽资源、网间结算费用、*IDC* 机房分布、网内 *ICP* 资源分布等存在较大差异。为了保证网内用户的访问质量，同时减少跨网结算，运营商在网内搭建了内容缓存服务器，通过把域名强行指向内容缓存服务器的 *IP* 地址，就实现了把本地本网流量完全留在了本地的目的。
+ 推送广告：有部分 *LocalDNS* 会把部分域名解析结果的所指向的内容缓存，并替换成第三方广告联盟的广告。

**域名解析转发**

解析转发是指运营商自身不进行域名递归解析，而是把域名解析请求转发到其它运营商的递归 *DNS* 上的行为。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-dns8-2021-06-28-23.png?raw=true)

部分小运营商为了节省资源，就直接将解析请求转发到了其它运营的递归 LocalDNS 上去了。

这样的直接后果就是权威 *DNS* 收到的域名解析请求的来源 *IP* 就成了其它运营商的 *IP*，最终导致用户流量被导向了错误的 *IDC*，用户访问变慢。

![image](https://github.com/lecc2cc/microgo/blob/master/images/11-1-dns9-2021-06-28-23.png?raw=true)

LocalDNS 递归出口 NAT (`网络地址转换协议`)指的是运营商的 LocalDNS 按照标准的 DNS 协议进行递归，但是因为在网络上存在多出口且配置了目标路由 NAT，结果导致 LocalDNS 最终进行递归解析的时候的出口 IP 就有概率不为本网的 IP 地址。

+ 这样的直接后果就是 *DNS* 收到的域名解析请求的来源 *IP* 还是成了其它运营商的 *IP*，最终导致用户流量被导向了错误的 *IDC*，用户访问变慢。