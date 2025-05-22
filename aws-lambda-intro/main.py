import boto3

BUCKET_NAME = "bradock-erase"

s3 = boto3.client("s3")


def main():
    put(bucket=BUCKET_NAME, key="path/to/key.json", payload="filetoupload")


def put(fn, bucket: str, key: str, payload: str):
    return fn(bucket, key, payload)


def s3_put(bucket: str, key: str, payload: str):
    print("hello")
    response = s3.put_object(
        Body=payload,
        Bucket=bucket,
        Key=key,
    )

    return response


def get():
    response = s3.get_object(Bucket="bradock-erase", Key="profile.jpg")
    print(response)


if __name__ == "__main__":
    main()
