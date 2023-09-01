# build
build:
	go build -o main main.go
clean:
	rm -rf main && go clean

help:
	go run main.go -h
helpD:
	go run main.go -D -h

# word
wdh:
	go run main.go -D help word
we:
	go run main.go word -m 5 -s YaoQiJun

# file convert
fch:
	go run main.go -D convert -h
fc:
	go run main.go convert -f dest/from.txt -t dest/to.txt

# model
mh:
	go run main.go -D model -h
m:
	go run main.go model -d dest

# api
ad:
	go run main.go -D api -h
abrick:
	go run main.go api -p brick -a 315713 -d dest
ahxxmis:
	go run main.go api -p hxxmis -a 643396 -d dest
aunit:
	go run main.go api -p unit -a 670477 -d dest
