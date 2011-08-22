include $(GOROOT)/src/Make.inc

.PHONY: mathgl install examples clean

all: mathgl

mathgl:
	gomake -C src

install:
	gomake -C src install

test:
	gomake -C src test

clean:
	gomake -C src clean
	rm -Rf out
