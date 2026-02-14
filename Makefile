c:
	go build -o bin/tarkac .

# Basically all the flags should be written before the positional argument because of 
# Go's flag package's internal working. That's why -d after assets/code.tk won't work
r:
	./bin/tarkac -d assets/code.tk

# Basically, to view docs in an HTML view
d:
	~/go/bin/pkgsite

t:
	go test ./internal/lexer
	go test ./internal/parser