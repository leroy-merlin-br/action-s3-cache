ACTION=$1

if [ -z "${ACTION}" ]; then
  echo "Action is not defined\nCommand: $0 [action]"
  exit 1
fi

source .env
export ACTION=$ACTION

mkdir -p temp/
go build -o temp/action-s3-cache
cd temp && ./action-s3-cache
