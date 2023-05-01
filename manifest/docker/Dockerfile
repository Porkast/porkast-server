FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ENV env  prod

ADD resource $WORKDIR/

ADD ./bin/linux_amd64/guoshao-fm-web $WORKDIR/guoshao-fm-web

RUN chmod +x $WORKDIR/guoshao-fm-web

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./guoshao-fm-web
