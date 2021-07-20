SHELL=/bin/bash

EXE = ws-service

all: $(EXE)

ws-service:
	@echo "building $@ ..."
	$(MAKE) -s -f make.inc s=static

clean:
	rm -f $(EXE)

