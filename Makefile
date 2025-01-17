# Makefile for cross-compilation
# Window icon Need：go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo

# make all VERSION=v0.9.12
# make md5SumThemAll VERSION=v0.9.9
# mingw32-make all VERSION=v0.9.9
# rm  resource.syso && make  Linux-armv8   VERSION=v0.9.6
# make Windows_i386  VERSION=v1.0.0
# make Windows_x86_64

# Windows：MSYS2 or mingw32 + find.exe make.exe zip.exe upx.exe ！！

#我应该下载哪个版本？
#
#MacOS intel芯片（2020年以前的Mac）: MacOS_x86_64.tar.gz
#MacOS apple （2020年以后的新Mac，M系列芯片）: MacOS_arm64.tar.gz
#Linux ARM 32位（树莓派1、树莓派Zero等老设备）: Linux_armv6.tar.gz
#Linux ARM 32位（树莓派2~4，其他arm设备）: Linux_armv7.tar.gz
#Linux ARM 64位（树莓派4或树莓派5，安装64位的ARM系统的时候）: Linux_armv8.tar.gz
#
#Windows 64位（大多数Windows用户）: Windows_x86_64.zip
#Windows 32位（比较老的Windows）: Windows_i386.zip
#Windows ARM版（Windows 骁龙Elite等）：Windows_arm64.zip

NAME=comi
SKETCH_NAME=sketch_66seconds
OS := $(shell uname)
BINDIR := ./bin
MD5_TEXTFILE := $(BINDIR)/md5Sums.txt
#go: cannot install cross-compiled binaries when GOBIN is set
unexport GOBIN

MAIN_FILE_DIR := ./
# -ldflags 指定编译参数。-s 去掉符号信息。 -w去掉调试信息。
GOBUILD=CGO_ENABLED=0 go build -ldflags "-s -w -X config.Version=${VERSION}"

ifeq ($(OS), Darwin)
  MD5_UTIL = md5
else
  MD5_UTIL = md5sum
endif

all: compileThemAll md5SumThemAll

# 因为sqlite（ent）库的关系，部分架构（Windows_i386）无法正常运行，需要写条件编译代码。 ent库的编译检测状态： https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2Fsqlite

compileThemAll: Windows_x86_64 Windows_i386  Windows_arm64 Linux_x86_64 Linux_i386 MacOS_x86_64 MacOS_arm64 Linux-armv6 Linux-armv7 Linux-armv8 

android: Linux-arm-android Linux-arm64-android

UPX := $(shell command -v upx 2> /dev/null)

gomobile:
	export ANDROID_NDK_HOME=/Users/bai/Library/Android/sdk/ndk/26.1.10909125
	gomobile bind -target=android -o comigo.aar -androidapi 26

md5SumThemAll:
	rm -f $(MD5_TEXTFILE)
	find $(BINDIR) -type f -name "$(NAME)_*" -exec $(MD5_UTIL) {} >> $(MD5_TEXTFILE) \;
	# 删除 $(MD5_TEXTFILE)里面的 ./bin/ 字符串
	sed -i '' 's|./bin/||g' $(MD5_TEXTFILE)
	cat $(MD5_TEXTFILE)

# upx可能导致报毒，取消windows平台的upx压缩
# 换行用TAB而不是空格
#64位Windows	$(NAME)_$(VERSION)_$@   
Windows_x86_64:
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo # Window icon Need
	GOARCH=amd64 GOOS=windows go generate #go: cannot install cross-compiled binaries when GOBIN is set
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME).exe 
	zip -m -r -j -9 $(BINDIR)/$(NAME)_$(VERSION)_$@.zip $(BINDIR)/$(NAME)_$(VERSION)_$@
	rmdir $(BINDIR)/$(NAME)_$(VERSION)_$@
	rm  resource.syso 

#32位Windows	
Windows_i386:
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo # Window icon Need
	GOARCH=386 GOOS=windows go generate
	GOARCH=386 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME).exe 
	zip -m -r -j -9 $(BINDIR)/$(NAME)_$(VERSION)_$@.zip $(BINDIR)/$(NAME)_$(VERSION)_$@
	rmdir $(BINDIR)/$(NAME)_$(VERSION)_$@
	rm   resource.syso	

#windows arm64 no upx
Windows_arm64:
	GOARCH=arm64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME).exe 
	zip -m -r -j -9 $(BINDIR)/$(NAME)_$(VERSION)_$@.zip $(BINDIR)/$(NAME)_$(VERSION)_$@
	rmdir $(BINDIR)/$(NAME)_$(VERSION)_$@
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

# golang支持的交叉编译架构的列表，参见 go tool dist list
# 看ARM处理器是否支持VFP功能:grep -i vfp /proc/cpuinfo

##Linux-armv5,GOARM=5：使用软件浮点；当 CPU 没有 VFP 协处理器时
#Linux-armv5:
#	GOARCH=arm GOOS=linux GOARM=5 $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
##ifdef UPX
##	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
##endif
#	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
#	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#Linux-armv6 RaspberryPi1,2,zero,GOARM=6：仅使用 VFPv1；交叉编译时默认；通常是 ARM11 或更好的内核（也支持 VFPv2 或更好的内核）
Linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#Linux-armv7，RaspberryPi3 官方32位armv7l系统。GOARM=7：使用 VFPv3；通常是 Cortex-A 内核.
Linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#linux，64位arm
Linux-armv8:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#Linux，x86_64
Linux_x86_64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#Linux，i386
Linux_i386:
	GOARCH=386 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#MACOS x86_64
MacOS_x86_64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#ifdef UPX
#	upx -9 $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
#endif
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@
	
#MACOS arm64 no upx
MacOS_arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME)
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@
	
#Android，32位arm，Termux	
Linux-arm-android:
	GOARCH=arm GOOS=android $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@

#Android，64位arm，Termux
Linux-arm64-android:
	GOARCH=arm64 GOOS=android $(GOBUILD) -o $(BINDIR)/$(NAME)_$(VERSION)_$@/$(NAME) 
	tar --directory=$(BINDIR)/$(NAME)_$(VERSION)_$@  -zcvf $(BINDIR)/$(NAME)_$(VERSION)_$@.tar.gz $(NAME)
	rm -rf $(BINDIR)/$(NAME)_$(VERSION)_$@
	
