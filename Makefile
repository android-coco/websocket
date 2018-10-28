#/bin/bash
# This is how we want to name the binary output
OUTPUT=bin/websocket
SRC=cmd/websocket.go

# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`



# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME}"

local:
	rm -f ./bin/websocket
	go build ${LDFLAGS} -o ${OUTPUT} ${SRC}

debug:
	rm -f ./bin/websocket
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT} ${SRC}

release:
	rm -f ./bin/websocket
	CGO_ENABLED=0 GOOS=linux go build ${LDFLAGS} -o ${OUTPUT} ${SRC}
	cd .. && tar -zcvf eospark_release.tar.gz --exclude build --exclude .git --exclude .gitmodules --exclude pkg --exclude src --exclude Makefile eospark && mv -f eospark_release.tar.gz eospark/build/eospark_release_${BUILD_TIME}.tar.gz && rm -f eospark/bin/eospark

clean:
	rm -f ./bin/websocket
	rm -f ./build/websocket_*