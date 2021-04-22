FROM grafana/grafana:master

COPY provisioning /etc/grafana/provisioning
COPY dist /var/lib/grafana/plugins/grafana-intersystems-datasource/

ENV COMPOSE_INTERACTIVE_NO_CLI=1
ENV TERM=linux
ENV GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
ENV GF_AUTH_ANONYMOUS_ENABLED=true
ENV GF_AUTH_BASIC_ENABLED=false
ENV GF_LOG_LEVEL=debug
ENV GF_DATAPROXY_LOGGING=true
ENV GF_FEATURE_TOGGLES_ENABLE="expressions inspect transformations newEdit live"
ENV GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS=grafana-intersystems-datasource
