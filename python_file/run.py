import urllib.request
import json


data = {
  "ToUserUid":776223056,
  "SendToType":2,
  "SendMsgType":"TextMsg",
  "Content":"yzb又瞎push了什么怪东西"
}
values = urllib.parse.urlencode(data).encode(encoding='UTF8')
headers = {'Content-Type': 'application/json'}
print(data)
print(values)
print(json.dumps(data))
print(json.dumps(data).encode())
request = urllib.request.Request(url='http://121.40.124.19:8888/v1/LuaApiCaller?qq=1004619265&funcname=SendMsgV2&timeout=100', headers=headers, data=json.dumps(data).encode())
response = urllib.request.urlopen(request)