all:
	go build

filter:	all

install:	filter
	cp $< ${HOME}/bin
