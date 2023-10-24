# Use the official MySQL image as the base image
FROM mysql:8.0

# Set any environment variables if necessary
ENV MYSQL_ROOT_PASSWORD=my-secret-pw
ENV MYSQL_DATABASE=eCombombDB

# Copy initialization scripts into the container
COPY ./docker/db-deployment/scripts/init_mysql/ /docker-entrypoint-initdb.d/

# Expose port 3306 for the MySQL service
EXPOSE 3306