# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOROOT=/usr/local/go

include $(GOROOT)/src/Make.inc

TARG=   IP \
        Mask \
        ResolveIP \
        LookupHost \
        LookupPort \
        DaytimeServer \
        GetHeadInfo \
        SimpleEchoServer \
        ThreadedEchoServer \
        UDPDaytimeServer \
        UDPDaytimeClient \
        ThreadedIPEchoServer \
        IPDaytimeClient \
        IPGetHeadInfo \
        Ping \
        EchoClient

OFILES=$(TARG:%=%.$O)

all: $(TARG)

$(TARG): %: %.$O
	$(LD) -o $@ $<

$(OFILES): %.$O: %.go Makefile
	$(GC) -o $@ $<

clean:
	rm -f *.[$(OS)] $(TARG) $(CLEANFILES)

