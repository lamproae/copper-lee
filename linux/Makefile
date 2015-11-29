.EXPORT_ALL_VARIABLES:

PROJECT_DIR	=	$(PWD)
MAKE_DIR	=	$(PROJECT_DIR)/make
APPS_DIR	=	$(PROJECT_DIR)/apps
MODULES_DIR	=	$(PROJECT_DIR)/modules
KERNEL_DIR	=	$(PROJECT_DIR)/kernel
SCRIPT_DIR	=	$(PROJECT_DIR)/script

BUILD_DIR	=	$(PROJECT_DIR)/build
ROOT_DIR	=	$(PROJECT_DIR)/rootfs
MODULES_DIR	=	$(PROJECT_DIR)/modules

ifeq ($(ARCH),)
    $(error Please set the target platform!) 
endif

all: prepare config build 

boot:
	sh $(SCRIPT_DIR)/boot.sh

prepare: 
	@case $(ARCH) in  \
		x86_64 | arm)	\
			ln -sf $(PROJECT_DIR)/make/rule/$(ARCH).mk $(PROJECT_DIR)/make/.rule;;  \
		*) 	\
			echo "Unsupported archetecture!";;  \
	esac

build:
	$(MAKE) -C $(PROJECT_DIR)/make 

modules:
	$(MAKE) -C $(PROJECT_DIR)/make $@

kernel:
	$(MAKE) -C $(PROJECT_DIR)/make $@

config:
	$(MAKE) -C $(PROJECT_DIR)/make $@

lib:
	$(MAKE) -C $(PROJECT_DIR)/make $@

clean:
	$(MAKE) -C $(PROJECT_DIR)/make $@

distclean:
	$(MAKE) -C $(PROJECT_DIR)/make $@

.PHONY: all build clean prepare modules make kernel
