# 9front environment setup

* Bootstrap Go for i386 or amd64 from another machine as described [here](https://github.com/golang/go/wiki/Plan9)
* Extract go on Plan 9, add GOPATH, GOROOT and go's bin path to glenda's profile:

	```
	bind -a /usr/glenda/go/bin /bin
	GOPATH=/usr/glenda/go/packages
	GOROOT=/usr/glenda/go
	```
* Install Mozilla's CA certs: `hget https://curl.haxx.se/ca/cacert.pem > /sys/lib/tls/ca.pem`
* Download the git wrapper for Plan 9:

	```
	hget https://9legacy.org/9legacy/tools/git > /usr/glenda/bin/rc/git
	chmod +x /usr/glenda/bin/rc/git
	```
