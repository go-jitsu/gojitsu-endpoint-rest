FROM fedora

ADD gojitsu-endpoint-rest /usr/bin/

ENTRYPOINT ["/usr/bin/gojitsu-endpoint-rest"]