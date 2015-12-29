
GCC = 6g
LINK = 6l
FLAGS = 

%.6: %.go
	$(GCC) $(FLAGS)  $<
	$(LINK) -o $* $@

OBJS = \
	Server.6


all: install $(OBJS)

Server.6:: dictionary/Dictionary.go \
	flashcards/FlashCards.go \
	templatefuncs/PinyinFormatter.go

# Other dependencies

install:
	(cd dictionary; make install)
	(cd templatefuncs; make install)
	(cd flashcards; make install)

clean:
	rm -f  $(OBJS) *~ */*~  *.zip core
	rm -rf *~
	echo $(OBJS) ' ' | xargs -I {} -d ' ' basename {} .6 | xargs rm -f
	(cd dictionary; make clean)
	(cd templatefuncs; make clean)
	(cd flashcards; make clean)







