GCC = gcc
GO = go
PROGRAM = program.txt

run: $(PROGRAM)
	$(GO) run main.go $(PROGRAM) > bfc.S
	$(GCC) -o bfc bfc.S
	./bfc
