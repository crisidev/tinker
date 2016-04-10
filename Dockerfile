FROM debian:jessie

RUN echo 'APT::Install-Suggests "0";' > /etc/apt/apt.conf.d/99local
RUN echo 'APT::Install-Recommends "0";' >> /etc/apt/apt.conf.d/99local
RUN apt-get update && apt-get -y install wget ca-certificates locales
RUN echo "en_IE.UTF-8 UTF-8\nen_US.UTF-8 UTF-8" > /etc/locale.gen
RUN locale-gen


