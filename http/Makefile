
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	Head.6 \
	Get.6 \
	LowLevelGet.6 \
	LowLevelMultipleGet.6 \
	FileServer.6 \
	PrintEnv.6 \
	LowLevelServer.6 \
	ProxyGet.6 \
	ProxyAuthGet.6 \
	HTTPSFileServer.6 \
	PersistentServer.6 \
	PersistentClient.6 \



all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm









