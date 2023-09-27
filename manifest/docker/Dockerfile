FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ENV env  prod

ADD resource $WORKDIR/

ADD ./bin/linux_amd64/porkast-server $WORKDIR/porkast-server

RUN chmod +x $WORKDIR/porkast-server

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./porkast-server
