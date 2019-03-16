FROM alpine:3.8 

COPY numfiles_exporter /opt/numfiles_exporter/
COPY targets.yaml /opt/numfiles_exporter/

WORKDIR /opt/numfiles_exporter/

CMD ["./numfiles_exporter"]
