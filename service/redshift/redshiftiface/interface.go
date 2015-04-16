package redshiftiface

import (
	"github.com/awslabs/aws-sdk-go/service/redshift"
)

type RedshiftAPI interface {
	AuthorizeClusterSecurityGroupIngress(*redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error)

	AuthorizeSnapshotAccess(*redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error)

	CopyClusterSnapshot(*redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error)

	CreateCluster(*redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error)

	CreateClusterParameterGroup(*redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error)

	CreateClusterSecurityGroup(*redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error)

	CreateClusterSnapshot(*redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error)

	CreateClusterSubnetGroup(*redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error)

	CreateEventSubscription(*redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error)

	CreateHSMClientCertificate(*redshift.CreateHSMClientCertificateInput) (*redshift.CreateHSMClientCertificateOutput, error)

	CreateHSMConfiguration(*redshift.CreateHSMConfigurationInput) (*redshift.CreateHSMConfigurationOutput, error)

	CreateTags(*redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error)

	DeleteCluster(*redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error)

	DeleteClusterParameterGroup(*redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error)

	DeleteClusterSecurityGroup(*redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error)

	DeleteClusterSnapshot(*redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error)

	DeleteClusterSubnetGroup(*redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error)

	DeleteEventSubscription(*redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error)

	DeleteHSMClientCertificate(*redshift.DeleteHSMClientCertificateInput) (*redshift.DeleteHSMClientCertificateOutput, error)

	DeleteHSMConfiguration(*redshift.DeleteHSMConfigurationInput) (*redshift.DeleteHSMConfigurationOutput, error)

	DeleteTags(*redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error)

	DescribeClusterParameterGroups(*redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error)

	DescribeClusterParameters(*redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error)

	DescribeClusterSecurityGroups(*redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error)

	DescribeClusterSnapshots(*redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error)

	DescribeClusterSubnetGroups(*redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error)

	DescribeClusterVersions(*redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error)

	DescribeClusters(*redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error)

	DescribeDefaultClusterParameters(*redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error)

	DescribeEventCategories(*redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error)

	DescribeEventSubscriptions(*redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error)

	DescribeEvents(*redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error)

	DescribeHSMClientCertificates(*redshift.DescribeHSMClientCertificatesInput) (*redshift.DescribeHSMClientCertificatesOutput, error)

	DescribeHSMConfigurations(*redshift.DescribeHSMConfigurationsInput) (*redshift.DescribeHSMConfigurationsOutput, error)

	DescribeLoggingStatus(*redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error)

	DescribeOrderableClusterOptions(*redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error)

	DescribeReservedNodeOfferings(*redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error)

	DescribeReservedNodes(*redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error)

	DescribeResize(*redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error)

	DescribeTags(*redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error)

	DisableLogging(*redshift.DisableLoggingInput) (*redshift.LoggingStatus, error)

	DisableSnapshotCopy(*redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error)

	EnableLogging(*redshift.EnableLoggingInput) (*redshift.LoggingStatus, error)

	EnableSnapshotCopy(*redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error)

	ModifyCluster(*redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error)

	ModifyClusterParameterGroup(*redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)

	ModifyClusterSubnetGroup(*redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error)

	ModifyEventSubscription(*redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error)

	ModifySnapshotCopyRetentionPeriod(*redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error)

	PurchaseReservedNodeOffering(*redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error)

	RebootCluster(*redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error)

	ResetClusterParameterGroup(*redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)

	RestoreFromClusterSnapshot(*redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error)

	RevokeClusterSecurityGroupIngress(*redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error)

	RevokeSnapshotAccess(*redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error)

	RotateEncryptionKey(*redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error)
}