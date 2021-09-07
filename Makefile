phony: build


extract-header: 
	@echo "Extracting header..."
	@cd lib/libsignal-client/rust/bridge/ffi && cbindgen --config cbindgen.toml --crate libsignal-ffi --output $(shell pwd)/lib/libsignal_ffi.h