ACCOUNT=YOUR_ACCOUNT_ID
REGION=us-west-2
IMAGE_NAME=apiservice

$(aws ecr get-login --no-include-email --region ${REGION})
docker tag ${IMAGE_NAME}:latest ${ACCOUNT}.dkr.ecr.${REGION}.amazonaws.com/${IMAGE_NAME}:latest
docker push ${ACCOUNT}.dkr.ecr.${REGION}.amazonaws.com/${IMAGE_NAME}:latest