
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	StrLength.6 \
	UTF8Server.6 \
	UTF16.6 \
	UTF16Server.6 \
	UTF16Client.6 \
	ISO8859.6



all: $(OBJS)

# Other dependencies

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm









