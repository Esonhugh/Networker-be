# networker-be

DN42 config generate Sever (Backend)

# API

## /api/v1/

### /config

获取配置文件

Method: GET

request

```json
```

return 

```json
{
    "nickname": "yourid like esonhugh",
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX",
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
        "nickname": "yourid like esonhugh",
        "asn": "4242421234"
    }, 
    {
        "id": "XXX",
        "nickname": "yourid like esonhugh",
        "asn": "4242421234"
    }, 
    ...
]
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
    "nickname": "yourid like esonhugh",
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX",
}
```

### /peerinfo

上传自己的 peer 的信息

Method: POST

request

```json
{
    "nickname": "==username",
    "asn": "4242421234",
    "public_access": "XXXX.dn42.youdomain.com or IPv4 IPv6",
    "wireguard_key": "[BASE64 String]==",
    "dn42_ipv4": "172.20.xx.xx",
    "dn42_ipv6": "fe80::XXXX",
}
```

return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": "",
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
    "email": "",
}
```

return 

200 === success

400 === fail

```json
{
    "errorcode": "",
    "errormsg": "",
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
    "errormsg": "",
}
```

# DTO

## Table: auth

|id |username|MD5(password)|email|verify|
|---|--------|-------------|-----|------|
|001|eson|pasadsijiadjaisd|eson.nin|true|

## Table: config

|id |username|asn|public_access|wireguard_key|dn42_ipv4|dn42_ipv6|
|---|--------|---|-------------|-------------|---------|---------|
|001|eson|4242420|42.eson.ninja|asdkoskdad2==|172.xxxxx|fe80::222|
