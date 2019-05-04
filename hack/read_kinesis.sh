SHARD_ITERATOR=`aws kinesis get-shard-iterator --shard-id shardId-000000000000 --shard-iterator-type TRIM_HORIZON --stream-name $AWS_EVENT_STREAM_NAME --query 'ShardIterator'`

while [ X$SHARD_ITERATOR != X ]
do
	GET_RECORDS=`aws kinesis get-records --shard-iterator $SHARD_ITERATOR`
	SHARD_ITERATOR=`echo $GET_RECORDS | jq '.NextShardIterator'`
	echo $GET_RECORDS | jq '.Records'
	echo $GET_RECORDS | jq '.MillisBehindLatest'
	echo $SHARD_ITERATOR
	sleep 5
done
