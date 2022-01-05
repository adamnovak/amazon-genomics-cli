# AGC CDK

This library provides a set of constructs and resources used by AGC to deploy the
infrastructure needed to run genomics workflows.

To deploy these constructs manually:

1. Run `npm install` in this directory; AGC installs some needed modules elsewhere.
2. Go to the directory in `apps/` for the construct you want to deploy.
3. Run the CDK to deploy the construct. For example, for a Cromwell context, you could do:
```
npm run cdk -- -v -v -v "deploy" "--all" "--profile" "" "--require-approval" "never" "--output" "${PWD}/cdk-output-test" "-c" "READ_WRITE_BUCKET_ARNS=" "-c" "USER_ID=anovak4rBHiA" "-c" "ENGINE_NAME=cromwell" "-c" "ADAPTER_REPOSITORY=" "-c" "ADAPTER_NAME=" "-c" "AGC_VERSION=1.1.2" "-c" "ENGINE_REPOSITORY=" "-c" "ENGINE_HEALTH_CHECK_PATH=" "-c" "BATCH_COMPUTE_INSTANCE_TYPES=" "-c" "OUTPUT_BUCKET=agc-318423852362-us-west-2" "-c" "ENGINE_DESIGNATION=cromwell" "-c" "ADAPTER_DESIGNATION=" "-c" "ARTIFACT_BUCKET=healthai-public-assets-us-east-1" "-c" "READ_BUCKET_ARNS=arn:aws:s3:::gatk-test-data,arn:aws:s3:::broad-references" "-c" "PROJECT=Demo" "-c" "CONTEXT=myContext" "-c" "USER_EMAIL=anovak@soe.ucsc.edu" "-c" "MAX_V_CPUS=256" "-c" "REQUEST_SPOT_INSTANCES=false" "-c" "ECR_WES_ACCOUNT_ID=680431765560" "-c" "ECR_WES_REGION=us-west-2" "-c" "ECR_WES_TAG=0.1.0" "-c" "ECR_WES_REPOSITORY=aws/wes-release" "-c" "ECR_CROMWELL_ACCOUNT_ID=680431765560" "-c" "ECR_CROMWELL_REGION=us-west-2" "-c" "ECR_CROMWELL_TAG=64" "-c" "ECR_CROMWELL_REPOSITORY=aws/cromwell-mirror" "-c" "ECR_NEXTFLOW_ACCOUNT_ID=680431765560" "-c" "ECR_NEXTFLOW_REGION=us-west-2" "-c" "ECR_NEXTFLOW_TAG=21.04.3" "-c" "ECR_NEXTFLOW_REPOSITORY=aws/nextflow-mirror" "-c" "ECR_MINIWDL_ACCOUNT_ID=680431765560" "-c" "ECR_MINIWDL_REGION=us-west-2" "-c" "ECR_MINIWDL_TAG=v0.1.6" "-c" "ECR_MINIWDL_REPOSITORY=aws/miniwdl-mirror"
```
Make sure to set `OUTPUT_BUCKET`, `USER_ID`, and `USER_EMAIL` to values that make sense for you. You may also want to remove `"--require-approval" "never"` if you want to check the CDK's plan.

Note that since we now use CDK 2, which uses the `--platform` option on Docker, you will need Docker with API version 1.38 or later to deploy any contexts.
