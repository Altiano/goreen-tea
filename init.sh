[ -z $1 ] && echo "please provide the base-url parameter" && return 1
[ -z $2 ] && echo "please provide the app-name parameter" && return 1

TEMP_BASE_URL=${1}
TEMP_APP_NAME=${2}
TEMP_FULL_NAME="${1}/${2}"

TEMP_BASE_URL_PLACEHOLDER="<base-url>"
TEMP_APP_NAME_PLACEHOLDER="<app-name>"
TEMP_MODULE_NAME=gitlab.com/altiano/goreen-tea


# Rename current folder
TEMP_THEDIR=$(pwd)
cd ../
mv $TEMP_THEDIR $TEMP_APP_NAME
cd $TEMP_APP_NAME

# Adjust appname
sed -i '' "s;$TEMP_MODULE_NAME;$TEMP_FULL_NAME;g" go.mod
sed -i '' "s/$TEMP_APP_NAME_PLACEHOLDER/$TEMP_APP_NAME/g" .gitlab-ci.disable.yml

sed -i '' "s;$TEMP_BASE_URL_PLACEHOLDER;$TEMP_BASE_URL;g" k8s/base/kustomization.yaml
sed -i '' "s;$TEMP_APP_NAME_PLACEHOLDER;$TEMP_APP_NAME;g" k8s/base/kustomization.yaml
sed -i '' "s;$TEMP_APP_NAME_PLACEHOLDER;$TEMP_APP_NAME;g" k8s/prod/replace-container-envFrom.yaml
sed -i '' "s;$TEMP_APP_NAME_PLACEHOLDER;$TEMP_APP_NAME;g" k8s/prod/replace-container-name.yaml

sed -i '' "s;$TEMP_MODULE_NAME;$TEMP_FULL_NAME;g" main.go
cd di
sed -i '' "s;$TEMP_MODULE_NAME;$TEMP_FULL_NAME;g" *
cd ../src
find ./ -type f -exec sed -i '' "s;$TEMP_MODULE_NAME;$TEMP_FULL_NAME;g" {} \;
cd ..


# create .env
cp .env.example .env

# Clean up
rm ./init.sh

# Yeah
pwd
echo ""
echo "It's ready"
echo "Let's build awesomeness"
echo " 🚀 🌝 💥 "