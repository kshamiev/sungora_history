FROM registry.services.mts.ru/docker/oraclelinux:7-slim
WORKDIR /bin

COPY torgi-back /bin/torgi-back
COPY data/config.yaml.sample /bin/application/config.yaml
EXPOSE 80
CMD ["/bin/torgi-back", "-c", "/etc/application/config.yaml"]