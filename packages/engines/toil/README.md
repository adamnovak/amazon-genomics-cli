## Toil AWS Mirror

A Toil mono-container WES server for use with Amazon Genomics CLI.

### Building the Container Manually

Go to this directory and run:

```bash
docker build . -f Dockerfile -t toil-agc
```

### Running for Testing

Having built the container, run:

```bash
docker run --name toil-agc-test -ti --rm -p "127.0.0.1:8000:8000" toil-agc
```

This will start the containerized server and make it available on port 8000 on the loopback interface. You can inspect the port mapping with:

```bash
docker port toil-agc-test
```

Then you can talk to it with e.g.:

```bash
curl -vvv "http://localhost:8000/ga4gh/wes/v1/service-info"
```

For debugging, you can get inside the container with:

```bash
docker exec -ti toil-agc-test /bin/bash
```

### Deploying

To push this to an Amazon ECR repo, where AGC can get at it, you can do something like:

```bash
AWS_REGION=<your-deployment-region> # For example, us-west-2
AWS_ACCOUNT=<your-account-number> # For example, 123456789012
ECR_REPO=<your-ecr-repo> # For example, yourname/toil-agc. Needs to be created in the ECR console.
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com
docker build -t ${ECR_REPO} .
docker tag ${ECR_REPO}:latest ${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_REPO}:latest
docker push ${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com/${ECR_REPO}:latest
```

### Using in AGC

To use a custom image in an AGC context, go to the directory for the project you want to use it with, set the environment variables, and deploy the context:

```bash
cd ../../../examples/demo-cwl-project/
export ECR_TOIL_ACCOUNT_ID="${AWS_ACCOUNT}"
export ECR_TOIL_TAG=latest
export ECR_TOIL_REPOSITORY="${ECR_REPO}"
export ECR_TOIL_REGION="${AWS_REGION}"
agc context deploy --context myContext
```
