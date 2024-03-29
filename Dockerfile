FROM docker.elastic.co/logstash/logstash:7.17.0

# Download JDBC connector for Logstash
RUN curl -L --output "mysql-connector-j-8.0.31.tar.gz" "https://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-j-8.0.31.tar.gz" \
    && tar -xf "mysql-connector-j-8.0.31.tar.gz" "mysql-connector-j-8.0.31/mysql-connector-j-8.0.31.jar" \
    && mv "mysql-connector-j-8.0.31/mysql-connector-j-8.0.31.jar" "mysql-connector-j-8.0.31.jar" \
    && rm -r "mysql-connector-j-8.0.31" "mysql-connector-j-8.0.31.tar.gz"

ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]