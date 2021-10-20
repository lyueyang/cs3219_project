import urllib3
import time

http = urllib3.PoolManager()
while True:
    r = http.request('GET', 'http://127.0.0.1/foo')
    print(r.status)

    # r1 = http.request('GET', 'http://127.0.0.1/bar')
    # print(r1.status)