FROM envoyproxy/envoy:v1.34.1

COPY envoy.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /dev/stdout