NAME ?= hms-bss 
VERSION ?= $(shell cat .version)

all : image 

image:
		docker build --pull ${DOCKER_ARGS} --tag '${NAME}:${VERSION}' .

