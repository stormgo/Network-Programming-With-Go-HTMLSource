
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	EchoServer.6 \
	EchoClient.6 \
	PersonClientJSON.6 \
	PersonServerJSON.6 \
	PersonClientXML.6 \
	PersonServerXML.6 \
	EchoServerTLS.6 \


	#PersonClientProtocol.6 \
	#PersonServerProtocol.6 \
all: install $(OBJS) 

install:
	(cd xml; make install)


# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm
	(cd xml; make clean)









