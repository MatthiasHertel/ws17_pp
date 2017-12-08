#!/bin/sh

cd utils/minio

while [ 1 ]
do
  if (docker inspect -f {{.State.Running}} minio_minio1_1 >/dev/null 2>&1 ) ; then

  echo 'start testing minio'

  sleep 3;

  go run FileUploader.go

  exit 1
  fi
  printf %s .
done
