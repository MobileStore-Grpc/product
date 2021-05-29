#!/bin/bash

EKS_ROLE_ASSUME_POLICY=$(echo -n '{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowAssumeEksWorkshopCodeBuildKubectlRole",
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "arn:aws:iam::890358416518:role/EksWorkshopCodeBuildKubectlRole"
    },
    {
        "Effect": "Allow",
        "Action": "codestar-connections:UseConnection",
        "Resource": "arn:aws:codestar-connections:us-east-1:890358416518:connection/f0db2dbb-e1b2-4351-8171-b3354930b326"
    },
    {
    "Action": [
        "appconfig:StartDeployment",
        "appconfig:GetDeployment",
        "appconfig:StopDeployment"
    ],
    "Resource": "*",
    "Effect": "Allow"
    },
    {
    "Action": [
        "codecommit:GetRepository"
    ],
    "Resource": "*",
    "Effect": "Allow"
    }
  ]
}')


TRUST="{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Effect\": \"Allow\", \"Principal\": { \"Service\": \"codepipeline.amazonaws.com\" }, \"Action\": \"sts:AssumeRole\" } ] }"

echo EKS_ROLE_ASSUME_POLICY=$EKS_ROLE_ASSUME_POLICY

aws iam create-role --role-name codebuild-eks-devops-cb-for-pipe-service-role --assume-role-policy-document "$TRUST" --output text --query 'Role.Arn'

aws iam put-role-policy --role-name codebuild-eks-devops-cb-for-pipe-service-role --policy-name assume-codebuild-kubectlrole --policy-document "$EKS_ROLE_ASSUME_POLICY"
aws iam attach-role-policy --role-name codebuild-eks-devops-cb-for-pipe-service-role --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess