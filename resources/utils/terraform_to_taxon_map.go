package utils

var terraformToTaxonName = map[string]struct {
	CloudServiceName string
}{
	"VPN (Site-to-Site)": {
		CloudServiceName: "AWS Site-to-Site VPN",
	},
	"WAF Classic Regional": {
		CloudServiceName: "AWS WAF Classic Regional",
	},
	"RAM (Resource Access Manager)": {
		CloudServiceName: "AWS Resource Access Manager",
	},
	"AppIntegrations": {
		CloudServiceName: "Amazon AppIntegrations Service",
	},
	"Cognito IDP (Identity Provider)": {
		CloudServiceName: "Amazon Cognito",
	},
	"CloudWatch Internet Monitor": {
		CloudServiceName: "Amazon CloudWatch Internet Monitor",
	},
	"WAF Classic": {
		CloudServiceName: "AWS WAF Classic",
	},
	"API Gateway": {
		CloudServiceName: "Amazon API Gateway",
	},
	"IAM (Identity & Access Management)": {
		CloudServiceName: "Amazon Identity and Access Management",
	},
	"SNS (Simple Notification)": {
		CloudServiceName: "Amazon Simple Notification Service",
	},
	"Storage Gateway": {
		CloudServiceName: "AWS Storage Gateway",
	},
	"Verified Permissions": {
		CloudServiceName: "Amazon Verified Permissions",
	},
	"VPC (Virtual Private Cloud)": {
		CloudServiceName: "Amazon Virtual Private Cloud",
	},
	"Elastic Beanstalk": {
		CloudServiceName: "AWS Elastic Beanstalk",
	},
	"IoT Core": {
		CloudServiceName: "AWS IoT Core",
	},
	"Kendra": {
		CloudServiceName: "Amazon Kendra",
	},
	"RDS (Relational Database)": {
		CloudServiceName: "Amazon Relational Database Service",
	},
	"SSO Admin": {
		CloudServiceName: "IAM Identity Center",
	},
	"AppStream 2.0": {
		CloudServiceName: "Amazon AppStream 2.0",
	},
	"CloudWatch": {
		CloudServiceName: "Amazon CloudWatch",
	},
	"DMS (Database Migration)": {
		CloudServiceName: "AWS Database Migration Service",
	},
	"Global Accelerator": {
		CloudServiceName: "AWS Global Accelerator",
	},
	"Glue": {
		CloudServiceName: "AWS Glue",
	},
	"Inspector Classic": {
		CloudServiceName: "Amazon Inspector Classic",
	},
	"Batch": {
		CloudServiceName: "AWS Batch",
	},
	"Auto Scaling": {
		CloudServiceName: "AWS Auto Scaling",
	},
	"S3 Control": {
		CloudServiceName: "Amazon S3 Control",
	},
	"SESv2 (Simple Email V2)": {
		CloudServiceName: "Amazon Simple Email Service",
	},
	"Detective": {
		CloudServiceName: "Amazon Detective",
	},
	"VPC IPAM (IP Address Manager)": {
		CloudServiceName: "Amazon VPC IP Address Manager (IPAM)",
	},
	"EMR Serverless": {
		CloudServiceName: "Amazon EMR Serverless",
	},
	"Backup": {
		CloudServiceName: "AWS Backup",
	},
	"Connect": {
		CloudServiceName: "Amazon Connect",
	},
	"SageMaker": {
		CloudServiceName: "Amazon SageMaker",
	},
	"SSM (Systems Manager)": {
		CloudServiceName: "AWS Systems Manager",
	},
	"EFS (Elastic File System)": {
		CloudServiceName: "Amazon Elastic File System",
	},
	"Kinesis": {
		CloudServiceName: "Amazon Kinesis",
	},
	"Service Catalog": {
		CloudServiceName: "AWS Service Catalog",
	},
	"DynamoDB": {
		CloudServiceName: "Amazon DynamoDB",
	},
	"ElastiCache": {
		CloudServiceName: "Amazon ElastiCache",
	},
	"FIS (Fault Injection Simulator)": {
		CloudServiceName: "AWS Fault Injection Service",
	},
	"Lambda": {
		CloudServiceName: "AWS Lambda",
	},
	"Lex Model Building": {
		CloudServiceName: "Amazon Lex",
	},
	"Redshift": {
		CloudServiceName: "Amazon Redshift",
	},
	"App Runner": {
		CloudServiceName: "AWS App Runner",
	},
	"ELB Classic": {
		CloudServiceName: "Elastic Load Balancing",
	},
	"Managed Streaming for Kafka": {
		CloudServiceName: "Amazon Managed Streaming for Apache Kafka",
	},
	"EC2 (Elastic Compute Cloud)": {
		CloudServiceName: "Amazon Elastic Compute Cloud (EC2)",
	},
	"Elemental MediaLive": {
		CloudServiceName: "AWS Elemental MediaLive",
	},
	"Route 53": {
		CloudServiceName: "Amazon Route 53",
	},
	"Chime SDK Media Pipelines": {
		CloudServiceName: "Amazon Chime SDK media pipelines",
	},
	"S3 Glacier": {
		CloudServiceName: "Amazon S3 Glacier",
	},
	"Neptune": {
		CloudServiceName: "Amazon Neptune",
	},
	"Redshift Serverless": {
		CloudServiceName: "Amazon Redshift Serverless",
	},
	"Security Hub": {
		CloudServiceName: "AWS Security Hub",
	},
	"WAF": {
		CloudServiceName: "AWS WAF",
	},
	"ELB (Elastic Load Balancing)": {
		CloudServiceName: "Elastic Load Balancing",
	},
	"DynamoDB Accelerator (DAX)": {
		CloudServiceName: "Amazon DynamoDB Accelerator (DAX)",
	},
	"Device Farm": {
		CloudServiceName: "AWS Device Farm",
	},
	"Pinpoint": {
		CloudServiceName: "Amazon Pinpoint",
	},
	"App Mesh": {
		CloudServiceName: "AWS App Mesh",
	},
	"CloudSearch": {
		CloudServiceName: "Amazon CloudSearch",
	},
	"Lightsail": {
		CloudServiceName: "Amazon Lightsail",
	},
}
