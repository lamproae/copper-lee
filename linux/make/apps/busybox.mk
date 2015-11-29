SOURCE=$(APPS_DIR)/busybox/busybox-$(VERSION)
VERSION=1.23.2

all: build install 

build: config
	@cd $(SOURCE) && make

config:
	cd $(SOURCE) && $(MAKE) defconfig ARCH=$(ARCH) CFLAGS="$(CFLAGS)" LDFLAGS="$(LDFLAGS)"

install:
	@cd $(SOURCE) && $(MAKE) CONFIG_PREFIX=$(ROOT_DIR) install

clean:
	@cd $(SOURCE) && make clean

distclean:
	@cd $(SOURCE) && make distclean

.PHONY: config build clean install
