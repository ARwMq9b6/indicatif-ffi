uname := $(shell uname -s)

ifeq ($(uname), Darwin)
    LDFLAGS := -Wl,-L$(PWD)/target/release -Wl,-Bstatic -lindicatif_ffi_c
    RUST_STATICLIB := target/release/libindicatif_ffi_c.a
    EXEXT :=
else ifeq ($(uname), Linux)
    LDFLAGS := -Wl,-L$(PWD)/target/release -Wl,-Bstatic -lindicatif_ffi_c -Wl,-Bdynamic -lpthread -ldl -lutil -lc -lrt -lm
    RUST_STATICLIB := target/release/libindicatif_ffi_c.a
    EXEXT :=
else
    LDFLAGS := -Wl,-L$(PWD)/target/release -Wl,-Bstatic -lindicatif_ffi_c -Wl,-Bdynamic -luserenv -lws2_32
    RUST_STATICLIB := target/release/indicatif_ffi_c.lib
    EXEXT := .exe
endif

build: $(RUST_STATICLIB)

# -------- c
c-example-download: indicatif-ffi-c/examples/download$(EXEXT)
	$<

indicatif-ffi-c/examples/download$(EXEXT): indicatif-ffi-c/examples/download.o indicatif-ffi-c/ruststr.o $(RUST_STATICLIB)
	$(CC) -o $@ $< indicatif-ffi-c/ruststr.o $(LDFLAGS)

indicatif-ffi-c/examples/download.o: indicatif-ffi-c/examples/download.c indicatif-ffi-c/indicatif.h indicatif-ffi-c/ruststr.h
	$(CC) -o $@ -c $< -Iindicatif-ffi-c

indicatif-ffi-c/ruststr.o: indicatif-ffi-c/ruststr.c indicatif-ffi-c/ruststr.h
	$(CC) -o $@ -c $<

# -------- go
# pure go
go-example-download: indicatif-ffi-go/examples/download$(EXEXT)
	$<

indicatif-ffi-go/examples/download$(EXEXT): indicatif-ffi-go/examples/download.go indicatif-ffi-go/indicatif.go $(RUST_STATICLIB)
	 go build -o $@ -compiler gccgo -gccgoflags '$(LDFLAGS)' $<

# cgo
go-cgo-example-download: indicatif-ffi-go-cgo/examples/download$(EXEXT)
	$<

indicatif-ffi-go-cgo/examples/download$(EXEXT): indicatif-ffi-go-cgo/examples/download.go indicatif-ffi-go-cgo/indicatif.go indicatif-ffi-c/indicatif.h indicatif-ffi-c/ruststr.h $(RUST_STATICLIB)
	export CGO_CFLAGS=-I$(PWD)/indicatif-ffi-c;\
	export CGO_LDFLAGS='$(LDFLAGS)';\
	go build -o $@ $<
#	or
#	go build -o $@ --compiler gccgo $<

# -------- py
py-example-download: indicatif-ffi-py/examples/download.py indicatif-ffi-py/.setup-develop
	python $<

indicatif-ffi-py/.setup-develop: indicatif-ffi-py/x/ffi_builder.py $(wildcard indicatif-ffi-py/indicatif/*.py) $(wildcard indicatif-ffi-c/*.[h|c])
	export C_INCLUDE_PATH=$(PWD)/indicatif-ffi-c;\
	export LIBRARY_PATH=$(PWD)/target/release;\
	cd indicatif-ffi-py;\
	python setup.py develop
	touch $@

# -------- rust lib
$(RUST_STATICLIB): src/lib.rs src/progress.rs Cargo.toml
	cargo build --release

clean:
	rm -rf target
	rm 	-rf indicatif-ffi-c/examples/download$(EXEXT) \
		indicatif-ffi-c/examples/download.o \
		indicatif-ffi-c/ruststr.o \
		indicatif-ffi-go/examples/download$(EXEXT) \
		indicatif-ffi-go-cgo/examples/download$(EXEXT) \
		indicatif-ffi-py/.setup-develop \
		indicatif-ffi-py/build \
		indicatif-ffi-py/indicatif.egg-info

.PHONY: build clean c-example-download go-example-download go-cgo-example-download py-example-download
