import urllib.request
import json
import sys
# import getpot
wqy = sys.argv
str=""
with open('../commit.log', 'r') as f:
	wcr = f.readlines()

for i in wcr:
	str+=i+'\n'
f.close()
data = {
	"ToUserUid":int(wqy[1]),
	"SendToType":2,
	"SendMsgType":"TextMsg",
	"Content": "yzb又瞎push了什么怪东西进github.com/UncleMadeleine/uselessBlog："+str
}
values = urllib.parse.urlencode(data).encode(encoding='UTF8')
headers = {'Content-Type': 'application/json'}
print(data)
print(values)
print(json.dumps(data))
print(json.dumps(data).encode())
request = urllib.request.Request(url=wqy[2], headers=headers, data=json.dumps(data).encode())
response = urllib.request.urlopen(request)
