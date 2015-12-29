

DIRS = arch golang socket serialisation protocol charsets security http template chinese html xml rpc channel  websockets 

all:
	@for d in $(DIRS);\
	do 	echo Making in $$d; \
		(cd $$d; make); \
	done;

# Other dependencies

tests:
	./testbin/testall $(DIRS)

clean:
	@for d in $(DIRS);\
	 do 	echo Making clean in $$d;\
		(cd $$d; make clean); \
	done;

pdf:
	@mkpdf $(DIRS)

distrib:
	(cd dist; wget -r http://localhost/go/index.html; echo "Dist in dist/localhost/go")




