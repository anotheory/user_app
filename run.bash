#!/bin/bash
echo "Start"
cd ./Liquibase

PROJECT_NAME=user_app
dbUsername="postgres"
dbPassword=""
dbUrl="jdbc:postgresql://postgres:5432/postgres?user=${dbUsername}&password=${dbPassword}"

cd /app/Liquibase/workspace

java -jar liquibase.jar \
    --driver=org.postgresql.Driver \
    --classpath=./lib/postgresql.jar \
    --url=$dbUrl \
    --defaultSchemaName=public \
    --databaseChangeLogLockTableName=${PROJECT_NAME}_databasechangeloglock \
    --databaseChangeLogTableName=${PROJECT_NAME}_databasechangelog \
    --changeLogFile=./changelog/changelog-initschema.xml update;

java -jar liquibase.jar \
	--driver=org.postgresql.Driver \
	--classpath=./lib/postgresql.jar \
	--url=$dbUrl \
	--defaultSchemaName=user_management \
	--changeLogFile=./changelog/changelog-master.xml \
	update;

cd /app && /user_app