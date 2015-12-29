
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	ASN1.6 \
	ASN1basic.6 \
	ASN1edge.6 \
	ASNDaytimeServer.6 \
	ASNDaytimeClient.6 \
	JSON.6 \
	SaveGob.6 \
	LoadGob.6 \
	GobEchoServer.6 \
	GobEchoClient.6 \
	SaveJSON.6 \
	LoadJSON.6 \
	JSONEchoServer.6 \
	JSONEchoClient.6 \
	Base64.6 \



all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm







