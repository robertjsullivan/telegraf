// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Output parameter of the GetRecords API. The existing child shard of the current
// shard.
type ChildShard struct {

	// The range of possible hash key values for the shard, which is a set of ordered
	// contiguous positive integers.
	//
	// This member is required.
	HashKeyRange *HashKeyRange

	// The current shard that is the parent of the existing child shard.
	//
	// This member is required.
	ParentShards []string

	// The shard ID of the existing child shard of the current shard.
	//
	// This member is required.
	ShardId *string

	noSmithyDocumentSerde
}

// An object that represents the details of the consumer you registered. This type
// of object is returned by RegisterStreamConsumer.
type Consumer struct {

	// When you register a consumer, Kinesis Data Streams generates an ARN for it. You
	// need this ARN to be able to call SubscribeToShard. If you delete a consumer and
	// then create a new one with the same name, it won't have the same ARN. That's
	// because consumer ARNs contain the creation timestamp. This is important to keep
	// in mind if you have IAM policies that reference consumer ARNs.
	//
	// This member is required.
	ConsumerARN *string

	//
	//
	// This member is required.
	ConsumerCreationTimestamp *time.Time

	// The name of the consumer is something you choose when you register the consumer.
	//
	// This member is required.
	ConsumerName *string

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// This member is required.
	ConsumerStatus ConsumerStatus

	noSmithyDocumentSerde
}

// An object that represents the details of a registered consumer. This type of
// object is returned by DescribeStreamConsumer.
type ConsumerDescription struct {

	// When you register a consumer, Kinesis Data Streams generates an ARN for it. You
	// need this ARN to be able to call SubscribeToShard. If you delete a consumer and
	// then create a new one with the same name, it won't have the same ARN. That's
	// because consumer ARNs contain the creation timestamp. This is important to keep
	// in mind if you have IAM policies that reference consumer ARNs.
	//
	// This member is required.
	ConsumerARN *string

	//
	//
	// This member is required.
	ConsumerCreationTimestamp *time.Time

	// The name of the consumer is something you choose when you register the consumer.
	//
	// This member is required.
	ConsumerName *string

	// A consumer can't read data while in the CREATING or DELETING states.
	//
	// This member is required.
	ConsumerStatus ConsumerStatus

	// The ARN of the stream with which you registered the consumer.
	//
	// This member is required.
	StreamARN *string

	noSmithyDocumentSerde
}

// Represents enhanced metrics types.
type EnhancedMetrics struct {

	// List of shard-level metrics. The following are the valid shard-level metrics.
	// The value "ALL" enhances every metric.
	//
	// * IncomingBytes
	//
	// * IncomingRecords
	//
	// *
	// OutgoingBytes
	//
	// * OutgoingRecords
	//
	// * WriteProvisionedThroughputExceeded
	//
	// *
	// ReadProvisionedThroughputExceeded
	//
	// * IteratorAgeMilliseconds
	//
	// * ALL
	//
	// For more
	// information, see Monitoring the Amazon Kinesis Data Streams Service with Amazon
	// CloudWatch
	// (https://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	ShardLevelMetrics []MetricsName

	noSmithyDocumentSerde
}

// The range of possible hash key values for the shard, which is a set of ordered
// contiguous positive integers.
type HashKeyRange struct {

	// The ending hash key of the hash key range.
	//
	// This member is required.
	EndingHashKey *string

	// The starting hash key of the hash key range.
	//
	// This member is required.
	StartingHashKey *string

	noSmithyDocumentSerde
}

// Represents the output for PutRecords.
type PutRecordsRequestEntry struct {

	// The data blob to put into the record, which is base64-encoded when the blob is
	// serialized. When the data blob (the payload before base64-encoding) is added to
	// the partition key size, the total size must not exceed the maximum record size
	// (1 MiB).
	//
	// This member is required.
	Data []byte

	// Determines which shard in the stream the data record is assigned to. Partition
	// keys are Unicode strings with a maximum length limit of 256 characters for each
	// key. Amazon Kinesis Data Streams uses the partition key as input to a hash
	// function that maps the partition key and associated data to a specific shard.
	// Specifically, an MD5 hash function is used to map partition keys to 128-bit
	// integer values and to map associated data records to shards. As a result of this
	// hashing mechanism, all data records with the same partition key map to the same
	// shard within the stream.
	//
	// This member is required.
	PartitionKey *string

	// The hash value used to determine explicitly the shard that the data record is
	// assigned to by overriding the partition key hash.
	ExplicitHashKey *string

	noSmithyDocumentSerde
}

// Represents the result of an individual record from a PutRecords request. A
// record that is successfully added to a stream includes SequenceNumber and
// ShardId in the result. A record that fails to be added to the stream includes
// ErrorCode and ErrorMessage in the result.
type PutRecordsResultEntry struct {

	// The error code for an individual record result. ErrorCodes can be either
	// ProvisionedThroughputExceededException or InternalFailure.
	ErrorCode *string

	// The error message for an individual record result. An ErrorCode value of
	// ProvisionedThroughputExceededException has an error message that includes the
	// account ID, stream name, and shard ID. An ErrorCode value of InternalFailure has
	// the error message "Internal Service Failure".
	ErrorMessage *string

	// The sequence number for an individual record result.
	SequenceNumber *string

	// The shard ID for an individual record result.
	ShardId *string

	noSmithyDocumentSerde
}

// The unit of data of the Kinesis data stream, which is composed of a sequence
// number, a partition key, and a data blob.
type Record struct {

	// The data blob. The data in the blob is both opaque and immutable to Kinesis Data
	// Streams, which does not inspect, interpret, or change the data in the blob in
	// any way. When the data blob (the payload before base64-encoding) is added to the
	// partition key size, the total size must not exceed the maximum record size (1
	// MiB).
	//
	// This member is required.
	Data []byte

	// Identifies which shard in the stream the data record is assigned to.
	//
	// This member is required.
	PartitionKey *string

	// The unique identifier of the record within its shard.
	//
	// This member is required.
	SequenceNumber *string

	// The approximate time that the record was inserted into the stream.
	ApproximateArrivalTimestamp *time.Time

	// The encryption type used on the record. This parameter can be one of the
	// following values:
	//
	// * NONE: Do not encrypt the records in the stream.
	//
	// * KMS: Use
	// server-side encryption on the records in the stream using a customer-managed
	// Amazon Web Services KMS key.
	EncryptionType EncryptionType

	noSmithyDocumentSerde
}

// The range of possible sequence numbers for the shard.
type SequenceNumberRange struct {

	// The starting sequence number for the range.
	//
	// This member is required.
	StartingSequenceNumber *string

	// The ending sequence number for the range. Shards that are in the OPEN state have
	// an ending sequence number of null.
	EndingSequenceNumber *string

	noSmithyDocumentSerde
}

// A uniquely identified group of data records in a Kinesis data stream.
type Shard struct {

	// The range of possible hash key values for the shard, which is a set of ordered
	// contiguous positive integers.
	//
	// This member is required.
	HashKeyRange *HashKeyRange

	// The range of possible sequence numbers for the shard.
	//
	// This member is required.
	SequenceNumberRange *SequenceNumberRange

	// The unique identifier of the shard within the stream.
	//
	// This member is required.
	ShardId *string

	// The shard ID of the shard adjacent to the shard's parent.
	AdjacentParentShardId *string

	// The shard ID of the shard's parent.
	ParentShardId *string

	noSmithyDocumentSerde
}

// The request parameter used to filter out the response of the ListShards API.
type ShardFilter struct {

	// The shard type specified in the ShardFilter parameter. This is a required
	// property of the ShardFilter parameter. You can specify the following valid
	// values:
	//
	// * AFTER_SHARD_ID - the response includes all the shards, starting with
	// the shard whose ID immediately follows the ShardId that you provided.
	//
	// *
	// AT_TRIM_HORIZON - the response includes all the shards that were open at
	// TRIM_HORIZON.
	//
	// * FROM_TRIM_HORIZON - (default), the response includes all the
	// shards within the retention period of the data stream (trim to tip).
	//
	// *
	// AT_LATEST - the response includes only the currently open shards of the data
	// stream.
	//
	// * AT_TIMESTAMP - the response includes all shards whose start timestamp
	// is less than or equal to the given timestamp and end timestamp is greater than
	// or equal to the given timestamp or still open.
	//
	// * FROM_TIMESTAMP - the response
	// incldues all closed shards whose end timestamp is greater than or equal to the
	// given timestamp and also all open shards. Corrected to TRIM_HORIZON of the data
	// stream if FROM_TIMESTAMP is less than the TRIM_HORIZON value.
	//
	// This member is required.
	Type ShardFilterType

	// The exclusive start shardID speified in the ShardFilter parameter. This property
	// can only be used if the AFTER_SHARD_ID shard type is specified.
	ShardId *string

	// The timestamps specified in the ShardFilter parameter. A timestamp is a Unix
	// epoch date with precision in milliseconds. For example,
	// 2016-04-04T19:58:46.480-00:00 or 1459799926.480. This property can only be used
	// if FROM_TIMESTAMP or AT_TIMESTAMP shard types are specified.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// The starting position in the data stream from which to start streaming.
type StartingPosition struct {

	// You can set the starting position to one of the following values:
	// AT_SEQUENCE_NUMBER: Start streaming from the position denoted by the sequence
	// number specified in the SequenceNumber field. AFTER_SEQUENCE_NUMBER: Start
	// streaming right after the position denoted by the sequence number specified in
	// the SequenceNumber field. AT_TIMESTAMP: Start streaming from the position
	// denoted by the time stamp specified in the Timestamp field. TRIM_HORIZON: Start
	// streaming at the last untrimmed record in the shard, which is the oldest data
	// record in the shard. LATEST: Start streaming just after the most recent record
	// in the shard, so that you always read the most recent data in the shard.
	//
	// This member is required.
	Type ShardIteratorType

	// The sequence number of the data record in the shard from which to start
	// streaming. To specify a sequence number, set StartingPosition to
	// AT_SEQUENCE_NUMBER or AFTER_SEQUENCE_NUMBER.
	SequenceNumber *string

	// The time stamp of the data record from which to start reading. To specify a time
	// stamp, set StartingPosition to Type AT_TIMESTAMP. A time stamp is the Unix epoch
	// date with precision in milliseconds. For example, 2016-04-04T19:58:46.480-00:00
	// or 1459799926.480. If a record with this exact time stamp does not exist,
	// records will be streamed from the next (later) record. If the time stamp is
	// older than the current trim horizon, records will be streamed from the oldest
	// untrimmed data record (TRIM_HORIZON).
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// Represents the output for DescribeStream.
type StreamDescription struct {

	// Represents the current enhanced monitoring settings of the stream.
	//
	// This member is required.
	EnhancedMonitoring []EnhancedMetrics

	// If set to true, more shards in the stream are available to describe.
	//
	// This member is required.
	HasMoreShards *bool

	// The current retention period, in hours. Minimum value of 24. Maximum value of
	// 168.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The shards that comprise the stream.
	//
	// This member is required.
	Shards []Shard

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// This member is required.
	StreamARN *string

	// The approximate time that the stream was created.
	//
	// This member is required.
	StreamCreationTimestamp *time.Time

	// The name of the stream being described.
	//
	// This member is required.
	StreamName *string

	// The current status of the stream being described. The stream status is one of
	// the following states:
	//
	// * CREATING - The stream is being created. Kinesis Data
	// Streams immediately returns and sets StreamStatus to CREATING.
	//
	// * DELETING - The
	// stream is being deleted. The specified stream is in the DELETING state until
	// Kinesis Data Streams completes the deletion.
	//
	// * ACTIVE - The stream exists and
	// is ready for read and write operations or deletion. You should perform read and
	// write operations only on an ACTIVE stream.
	//
	// * UPDATING - Shards in the stream
	// are being merged or split. Read and write operations continue to work while the
	// stream is in the UPDATING state.
	//
	// This member is required.
	StreamStatus StreamStatus

	// The server-side encryption type used on the stream. This parameter can be one of
	// the following values:
	//
	// * NONE: Do not encrypt the records in the stream.
	//
	// * KMS:
	// Use server-side encryption on the records in the stream using a customer-managed
	// Amazon Web Services KMS key.
	EncryptionType EncryptionType

	// The GUID for the customer-managed Amazon Web Services KMS key to use for
	// encryption. This value can be a globally unique identifier, a fully specified
	// ARN to either an alias or a key, or an alias name prefixed by "alias/".You can
	// also use a master key owned by Kinesis Data Streams by specifying the alias
	// aws/kinesis.
	//
	// * Key ARN example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	// *
	// Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	// *
	// Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	// * Alias
	// name example: alias/MyAliasName
	//
	// * Master key owned by Kinesis Data Streams:
	// alias/aws/kinesis
	KeyId *string

	// Specifies the capacity mode to which you want to set your data stream.
	// Currently, in Kinesis Data Streams, you can choose between an on-demand capacity
	// mode and a provisioned capacity mode for your data streams.
	StreamModeDetails *StreamModeDetails

	noSmithyDocumentSerde
}

// Represents the output for DescribeStreamSummary
type StreamDescriptionSummary struct {

	// Represents the current enhanced monitoring settings of the stream.
	//
	// This member is required.
	EnhancedMonitoring []EnhancedMetrics

	// The number of open shards in the stream.
	//
	// This member is required.
	OpenShardCount *int32

	// The current retention period, in hours.
	//
	// This member is required.
	RetentionPeriodHours *int32

	// The Amazon Resource Name (ARN) for the stream being described.
	//
	// This member is required.
	StreamARN *string

	// The approximate time that the stream was created.
	//
	// This member is required.
	StreamCreationTimestamp *time.Time

	// The name of the stream being described.
	//
	// This member is required.
	StreamName *string

	// The current status of the stream being described. The stream status is one of
	// the following states:
	//
	// * CREATING - The stream is being created. Kinesis Data
	// Streams immediately returns and sets StreamStatus to CREATING.
	//
	// * DELETING - The
	// stream is being deleted. The specified stream is in the DELETING state until
	// Kinesis Data Streams completes the deletion.
	//
	// * ACTIVE - The stream exists and
	// is ready for read and write operations or deletion. You should perform read and
	// write operations only on an ACTIVE stream.
	//
	// * UPDATING - Shards in the stream
	// are being merged or split. Read and write operations continue to work while the
	// stream is in the UPDATING state.
	//
	// This member is required.
	StreamStatus StreamStatus

	// The number of enhanced fan-out consumers registered with the stream.
	ConsumerCount *int32

	// The encryption type used. This value is one of the following:
	//
	// * KMS
	//
	// * NONE
	EncryptionType EncryptionType

	// The GUID for the customer-managed Amazon Web Services KMS key to use for
	// encryption. This value can be a globally unique identifier, a fully specified
	// ARN to either an alias or a key, or an alias name prefixed by "alias/".You can
	// also use a master key owned by Kinesis Data Streams by specifying the alias
	// aws/kinesis.
	//
	// * Key ARN example:
	// arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	// *
	// Alias ARN example:  arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	// *
	// Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	// * Alias
	// name example: alias/MyAliasName
	//
	// * Master key owned by Kinesis Data Streams:
	// alias/aws/kinesis
	KeyId *string

	// Specifies the capacity mode to which you want to set your data stream.
	// Currently, in Kinesis Data Streams, you can choose between an on-demand
	// ycapacity mode and a provisioned capacity mode for your data streams.
	StreamModeDetails *StreamModeDetails

	noSmithyDocumentSerde
}

// Specifies the capacity mode to which you want to set your data stream.
// Currently, in Kinesis Data Streams, you can choose between an on-demand capacity
// mode and a provisioned capacity mode for your data streams.
type StreamModeDetails struct {

	// Specifies the capacity mode to which you want to set your data stream.
	// Currently, in Kinesis Data Streams, you can choose between an on-demand capacity
	// mode and a provisioned capacity mode for your data streams.
	//
	// This member is required.
	StreamMode StreamMode

	noSmithyDocumentSerde
}

// After you call SubscribeToShard, Kinesis Data Streams sends events of this type
// over an HTTP/2 connection to your consumer.
type SubscribeToShardEvent struct {

	// Use this as SequenceNumber in the next call to SubscribeToShard, with
	// StartingPosition set to AT_SEQUENCE_NUMBER or AFTER_SEQUENCE_NUMBER. Use
	// ContinuationSequenceNumber for checkpointing because it captures your shard
	// progress even when no data is written to the shard.
	//
	// This member is required.
	ContinuationSequenceNumber *string

	// The number of milliseconds the read records are from the tip of the stream,
	// indicating how far behind current time the consumer is. A value of zero
	// indicates that record processing is caught up, and there are no new records to
	// process at this moment.
	//
	// This member is required.
	MillisBehindLatest *int64

	//
	//
	// This member is required.
	Records []Record

	// The list of the child shards of the current shard, returned only at the end of
	// the current shard.
	ChildShards []ChildShard

	noSmithyDocumentSerde
}

// This is a tagged union for all of the types of events an enhanced fan-out
// consumer can receive over HTTP/2 after a call to SubscribeToShard.
//
// The following types satisfy this interface:
//  SubscribeToShardEventStreamMemberSubscribeToShardEvent
type SubscribeToShardEventStream interface {
	isSubscribeToShardEventStream()
}

// After you call SubscribeToShard, Kinesis Data Streams sends events of this type
// to your consumer. For an example of how to handle these events, see Enhanced
// Fan-Out Using the Kinesis Data Streams API.
type SubscribeToShardEventStreamMemberSubscribeToShardEvent struct {
	Value SubscribeToShardEvent

	noSmithyDocumentSerde
}

func (*SubscribeToShardEventStreamMemberSubscribeToShardEvent) isSubscribeToShardEventStream() {}

// Metadata assigned to the stream, consisting of a key-value pair.
type Tag struct {

	// A unique identifier for the tag. Maximum length: 128 characters. Valid
	// characters: Unicode letters, digits, white space, _ . / = + - % @
	//
	// This member is required.
	Key *string

	// An optional string, typically used to describe or define the tag. Maximum
	// length: 256 characters. Valid characters: Unicode letters, digits, white space,
	// _ . / = + - % @
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isSubscribeToShardEventStream() {}
