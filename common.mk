# Usage: echo 'include common.mk' > Makefile

TOP_PKG := $(shell go list .)
NAME := $(shell go list . | awk -F/ '{print $$NF}')
PKGS := $(shell go list ./... | grep -v "^$(TOP_PKG)/\(vendor\|prototype\)")

# PKGS := github.com/openacid/slimarray/array \
#         github.com/openacid/slimarray/bit \
#         github.com/openacid/slimarray/trie \

SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))

# gofmt check vendor dir. we need to skip vendor manually
GOFILES := $(shell find $(SRCDIRS) -not -path "*/vendor/*" -name "*.go")
GO := go

check: test vet gofmt misspell unconvert staticcheck ineffassign unparam

travis: vet gofmt misspell unconvert ineffassign unparam test

test:
	# fail fast with severe bugs
	$(GO) test -short      $(PKGS)
	$(GO) test -tags debug $(PKGS)
	# test release version and generate coverage data for task `coveralls`.
	$(GO) test -covermode=count -coverprofile=coverage.out $(PKGS)

vet:
	$(GO) vet $(PKGS)

staticcheck:
	$(GO) get honnef.co/go/tools/cmd/staticcheck
	# ST1016: methods on the same type should have the same receiver name
	#         .pb.go have this issue.
	staticcheck -checks all,-ST1016 $(PKGS)

misspell:
	$(GO) get github.com/client9/misspell/cmd/misspell
	find $(SRCDIRS) -name '*.go' -or -name '*.md' | grep -v "\bvendor/" | xargs misspell \
		-locale US \
		-error
	misspell \
		-locale US \
		-error \
		*.md *.go

unconvert:
	$(GO) get github.com/mdempsky/unconvert
	unconvert -v $(PKGS)

ineffassign:
	$(GO) get github.com/gordonklaus/ineffassign
	find $(SRCDIRS) -name '*.go' | grep -v "\bvendor/" | xargs ineffassign

pedantic: check errcheck

unparam:
	$(GO) get mvdan.cc/unparam
	unparam ./...

errcheck:
	$(GO) get github.com/kisielk/errcheck
	errcheck $(PKGS)

gofmt:
	@echo Checking code is gofmted
	@test -z "$(shell gofmt -s -l -d -e $(GOFILES) | tee /dev/stderr)"

ben: test
	$(GO) test ./... -run=none -bench=. -benchmem

gen:
	$(GO) generate ./...

doc:
	# build a markdown version of package doc to embed to README.md
	# $(GO) get github.com/robertkrimen/godocdown/godocdown
	godocdown . > docs/$(NAME).md
	# "package" is the first phrase in a go doc.
	# "## Usage" is the start of API section.
	cat docs/$(NAME).md | awk '/^package /,/^## Usage/' | grep -v '^## Usage' > docs/$(NAME)-package.md


readme: doc
	python ./scripts/build_md.py
	# brew install nodejs
	# npm install -g doctoc
	doctoc --title '' --github README.md

fix:
	gofmt -s -w $(GOFILES)
	unconvert -v -apply $(PKGS)


# local coverage
coverage:
	$(GO) test -covermode=count -coverprofile=coverage.out $(PKGS)
	go tool cover -func=coverage.out
	# go tool cover -html=coverage.out

# send coverage to coveralls
coveralls:
	# this job relies on the output of task test: `-coverprofile=coverage.out`
	$(GO) get golang.org/x/tools/cmd/cover
	$(GO) get github.com/mattn/goveralls
	goveralls -ignore='*.pb.go' -coverprofile=coverage.out -service=travis-ci
