// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translator

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/model/pdata"
	semconventions "go.opentelemetry.io/collector/translator/conventions"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func TestAwsFromEc2Resource(t *testing.T) {
	instanceID := "i-00f7c0bcb26da2a99"
	hostType := "m5.xlarge"
	imageID := "ami-0123456789"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudPlatform, semconventions.AttributeCloudPlatformAWSEC2)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, "us-east-1c")
	attrs.InsertString(semconventions.AttributeHostID, instanceID)
	attrs.InsertString(semconventions.AttributeHostType, hostType)
	attrs.InsertString(semconventions.AttributeHostImageID, imageID)
	attrs.CopyTo(resource.Attributes())

	attributes := make(map[string]pdata.AttributeValue)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.NotNil(t, awsData.EC2)
	assert.Nil(t, awsData.ECS)
	assert.Nil(t, awsData.Beanstalk)
	assert.Nil(t, awsData.EKS)
	assert.Equal(t, "123456789", *awsData.AccountID)
	assert.Equal(t, &awsxray.EC2Metadata{
		InstanceID:       aws.String(instanceID),
		AvailabilityZone: aws.String("us-east-1c"),
		InstanceSize:     aws.String(hostType),
		AmiID:            aws.String(imageID),
	}, awsData.EC2)
}

func TestAwsFromEcsResource(t *testing.T) {
	instanceID := "i-00f7c0bcb26da2a99"
	containerName := "signup_aggregator-x82ufje83"
	containerID := "0123456789A"
	az := "us-east-1c"
	launchType := "fargate"
	family := "family"
	taskArn := "arn:aws:ecs:us-west-2:123456789123:task/123"
	clusterArn := "arn:aws:ecs:us-west-2:123456789123:cluster/my-cluster"
	containerArn := "arn:aws:ecs:us-west-2:123456789123:container-instance/123"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudPlatform, semconventions.AttributeCloudPlatformAWSECS)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, az)
	attrs.InsertString(semconventions.AttributeContainerImage, "otel/signupaggregator")
	attrs.InsertString(semconventions.AttributeContainerTag, "v1")
	attrs.InsertString(semconventions.AttributeContainerName, containerName)
	attrs.InsertString(semconventions.AttributeContainerID, containerID)
	attrs.InsertString(semconventions.AttributeHostID, instanceID)
	attrs.InsertString(semconventions.AttributeAWSECSClusterARN, clusterArn)
	attrs.InsertString(semconventions.AttributeAWSECSContainerARN, containerArn)
	attrs.InsertString(semconventions.AttributeAWSECSTaskARN, taskArn)
	attrs.InsertString(semconventions.AttributeAWSECSTaskFamily, family)
	attrs.InsertString(semconventions.AttributeAWSECSLaunchType, launchType)
	attrs.InsertString(semconventions.AttributeHostType, "m5.xlarge")

	attrs.CopyTo(resource.Attributes())

	attributes := make(map[string]pdata.AttributeValue)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.NotNil(t, awsData.ECS)
	assert.Nil(t, awsData.EC2)
	assert.Nil(t, awsData.Beanstalk)
	assert.Nil(t, awsData.EKS)
	assert.Equal(t, &awsxray.ECSMetadata{
		ContainerName:    aws.String(containerName),
		ContainerID:      aws.String(containerID),
		AvailabilityZone: aws.String(az),
		ClusterArn:       aws.String(clusterArn),
		ContainerArn:     aws.String(containerArn),
		TaskArn:          aws.String(taskArn),
		TaskFamily:       aws.String(family),
		LaunchType:       aws.String(launchType),
	}, awsData.ECS)
}

func TestAwsFromBeanstalkResource(t *testing.T) {
	deployID := "232"
	versionLabel := "4"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudPlatform, semconventions.AttributeCloudPlatformAWSElasticBeanstalk)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, "us-east-1c")
	attrs.InsertString(semconventions.AttributeServiceNamespace, "production")
	attrs.InsertString(semconventions.AttributeServiceInstance, deployID)
	attrs.InsertString(semconventions.AttributeServiceVersion, versionLabel)
	attrs.CopyTo(resource.Attributes())

	attributes := make(map[string]pdata.AttributeValue)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Nil(t, awsData.EC2)
	assert.Nil(t, awsData.ECS)
	assert.NotNil(t, awsData.Beanstalk)
	assert.Nil(t, awsData.EKS)
	assert.Equal(t, &awsxray.BeanstalkMetadata{
		Environment:  aws.String("production"),
		VersionLabel: aws.String(versionLabel),
		DeploymentID: aws.Int64(232),
	}, awsData.Beanstalk)
}

func TestAwsFromEksResource(t *testing.T) {
	instanceID := "i-00f7c0bcb26da2a99"
	containerName := "signup_aggregator-x82ufje83"
	containerID := "0123456789A"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudPlatform, semconventions.AttributeCloudPlatformAWSEKS)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, "us-east-1c")
	attrs.InsertString(semconventions.AttributeContainerImage, "otel/signupaggregator")
	attrs.InsertString(semconventions.AttributeContainerTag, "v1")
	attrs.InsertString(semconventions.AttributeK8sCluster, "production")
	attrs.InsertString(semconventions.AttributeK8sNamespace, "default")
	attrs.InsertString(semconventions.AttributeK8sDeployment, "signup_aggregator")
	attrs.InsertString(semconventions.AttributeK8sPod, "my-deployment-65dcf7d447-ddjnl")
	attrs.InsertString(semconventions.AttributeContainerName, containerName)
	attrs.InsertString(semconventions.AttributeContainerID, containerID)
	attrs.InsertString(semconventions.AttributeHostID, instanceID)
	attrs.InsertString(semconventions.AttributeHostType, "m5.xlarge")
	attrs.CopyTo(resource.Attributes())

	attributes := make(map[string]pdata.AttributeValue)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.NotNil(t, awsData.EKS)
	assert.Nil(t, awsData.EC2)
	assert.Nil(t, awsData.ECS)
	assert.Nil(t, awsData.Beanstalk)
	assert.Equal(t, &awsxray.EKSMetadata{
		ClusterName: aws.String("production"),
		Pod:         aws.String("my-deployment-65dcf7d447-ddjnl"),
		ContainerID: aws.String(containerID),
	}, awsData.EKS)
}

func TestAwsWithAwsSqsResources(t *testing.T) {
	instanceID := "i-00f7c0bcb26da2a99"
	containerName := "signup_aggregator-x82ufje83"
	containerID := "0123456789A"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, "us-east-1c")
	attrs.InsertString(semconventions.AttributeContainerName, containerName)
	attrs.InsertString(semconventions.AttributeContainerImage, "otel/signupaggregator")
	attrs.InsertString(semconventions.AttributeContainerTag, "v1")
	attrs.InsertString(semconventions.AttributeK8sCluster, "production")
	attrs.InsertString(semconventions.AttributeK8sNamespace, "default")
	attrs.InsertString(semconventions.AttributeK8sDeployment, "signup_aggregator")
	attrs.InsertString(semconventions.AttributeK8sPod, "my-deployment-65dcf7d447-ddjnl")
	attrs.InsertString(semconventions.AttributeContainerName, containerName)
	attrs.InsertString(semconventions.AttributeContainerID, containerID)
	attrs.InsertString(semconventions.AttributeHostID, instanceID)
	attrs.InsertString(semconventions.AttributeHostType, "m5.xlarge")

	queueURL := "https://sqs.use1.amazonaws.com/Meltdown-Alerts"
	attributes := make(map[string]pdata.AttributeValue)
	attributes[awsxray.AWSOperationAttribute] = pdata.NewAttributeValueString("SendMessage")
	attributes[awsxray.AWSAccountAttribute] = pdata.NewAttributeValueString("987654321")
	attributes[awsxray.AWSRegionAttribute] = pdata.NewAttributeValueString("us-east-2")
	attributes[awsxray.AWSQueueURLAttribute] = pdata.NewAttributeValueString(queueURL)
	attributes["employee.id"] = pdata.NewAttributeValueString("XB477")

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, queueURL, *awsData.QueueURL)
	assert.Equal(t, "us-east-2", *awsData.RemoteRegion)
}

func TestAwsWithSqsAlternateAttribute(t *testing.T) {
	queueURL := "https://sqs.use1.amazonaws.com/Meltdown-Alerts"
	attributes := make(map[string]pdata.AttributeValue)
	attributes[awsxray.AWSQueueURLAttribute2] = pdata.NewAttributeValueString(queueURL)

	filtered, awsData := makeAws(attributes, pdata.NewResource())

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, queueURL, *awsData.QueueURL)
}

func TestAwsWithAwsDynamoDbResources(t *testing.T) {
	instanceID := "i-00f7c0bcb26da2a99"
	containerName := "signup_aggregator-x82ufje83"
	containerID := "0123456789A"
	resource := pdata.NewResource()
	attrs := pdata.NewAttributeMap()
	attrs.InsertString(semconventions.AttributeCloudProvider, semconventions.AttributeCloudProviderAWS)
	attrs.InsertString(semconventions.AttributeCloudAccount, "123456789")
	attrs.InsertString(semconventions.AttributeCloudAvailabilityZone, "us-east-1c")
	attrs.InsertString(semconventions.AttributeContainerName, "signup_aggregator")
	attrs.InsertString(semconventions.AttributeContainerImage, "otel/signupaggregator")
	attrs.InsertString(semconventions.AttributeContainerTag, "v1")
	attrs.InsertString(semconventions.AttributeK8sCluster, "production")
	attrs.InsertString(semconventions.AttributeK8sNamespace, "default")
	attrs.InsertString(semconventions.AttributeK8sDeployment, "signup_aggregator")
	attrs.InsertString(semconventions.AttributeK8sPod, "my-deployment-65dcf7d447-ddjnl")
	attrs.InsertString(semconventions.AttributeContainerName, containerName)
	attrs.InsertString(semconventions.AttributeContainerID, containerID)
	attrs.InsertString(semconventions.AttributeHostID, instanceID)
	attrs.InsertString(semconventions.AttributeHostType, "m5.xlarge")

	tableName := "WIDGET_TYPES"
	attributes := make(map[string]pdata.AttributeValue)
	attributes[awsxray.AWSOperationAttribute] = pdata.NewAttributeValueString("PutItem")
	attributes[awsxray.AWSRequestIDAttribute] = pdata.NewAttributeValueString("75107C82-EC8A-4F75-883F-4440B491B0AB")
	attributes[awsxray.AWSTableNameAttribute] = pdata.NewAttributeValueString(tableName)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, "PutItem", *awsData.Operation)
	assert.Equal(t, "75107C82-EC8A-4F75-883F-4440B491B0AB", *awsData.RequestID)
	assert.Equal(t, tableName, *awsData.TableName)
}

func TestAwsWithDynamoDbAlternateAttribute(t *testing.T) {
	tableName := "MyTable"
	attributes := make(map[string]pdata.AttributeValue)
	attributes[awsxray.AWSTableNameAttribute2] = pdata.NewAttributeValueString(tableName)

	filtered, awsData := makeAws(attributes, pdata.NewResource())

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, tableName, *awsData.TableName)
}

func TestAwsWithRequestIdAlternateAttribute(t *testing.T) {
	requestid := "12345-request"
	attributes := make(map[string]pdata.AttributeValue)
	attributes[awsxray.AWSRequestIDAttribute2] = pdata.NewAttributeValueString(requestid)

	filtered, awsData := makeAws(attributes, pdata.NewResource())

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, requestid, *awsData.RequestID)
}

func TestJavaSDK(t *testing.T) {
	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKName, "opentelemetry")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKLanguage, "java")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKVersion, "1.2.3")

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, "opentelemetry for java", *awsData.XRay.SDK)
	assert.Equal(t, "1.2.3", *awsData.XRay.SDKVersion)
}

func TestJavaAutoInstrumentation(t *testing.T) {
	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKName, "opentelemetry")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKLanguage, "java")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKVersion, "1.2.3")
	resource.Attributes().InsertString(semconventions.AttributeTelemetryAutoVersion, "3.4.5")

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, "opentelemetry for java", *awsData.XRay.SDK)
	assert.Equal(t, "1.2.3", *awsData.XRay.SDKVersion)
	assert.True(t, *awsData.XRay.AutoInstrumentation)
}

func TestGoSDK(t *testing.T) {
	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKName, "opentelemetry")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKLanguage, "go")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKVersion, "2.0.3")

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, "opentelemetry for go", *awsData.XRay.SDK)
	assert.Equal(t, "2.0.3", *awsData.XRay.SDKVersion)
}

func TestCustomSDK(t *testing.T) {
	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKName, "opentracing")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKLanguage, "java")
	resource.Attributes().InsertString(semconventions.AttributeTelemetrySDKVersion, "2.0.3")

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, "opentracing for java", *awsData.XRay.SDK)
	assert.Equal(t, "2.0.3", *awsData.XRay.SDKVersion)
}

func TestLogGroups(t *testing.T) {
	cwl1 := awsxray.LogGroupMetadata{
		LogGroup: awsxray.String("group1"),
	}
	cwl2 := awsxray.LogGroupMetadata{
		LogGroup: awsxray.String("group2"),
	}

	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	lg := pdata.NewAttributeValueArray()
	ava := lg.ArrayVal()
	ava.EnsureCapacity(2)
	ava.AppendEmpty().SetStringVal("group1")
	ava.AppendEmpty().SetStringVal("group2")

	resource.Attributes().Insert(semconventions.AttributeAWSLogGroupNames, lg)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, 2, len(awsData.CWLogs))
	assert.Contains(t, awsData.CWLogs, cwl1)
	assert.Contains(t, awsData.CWLogs, cwl2)
}

func TestLogGroupsFromArns(t *testing.T) {
	group1 := "arn:aws:logs:us-east-1:123456789123:log-group:group1"
	cwl1 := awsxray.LogGroupMetadata{
		LogGroup: awsxray.String("group1"),
		Arn:      awsxray.String(group1),
	}
	group2 := "arn:aws:logs:us-east-1:123456789123:log-group:group2"
	cwl2 := awsxray.LogGroupMetadata{
		LogGroup: awsxray.String("group2"),
		Arn:      awsxray.String(group2),
	}

	attributes := make(map[string]pdata.AttributeValue)
	resource := pdata.NewResource()
	lga := pdata.NewAttributeValueArray()
	ava := lga.ArrayVal()
	ava.EnsureCapacity(2)
	ava.AppendEmpty().SetStringVal(group1)
	ava.AppendEmpty().SetStringVal(group2)

	resource.Attributes().Insert(semconventions.AttributeAWSLogGroupARNs, lga)

	filtered, awsData := makeAws(attributes, resource)

	assert.NotNil(t, filtered)
	assert.NotNil(t, awsData)
	assert.Equal(t, 2, len(awsData.CWLogs))
	assert.Contains(t, awsData.CWLogs, cwl1)
	assert.Contains(t, awsData.CWLogs, cwl2)
}
