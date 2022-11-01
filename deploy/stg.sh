APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME} #stop container ${APP_NAME}

docker run -d --name ${APP_NAME} \
--network my-net \
-e VIRTUAL_HOST="g05.custohub.com" \
-e LETSENCRYPT_HOST="g05.custohub.com" \
-e LETSENCRYPT_EMAIL="stg@g05.custohub.com" \
-e MYSQL_CONN_STRING="demo:root_password@tcp(mysql:3306)/fd?charset=utf8mb4&parseTime=True&loc=Local" \
-e S3BucketName=g123456-my-bucket \
-e S3Region=ap-southeast-1 \
-e S3APIKey= AKIAZDLBDNV4DS54J4LG\
-e S3SecretKey="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" \
-e S3Domain="https://d3s5ma63l4xcbq.cloudfront.net" \
-e SYSTEM_SECRET=dogsupercute \
-p 8080:8080 \
${APP_NAME}
