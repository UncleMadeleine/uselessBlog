import urllib.request
import json


data = {
  "ToUserUid":393826682,
  "SendToType":1,
  "SendMsgType":"TextMsg",
  "Content":"test"
}
values = urllib.parse.urlencode(data).encode(encoding='UTF8')//注释1
headers = {'Content-Type': 'application/json'}
print(data)
print(values)
print(json.dumps(data))
print(json.dumps(data).encode())
request = urllib.request.Request(url='http://121.40.124.19:8888/v1/LuaApiCaller?qq=1004619265&funcname=SendMsgV2&timeout=100', headers=headers, data=json.dumps(data).encode())
response = urllib.request.urlopen(request)