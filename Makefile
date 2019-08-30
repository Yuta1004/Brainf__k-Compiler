GCC = gcc
GO = go

run:
	$(GO) run main.go > bfc.S
	$(GCC) -o bfc bfc.S
	./bfc
