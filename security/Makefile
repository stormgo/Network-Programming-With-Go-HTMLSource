
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	MD5Hash.6 \
	Blowfish.6 \
	GenRSAKeys.6 \
	LoadRSAKeys.6 \
	GenX509Cert.6 \
	TLSEchoServer.6 \
	TLSEchoClient.6


all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm









