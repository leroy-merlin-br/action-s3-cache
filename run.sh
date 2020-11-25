source .env
mkdir -p temp/
go build -o temp/action-s3-cache

cd temp && ./action-s3-cache