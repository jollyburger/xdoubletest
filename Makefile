SRCDIRS_EXCLUDE = log test utils proto scripts vendor Godeps conf logic examples schema build deploy
SRCDIRS_ALL = $(sort $(subst /,,$(dir $(wildcard */*))))
SRCDIRS = $(filter-out $(SRCDIRS_EXCLUDE), $(SRCDIRS_ALL))

PKGDIRS_EXCLUDE=$(GOROOT)/pkg
PKGDIRS_ALL = $(addsuffix /pkg, $(subst :, ,$(GOPATH)))
PKGDIRS = $(filter-out $(PKGDIRS_EXCLUDE), $(PKGDIRS_ALL))

all: build_main 
	@for subdir in $(SRCDIRS);do \
		cd $$subdir; go install; cd ..; \
	done 

build_main: 
	@CGO_ENABLED=0 go build -o build/xdoubletest main.go;

show:
	@echo "==================src====================="
	@echo SRCDIRS_ALL: $(SRCDIRS_ALL)
	@echo SRCDIRS_EXCLUDE: $(SRCDIRS_EXCLUDE)
	@echo SRCDIRS: $(SRCDIRS)
	@echo "==================pkg====================="
	@echo PKGDIRS_EXCLUDE: $(PKGDIRS_EXCLUDE)
	@echo PKGDIRS_ALL: $(PKGDIRS_ALL)
	@echo PKGDIRS: $(PKGDIRS)
	@echo "================clean====================="
	@for subdir in $(PKGDIRS); do \
		cd $$subdir;\
		module_name=`echo $(CURDIR)|awk -F"/" '{print $$(NF)}'`;\
		result=`find . |grep $$module_name |head -n1|awk -F"." '{print $$2}'`; \
		if [ -n "$$result" ];then \
			echo clean_dirs:$$subdir$$result; \
		fi \
	done

clean:
	@for subdir in $(PKGDIRS); do \
		cd $$subdir;\
		module_name=`echo $(CURDIR)|awk -F"/" '{print $$(NF)}'`;\
		result=`find . |grep $$module_name|head -n1|awk -F"." '{print $$2}'`; \
		if [ -n "$$result" ];then \
			cd $$subdir$$result; rm -rf *; \
		fi \
	done
