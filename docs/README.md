# networker-be

DN42 config generate Sever (Backend)

# Rules

用户名需要 4-20 长度

密码需要 8-20 长度

错误情况: 400 为默认返回错误, 200 为请求成功 看情况进行解析

ErrorCode 和 Message 用于联合标识错误 ErrorCode 目前没有其他意义 解析的时候可以丢弃 

较为重要的是 ErrorMsg

不存在 返回 非200 的同时 返回正常数据结构体 是必然返回 ErrorCode 和 ErrorMsg 的结构体

# VO

## /api/v1/ APIs

### /config

获取配置文件

Method: GET

request

```json
```

return 

```json
{
    "username": "yourid like esonhugh",
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX"
}
```

### /peerinfo/list

列出互相 peer 的用户

Method: GET

request

```json
```

return

```json
[
    {
        "id": "XXX",
        "username": "yourid like esonhugh",
        "asn": "4242421234"
    }, 
    {
        "id": "XXX",
        "username": "yourid like esonhugh",
        "asn": "4242421234"
    }, 
    ...
]
```

### /peerinfo/me

快速获取当前用户的配置信息

Method: GET

request

```json
```

return

200 == success

```json
{
  "id": "XXX",
  "username": "yourid like esonhugh",
  "asn": "4242421234",
  "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
  "wireguard_key": "[BASE64 String]==",
  "dn42_ipv4": "172.20.xx.xx",
  "dn42_ipv6": "fe80::XXXX"
}
```


### /peerinfo/{id}

获取相互 peer 的用户的具体信息

Method: GET

request

```json
```

return 

200 === success

```json
{
    "id": "XXX",
    "username": "yourid like esonhugh",
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX"
}
```

### /peerinfo

上传自己的 peer 的信息 也是更新

Method: POST

request

```json
{
    "username": "==username", // 实际返回的请求不用传 直接读 Cookie 拿到用户
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX"
}
```

return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": ""
}
```

### /auth/register

注册用户

Method: POST

request
```json
{
    "username": "xxxxxx",
    "password": "xxxxxxxxxxxxxxxxx",
    "email": ""
}
```

return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": ""
}
```

### /auth/verify/{ticket}

验证邮箱

request
```json
```

return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": ""
}
```

### /auth/login

登陆鉴权

Method: POST

request
```json
{
    "username": "XXXX",
    "password": "xxxxxxxxxx"
}
```


return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": ""
}
```

# DTO

## User

``` json
{
    "id": 001,
    "username": "eson",
    "email": "data@Eson.ninja",
    "verify": true,
}
```

## Configs

``` json
{
    "id": 001,
    "username": "eson",
    "asn": "data",
    "public_access": "dn42.lv.eson.ninja",
    "wireguard_key": "This is The PublicKey",
    "dn42_ipv4": "",
    "dn42_ipv6": ""
}
```

# PO

## Table: auth

|id |username| Bcrypt(password) |email|verify|
|---|--------|------------------|-----|------|
|001|eson| pasadsijiadjaisd |eson.nin|true|

## Table: config

|id |username|asn|public_access|wireguard_key|dn42_ipv4|dn42_ipv6|
|---|--------|---|-------------|-------------|---------|---------|
|001|eson|4242420|42.eson.ninja|asdkoskdad2==|172.xxxxx|fe80::222|

## K-V: verify

key: user id

value: ticket

expire: 10min