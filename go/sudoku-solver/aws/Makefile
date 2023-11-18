PHONY: deploy

deploy: build
	sam deploy --profile catdevman --stack-name sudoku-solver --region us-east-2 --resolve-s3 --capabilities CAPABILITY_IAM

build:
	sam build
