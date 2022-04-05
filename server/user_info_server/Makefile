
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
	cp -rf ${BASE_DIR}/configs/* ${TARGET_CONF}
endef

all:
	$(call compile_and_pack)

clean:
	@rm -rf ${RELEASE}
