
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	ArithServer.6 \
	ArithClient.6 \
	TCPArithServer.6 \
	TCPArithClient.6 \
	JSONArithServer.6 \
	JSONArithClient.6



all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm









