ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: build
build: clean
	@mkdir lib/
	@cd akatsuki-pp-ffi/ && cargo build --release && cargo test
	@cp akatsuki-pp-ffi/target/release/akatsuki_pp_ffi.dll lib/
	@cp akatsuki-pp-ffi/bindings/akatsuki_pp_ffi.h lib/
	go build -ldflags="-r $(ROOT_DIR)lib" main.go

.PHONY: clean
clean:
	rm -rf lib/ main