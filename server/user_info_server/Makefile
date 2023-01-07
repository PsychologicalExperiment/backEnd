
.PHONY: all

NAME=user_info_server
BASE_DIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
RELEASE=${BASE_DIR}/release
TARGET_BIN=${RELEASE}/bin
TARGET_CONF=${RELEASE}/conf
TARGET_LOG=${RELEASE}/log

DATE=$(shell date +"%Y%m%d%H%M%S")

FLAGS:= -X 'main.gitversion=$(shell git rev-parse --short HEAD)'
FLAGS+= -X 'main.gitbranch=$(shell git branch | grep "\*" | grep " .*" -o | sed "s/\ //g")'
FLAGS+= -X 'main.buildstamp=${DATE}'
FLAGS+= -X 'main.goversion=$(shell go version)'
FLAGS+= -X 'main.sysuname=$(shell uname -mns)'

GOBUILD= go build

define compile_and_pack
    @rm -rf ${RELEASE}
	@mkdir -p ${TARGET_BIN}
	@mkdir -p ${TARGET_CONF}
	@mkdir -p ${TARGET_LOG}
	$(GOBUILD) -o ${TARGET_BIN}/$(NAME) -ldflags "${FLAGS}" ${BASE_DIR}/cmd/main.go
	cp -rf ${BASE_DIR}/configs/* ${ffile}
endef

all:
	$(call compile_and_pack)

clean:
	@rm -rf ${RELEASE}

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${NAME} cmd/main.go
	ssh Javis-root "sudo rm -rf /data/server/${NAME} &> /dev/null"
	ssh -q Javis-root "kill -9 `ps -ef | grep -v grep | grep ${NAME} | awk '{print $$2}'` > /dev/null 2>&1 &"
	scp bin/${NAME} Javis-root:/data/server/
	ssh Javis-root "cd /data/server/ && chmod 777 ${NAME} && ./${NAME} > /dev/null 2>&1 &"

.PHONY: all build clean