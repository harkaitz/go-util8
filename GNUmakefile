DESTDIR =
PREFIX  =/usr/local


all:
clean:
install:
update:
## -- AUTO-GO --
GO_PROGRAMS += bin/$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
## -- AUTO-SERVICE --

## -- AUTO-SERVICE --
