# Archived, you can use [vmessconfig](https://github.com/yindaheng98/vmessconfig) for alternative.

# v2autosub

一个小小的go程序，通过订阅链接自动获取vmess订阅信息、测量连接情况、并生成包含多个outbound的负载均衡配置文件。可用于实现自动订阅+负载均衡。

## 如何使用？

```sh
git clone https://github.com/yindaheng98/v2autosub
cd v2autosub
go build github.com/yindaheng98/v2autosub
v2autosub -h
```

## vmess订阅信息是如何获取的？连接情况是如何测量的？

基于[v2gen](https://github.com/iochen/v2gen)

## 配置文件是如何生成的？

配置文件的生成都是基于模板替换的：

1. 首先，对于vmess订阅中的每个服务器，借助[v2gen](https://github.com/iochen/v2gen)的模板功能，用命令行参数`-node_template`所指定**outbound配置模板**生成单独的outbound配置
2. 将每个单独的outbound配置拼接好后替换到命令行参数`-nodes_template`所指定**全局配置模板**的`{{outbounds}}`处，然后将生成的tags替换到`{{tags}}`处

所以，在生成负载均衡配置的时候，编写好单独的outbound配置模板之后，在编写全局配置模板时把`{{outbounds}}`放在配置文件的outbounds配置里，并把`{{tags}}`放在配置文件的balancer里即可。

## 有哪些可调参数？

```
$ v2autosub -h
Usage of ./v2autosub:
  -c int
    	ping count for each node (default 3)
  -config string
    	v2gen config path (default "/etc/v2ray/v2gen.ini")
  -dst string
    	test destination url (vmess ping only) (default "https://cloudflare.com/cdn-cgi/trace")
  -max int
    	max number of nodes (default 8)
  -node_tag_fmt string
    	format of the tag for nodes (default "v2gen_%d")
  -node_template string
    	V2Ray template path for single nodes
  -nodes_template string
    	V2Ray template path for all nodes
  -o string
    	output path (default "/etc/v2ray/config.json")
  -ping
    	ping nodes (default true)
  -thread int
    	threads used when pinging (default 3)
  -u string
    	subscription address(URL)

```
