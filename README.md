# sleepy-reverse-proxy
Tool to help test application performance when third-party APIs slow down.

### Installing
```sh
go get github.com/dpromanko/sleepy-reverse-proxy
```

### Usage with flags
```
./sleepy-reverse-proxy -h

Usage of ./sleepy-reverse-proxy:
  -p string
    	Port this application will run on
  -s duration
    	Duration to sleep for (ex: 50us, 1000ms, 1s, 1h)
  -u string
    	Proxy URL


./sleepy-reverse-proxy -p=8080 -s=1000ms -u=http://the-api-you-want-to-hit:8080
```
> Sleep duration supports any format defined by Go's [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration)

### Usage with environment variables
The following environment variables can be set in favor of using the flags above
```
export PORT=8080
export SLEEP_DURATION=1000ms
export PROXY_URL=http://the-api-you-want-to-hit:8080

./sleepy-reverse-proxy
```

> Sleep duration supports any format defined by Go's [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration)