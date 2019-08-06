default: travis

travis: link fromScratch

link:
	mkdir -p $GOPATH/src/opensource2fa
	ln -s $GOPATH/src/opensource2fa/server $PWD

fromScratch:
	./build.sh fromscratch
