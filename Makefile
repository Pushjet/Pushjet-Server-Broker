all: clean build

build:
	go get -v || true
	go build -v -x -o pushjet-broker
	echo DONE

clean:
	rm -f ./pushjet-broker

install:
        cp -v pushjet-broker /usr/bin/pushjet-broker
