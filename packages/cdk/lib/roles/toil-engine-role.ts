import { PolicyOptions } from "../types/engine-options";
import { BucketOperations } from "../common/BucketOperations";
import { ToilBatchPolicy } from "./policies/toil-batch-policy";
import { Arn, Aws, Stack } from "aws-cdk-lib";
import { Construct } from "constructs";
import { Role, ServicePrincipal, PolicyDocument, PolicyStatement, Effect } from "aws-cdk-lib/aws-iam";

interface ToilEngineRoleProps {
  readOnlyBucketArns: string[];
  readWriteBucketArns: string[];
  policies: PolicyOptions;
  // This is the queue to which we are authorizing jobs to be submitted by
  // something with this role.
  jobQueueArn: string;
}

export class ToilEngineRole extends Role {
  constructor(scope: Construct, id: string, props: ToilEngineRoleProps) {
    const toilJobArn = Arn.format(
      {
        account: Aws.ACCOUNT_ID,
        region: Aws.REGION,
        partition: Aws.PARTITION,
        resource: "job-definition/*",
        service: "batch",
      },
      scope as Stack
    );
    super(scope, id, {
      assumedBy: new ServicePrincipal("ecs-tasks.amazonaws.com"),
      inlinePolicies: {
        ToilEngineBatchPolicy: new ToilBatchPolicy({
          ...props,
          toilJobArn: toilJobArn,
        }),
        ToilEcsDescribeInstances: new PolicyDocument({
          assignSids: true,
          statements: [
            new PolicyStatement({
              effect: Effect.ALLOW,
              actions: ["ecs:DescribeContainerInstances", "s3:ListAllMyBuckets"],
              resources: ["*"],
            }),
          ],
        }),
        // TODO: Remove this when Toil no longer uses its own SimpleDB domains
        ToilSimpleDBFullAccess: new PolicyDocument({
          assignSids: true,
          statements: [
            new PolicyStatement({
              effect: Effect.ALLOW,
              actions: ["sdb:*"],
              resources: ["*"],
            }),
          ],
        }),
        // TODO: Remove this when Toil is taught to use AGC buckets to store
        // its workflow state and doesn't need to make and destroy its own.
        ToilS3FullAccess: new PolicyDocument({
          assignSids: true,
          statements: [
            new PolicyStatement({
              effect: Effect.ALLOW,
              actions: ["s3:*"],
              resources: ["*"],
            }),
          ],
        }),
      },
      ...props.policies,
    });

    BucketOperations.grantBucketAccess(this, this, props.readOnlyBucketArns, true);
    BucketOperations.grantBucketAccess(this, this, props.readWriteBucketArns);
  }
}
