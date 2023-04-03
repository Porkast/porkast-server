FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ADD resource $WORKDIR/

ADD ./bin/linux_amd64/guoshao-fm-web $WORKDIR/main

RUN chmod +x $WORKDIR/guoshao-fm-web

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./guoshao-fm-web
