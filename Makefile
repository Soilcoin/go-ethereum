# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gsoil android ios gsoil-cross swarm evm all test clean
.PHONY: gsoil-linux gsoil-linux-386 gsoil-linux-amd64 gsoil-linux-mips64 gsoil-linux-mips64le
.PHONY: gsoil-linux-arm gsoil-linux-arm-5 gsoil-linux-arm-6 gsoil-linux-arm-7 gsoil-linux-arm64
.PHONY: gsoil-darwin gsoil-darwin-386 gsoil-darwin-amd64
.PHONY: gsoil-windows gsoil-windows-386 gsoil-windows-amd64

GOBIN = build/bin
GO ?= latest

gsoil:
	build/env.sh go run build/ci.go install ./cmd/gsoil
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gsoil\" to launch gsoil."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gsoil.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gsoil-cross: gsoil-linux gsoil-darwin gsoil-windows gsoil-android gsoil-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-*

gsoil-linux: gsoil-linux-386 gsoil-linux-amd64 gsoil-linux-arm gsoil-linux-mips64 gsoil-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-*

gsoil-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gsoil
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep 386

gsoil-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gsoil
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep amd64

gsoil-linux-arm: gsoil-linux-arm-5 gsoil-linux-arm-6 gsoil-linux-arm-7 gsoil-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep arm

gsoil-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gsoil
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep arm-5

gsoil-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gsoil
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep arm-6

gsoil-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gsoil
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep arm-7

gsoil-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gsoil
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep arm64

gsoil-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gsoil
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep mips

gsoil-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gsoil
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep mipsle

gsoil-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gsoil
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep mips64

gsoil-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gsoil
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-linux-* | grep mips64le

gsoil-darwin: gsoil-darwin-386 gsoil-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-darwin-*

gsoil-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gsoil
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-darwin-* | grep 386

gsoil-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gsoil
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-darwin-* | grep amd64

gsoil-windows: gsoil-windows-386 gsoil-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-windows-*

gsoil-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gsoil
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-windows-* | grep 386

gsoil-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gsoil
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gsoil-windows-* | grep amd64
