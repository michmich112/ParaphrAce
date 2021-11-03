# Server
Back end server for ParaphrAce Written in Go
Exposes the api on port 8080 by default.

## Requirements
You will need the following to run the server:
- PostgreSQL Database: A postgresql database (v9 or higher);
- AWS S3 Bucket: An S3 bucket with credentials for an IAM account with the permissions to put objects;
- Paraphrasing Api: A Paraphrasing API, specifically the one in this monorepo that implements the Pegasus model.

> Note: The PostgreSQL database must be using version 9 or later

## Routes
| Route                   | Description                                              |
|-------------------------|----------------------------------------------------------|
| `api/user/create`       | Creates a new user and retrieve the valid session token. |
| `api/paraphrase/create` | Create a new paraphrase from an input sentence.          |

## Environment variables
| Variable                | Required | Description                                                  |
|-------------------------|----------|--------------------------------------------------------------|
| `POSTGRES_HOST`         | Yes      | Host address for the PostgreSQL database                     |
| `POSTGRES_PORT`         | Yes      | Port for the PostgreSQL database                             |
| `POSTGRES_DB_NAME`      | Yes      | Name of the PostgreSQL database to connect to                |
| `POSTGRES_USER`         | Yes      | User with which to connect to the PostgreSQL database        |
| `POSTGRES_PASSWORD`     | Yes      | Password of the User with which to connect to the PostgreSQL |
| `POSTGRES_SSL`          | No       | Ssl connection mode, enabled by default                      |         
| `AWS_ACCESS_KEY_ID`     | Yes      | AWS access key for s3                                        |
| `AWS_SECRET_ACCESS_KEY` | Yes      | AWS secret for accessing s3                                  |
| `AWS_REGION`            | Yes      | AWS region to use                                            |
| `S3_BUCKET_NAME`        | Yes      | Name of the S3 Bucket to use                                 |
| `PEGASUS_API_URL`       | Yes      | Url of the Pegasus api. Api must expose `/paraphrase` route. |
