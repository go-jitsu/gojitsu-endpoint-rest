FROM fedora:26

MAINTAINER hekonsek@gmail.com

ADD gojitsu-endpoint-rest /usr/bin/

ENTRYPOINT ["/usr/bin/gojitsu-endpoint-rest"]