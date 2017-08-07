NAME    := $(shell basename $(shell pwd))
VERSION ?= 0.0.0

all: test

clean: cleanbuild cleanpkg
cleanbuild:
	rm -fr ./build
cleanpkg:
	rm -fr ./pkg

show:
	echo "Building '${NAME}'"

test:
	go test -ldflags # "-X main.name=${NAME}"

build: build/${NAME}
build/%:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o build/$*

pkg: show build pkg/${NAME}.deb
pkg/%.deb:
	mkdir -p ./pkg
# 	https://github.com/jordansissel/fpm/wiki
	fpm --verbose -s dir -t deb \
		--name ${NAME} \
		--package ./pkg/${NAME}.deb \
		--force \
                --category admin \
		--epoch $(shell /bin/date +%s) \
		--iteration $(VERSION) \
		--deb-compression bzip2 \
		--url https://github.com/go-cc/cc2py \
		--description "Chinese-Character to Pinyin converter" \
		--maintainer "Tong Sun <suntong@cpan.org>" \
		--license "MIT" \
                --vendor "contrib" \
		--version $(VERSION) \
		--architecture amd64 \
		--depends apt \
		./build/=/usr/bin/

#		--category universe/text \
