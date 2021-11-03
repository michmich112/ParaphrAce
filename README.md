# ParaphrAce

## Get started
First verify that you have Docker and Docker compose installed on your device.\
Create an S3 bucket in AWS, and an IAM user in the AWS console. The IAM user must be authenticated by Access key and Secret.\
Once the bucket created create a `.env` file at the root of the project with the following variables:

| Variable                | Description                                                  |
|-------------------------|--------------------------------------------------------------|
| `AWS_ACCESS_KEY_ID`     | AWS access key for s3                                        |
| `AWS_SECRET_ACCESS_KEY` | AWS secret for accessing s3                                  |
| `AWS_REGION`            | AWS region to use                                            |
| `S3_BUCKET_NAME`        | Name of the S3 Bucket to use                                 |

Once set run the following command: `docker compose build && docker compose up`
And the services will start up and you can start using the api on port `8080`. 

You can connect to the Postgresql instace with the following credentials.
- host: `localhost`
- port: `5432`
- username: `paraphrace`
- password: `password`
- db name: `paraphrace`
- SSL: `disable`

The model is exported on port `8000` by default.

> Note: You may need to restart the application after the initial boot. Just press `Ctrl+C` and then run `docker compose up once more`. The boot process can take up to a minute.

> Troubleshooting: if you are unable to start the application due to a docker error, try allowing more memory to the docker engine (through the docker dashboard). We've tested it with 4 CPUs & 8 GB of memory.

## Client
Front end client source code written in svelte and typescript

## Server
Back end server application written in GoLang

## Model
Machine learning model and interface

