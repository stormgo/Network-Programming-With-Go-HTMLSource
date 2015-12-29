
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	EchoClient.6 \
	EchoServer.6 \
	EchoChanClient.6 \
	EchoChanServer.6 \
	EchoRevClient.6 \
	EchoRevServer.6 \
	TestImportReceiver.6 \
	TestExportSender.6 \
	TestExportReceiver.6 \
	TestImportSender.6 \
	TestSenderReciever.6


all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm









