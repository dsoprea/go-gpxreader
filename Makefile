export GOPATH=${PWD}

GPXPARSE_EXECUTABLE_FILENAME=gpxparse

GRINTERNAL_FQ_PACKAGE=gpxreader/grinternal
GRINTERNAL_SOURCEFILES=src/${GRINTERNAL_FQ_PACKAGE}/*.go
GPXPARSE_SOURCEFILES=${GRINTERNAL_SOURCEFILES} src/gpxreader/${GPXPARSE_EXECUTABLE_FILENAME}/*.go

.PHONY: all clean

all: bin/${GPXPARSE_EXECUTABLE_FILENAME}

clean:
	rm -fr bin pkg 

bin/${GPXPARSE_EXECUTABLE_FILENAME}: ${GPXPARSE_SOURCEFILES}
	echo ${GPXPARSE_SOURCEFILES}

	go get gpxreader/${GPXPARSE_EXECUTABLE_FILENAME}
	go install gpxreader/${GPXPARSE_EXECUTABLE_FILENAME}
