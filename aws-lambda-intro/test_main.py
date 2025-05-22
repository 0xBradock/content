from main import put
# import moto
# import os
# import boto3
# import pytest

ACCOUNT = "012345678912"
REGION = "us-east-1"
QUEUE_NAME = "queue-name"
BUCKET_NAME = "bucket-name"
S3_KEY = "path/to/key.json"
SQS_MSG = "hello from tests"


def mock_s3_put(b, k, p):
    return {
        "ResponseMetadata": {
            "RequestId": "HST2S7622FVP63VS",
            "HostId": "RtIrgv6qeCk/g9ml8Dr/d1puePRyNB19vmbWtVX+vQGawRV2AbD+RWms0XAtw88lRg186RX0WYW+xHgtsp7JIg==",
            "HTTPStatusCode": 200,
            "HTTPHeaders": {
                "x-amz-id-2": "RtIrgv6qeCk/g9ml8Dr/d1puePRyNB19vmbWtVX+vQGawRV2AbD+RWms0XAtw88lRg186RX0WYW+xHgtsp7JIg==",
                "x-amz-request-id": "HST2S7622FVP63VS",
                "date": "Sun, 18 May 2025 17:29:53 GMT",
                "x-amz-server-side-encryption": "AES256",
                "etag": "f6cae9529f28a6529502af76b0dbc882",
                "x-amz-checksum-crc32": "uZdj9Q==",
                "x-amz-checksum-type": "FULL_OBJECT",
                "content-length": "0",
                "server": "AmazonS3",
            },
            "RetryAttempts": 1,
        },
        "ETag": "f6cae9529f28a6529502af76b0dbc882",
        "ChecksumCRC32": "uZdj9Q==",
        "ChecksumType": "FULL_OBJECT",
        "ServerSideEncryption": "AES256",
    }


def test_putres():
    res = put(mock_s3_put, BUCKET_NAME, S3_KEY, "anything")
    assert res["ResponseMetadata"]["HTTPStatusCode"] == 200


# @pytest.fixture(scope="function")
# def aws_credentials():
#     """Mocked AWS Credentials for moto."""
#     os.environ["AWS_ACCESS_KEY_ID"] = "testing"
#     os.environ["AWS_SECRET_ACCESS_KEY"] = "testing"
#     os.environ["AWS_SECURITY_TOKEN"] = "testing"
#     os.environ["AWS_SESSION_TOKEN"] = "testing"
#     os.environ["AWS_DEFAULT_REGION"] = REGION


# @pytest.fixture(scope="function")
# def s3(aws_credentials):
#     with moto.mock_aws():
#         yield boto3.client("s3", region_name=REGION)


# @pytest.fixture
# def create_bucket1(s3):
#     s3.create_bucket(Bucket=BUCKET_NAME)


# def test_s3_bucket_creation_through_fixtures(create_bucket1):
#     result = boto3.client("s3").list_buckets()
#     assert len(result["Buckets"]) == 1

#     run(message(SQS_MSG), BUCKET_NAME, S3_KEY)
#     content = boto3.client("s3").get_object(
#         Bucket=BUCKET_NAME,
#         Key=S3_KEY,
#     )
#     object_content = content["Body"].read().decode("utf-8")
#     assert object_content == SQS_MSG


# def message(body: str):
#     return {
#         "Records": [
#             {
#                 "messageId": "3655ce52-0b45-4f73-b833-b2b71c5f55c1",
#                 "receiptHandle": "xr9SRbsXrsXn9zADdUtoj3kSqAB+sudM",
#                 "body": body,
#                 "attributes": {
#                     "ApproximateReceiveCount": "1",
#                     "SentTimestamp": "1747581896030",
#                     "SenderId": ACCOUNT,
#                     "ApproximateFirstReceiveTimestamp": "1747581896042",
#                 },
#                 "messageAttributes": {},
#                 "md5OfBody": "3c83a1eeba0229b3856d8a6824e1da4c",
#                 "eventSource": "aws:sqs",
#                 "eventSourceARN": f"arn:aws:sqs:{REGION}:{ACCOUNT}:{QUEUE_NAME}",
#                 "awsRegion": REGION,
#             }
#         ]
#     }
