#!/usr/bin/env node
import codebuild = require('@aws-cdk/aws-codebuild');
import codepipeline = require('@aws-cdk/aws-codepipeline');
import actions = require('@aws-cdk/aws-codepipeline-actions');
import iam = require('@aws-cdk/aws-iam');
import cdk = require('@aws-cdk/core');

/**
 * Simple three-stage pipeline to build the base image for the local container endpoints image.
 * [GitHub source] -> [CodeBuild build, pushes image to DockerHub] -> [CodeBuild Verify, verifies the pushed images]
 *
 * TODO: use Docker manifest to create manifest for images under a common tag and ECR public
 */
class EcsLocalContainerEndpointsImagePipeline extends cdk.Stack {
  constructor(parent: cdk.App, name: string, props?: cdk.StackProps) {
    super(parent, name, props);

    // Instantiate pipeline
    const pipeline = new codepipeline.Pipeline(this, 'Pipeline', {
      pipelineName: 'local-container-endpoints-image',
    });

    // Source stage
    // Secret under ecs-local-container-endpoints+release@amazon.com
    const githubAccessToken = cdk.SecretValue.secretsManager('EcsDevXGitHubToken');

    const sourceOutput = new codepipeline.Artifact('SourceArtifact');
    const sourceAction = new actions.GitHubSourceAction({
      actionName: 'GitHubSource',
      owner: 'SoManyHs',
      repo: 'amazon-ecs-local-container-endpoints',
      oauthToken: githubAccessToken,
      branch: 'mainline',
      output: sourceOutput
    });

    pipeline.addStage({
      stageName: 'Source',
      actions: [sourceAction],
    });

    // Build stage containing build project for each architecture image
    const buildStage = pipeline.addStage({
      stageName: 'Build',
    });

    const verifyStage = pipeline.addStage({
      stageName: 'Verify',
    });

    const platforms = [
      {'arch': 'amd64', 'buildImage': codebuild.LinuxBuildImage.AMAZON_LINUX_2_3},
      {'arch': 'arm64', 'buildImage': codebuild.LinuxBuildImage.AMAZON_LINUX_2_ARM},
    ];

    // Create build and verify project for each platform
    for (const platform of platforms) {
      const arch = platform['arch'];

      const buildProject = new codebuild.PipelineProject(this, `BuildImage-${arch}`, {
        buildSpec: codebuild.BuildSpec.fromSourceFilename('./buildspec.yml'),
        environment: {
          buildImage: platform['buildImage'],
          privileged: true,
          environmentVariables: {
            ARCH_SUFFIX: { value: arch },
          }
        }
      });

      const verifyProject = new codebuild.PipelineProject(this, `VerifyImage-${arch}`, {
        buildSpec: codebuild.BuildSpec.fromSourceFilename('./buildspec_verify.yml'),
        environment: {
          buildImage: platform['buildImage'],
          privileged: true,
          environmentVariables: {
            ARCH_SUFFIX: { value: arch },
          }
        }
      });

      buildProject.addToRolePolicy(new iam.PolicyStatement({
        actions: ["ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:GetRepositoryPolicy",
          "ecr:DescribeRepositories",
          "ecr:ListImages",
          "ecr:DescribeImages",
          "ecr:BatchGetImage",
          "ecr:InitiateLayerUpload",
          "ecr:UploadLayerPart",
          "ecr:CompleteLayerUpload",
          "ecr:PutImage",
          "ecr-public:*",
          "secretsmanager:GetSecretValue",
          "sts:GetServiceBearerToken",
          "sts:AssumeRole",
        ],
        resources: ["*"]
      }));

      verifyProject.addToRolePolicy(new iam.PolicyStatement({
        actions: [
          "secretsmanager:GetSecretValue",
        ],
        resources: ["com.amazonaws.ec2.madison.dockerhub.amazon-ecs-local-container-endpoints.credentials"]
      }));

      const buildAction = new actions.CodeBuildAction({
        actionName: `Build-${platform['arch']}`,
        project: buildProject,
        input: sourceOutput
      });

      const verifyAction = new actions.CodeBuildAction({
        actionName: `Verify-${platform['arch']}`,
        project: verifyProject,
        input: sourceOutput
      });

      // Add build action for each platform to the build stage
      buildStage.addAction(buildAction);
      verifyStage.addAction(verifyAction);
    }
  }

  // TODO Build stage for creating manifest
}

const app = new cdk.App();

new EcsLocalContainerEndpointsImagePipeline(app, 'EcsLocalContainerEndpointsImagePipeline', {
  env: {
    account: process.env['CDK_DEFAULT_ACCOUNT'],
    region: 'us-west-2'
  },
  tags: {
    project: "amazon-ecs-local-container-endpoints"
  }
});
app.synth();
