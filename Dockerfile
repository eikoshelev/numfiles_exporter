FROM alpine:3.8 

COPY numfiles_exporter /bin/
COPY targets.yaml /opt/numfiles_exporter/

CMD ["numfiles_exporter"] # specify the required flags!
