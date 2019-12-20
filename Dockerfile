FROM alpine
ADD config.json /config.json
ADD apiservice /apiservice
ENTRYPOINT ["/apiservice"]