FROM scratch
ADD CaPgConnector ./CaPgConnector
ADD ca-certificates.crt /etc/ssl/certs/

EXPOSE 9000
ENTRYPOINT ["/CaPgConnector"]