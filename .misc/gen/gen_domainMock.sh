[ -z $1 ] && echo "DOMAIN_NAME is required" && exit 1

DOMAIN_NAME=${1}
DIR=./src/domain/$DOMAIN_NAME

if [ -d "$DIR" ]; then
    cd $DIR
else 
    echo "Directory: ${DIR} not found."
    exit 1
fi

mockgen -source=domain.go -destination=../mocks/mock_domain.go -package mocks
