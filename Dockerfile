FROM alpine:latest

RUN apk add ca-certificates --no-cache
COPY dailymotion_updater.linux64 /usr/local/bin/dailymotion_updater
RUN chmod a+x /usr/local/bin/dailymotion_updater
USER daemon
CMD /usr/local/bin/dailymotion_updater