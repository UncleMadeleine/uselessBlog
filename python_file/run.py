import urllib.request
import json
import sys
# import getpot
wqy = sys.argv
data = {
  "ToUserUid":int(wqy[1]),
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
request = urllib.request.Request(url=wqy[2], headers=headers, data=json.dumps(data).encode())
response = urllib.request.urlopen(request)