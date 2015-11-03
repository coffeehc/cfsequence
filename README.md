cfsequence
=====
[![GoDoc](https://godoc.org/github.com/coffeehc/cfsequence?status.png)](http://godoc.org/github.com/coffeehc/cfsequence)

#基于snowflake的一个序列号生成服务

## 数据结构



## 启动参数
-nodeid   指定节点编号,默认0
-http_ip 指定http服务的地址,默认:0.0.0.0
-http_port 指定http服务端口,默认:8888

## 服务的安装

```bash
    #安装服务
    ./cfsequence install
    service cfsequence start
    
```

## REST API

POST /service/sequences

response:

```json
{
  "code": 200,
  "msg": 25918459084800,
  "request_id": 0
}
```
GET /service/sequences/{id}

response:

```json
{
  "code": 200,
  "msg": {
    "createTime": "2015-10-09T16:34:18.87+08:00",
    "index": 84,
    "nodeId": 0,
    "sequence": 25612650523648
  },
  "request_id": 0
}
```


