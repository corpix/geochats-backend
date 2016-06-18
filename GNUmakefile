.DEFAULT_GOAL := all

NAME = geochats-backend
PKG  = github.com/corpix/$(NAME)

VERSION = $(shell date +'%Y%m%d%H').$(shell git rev-parse --short=8 HEAD)
LDFLAGS = \
	-X $(NAME)/cli.version=$(VERSION) \
	-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n')

-include misc.mk
-include release.mk
.PHONY: all $(NAME) release.mk release test tools save
all:
	go get github.com/tools/godep
	go get github.com/rogpeppe/godef
	go get github.com/nsf/gocode
	go get github.com/alecthomas/gometalinter
	godep restore
	gometalinter --install --update
	go install $(PKG)/...
$(NAME): all
	go build -a -ldflags "$(LDFLAGS)" -v
release.mk:
	echo 'VERSION = $(VERSION)' > $@
release: release.mk
	mkdir -p $(NAME)-"$(VERSION)"/$(PKG)

	rsync -avzr --delete \
		--filter='- $(NAME)-*' \
		--filter='- /$(NAME)' \
		--filter='- .*' \
		--filter='- *~' \
		. $(NAME)-"$(VERSION)"/$(PKG)

	tar czf $(NAME)-"$(VERSION)".tgz $(NAME)-"$(VERSION)"
test:
	go vet ./...
	go test -v ./...
lint:
	gometalinter --deadline=5m --concurrency=2 --vendor ./...
save:
	godep save $(PKG)/...
