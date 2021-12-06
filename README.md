# S3 Cache for GitHub Actions
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/leroy-merlin-br/action-s3-cache/Build%20and%20publish?style=flat-square) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/leroy-merlin-br/action-s3-cache?style=flat-square) ![Codacy grade](https://img.shields.io/codacy/grade/71fc49e81b654ddfa1379a2c50f6ea8a?style=flat-square)

GitHub Action that allows you to cache build artifacts to S3

## Prerequisites
- An AWS account. [Sign up here](https://aws.amazon.com/pt/resources/create-account/).
- AWS Access and Secret Keys. More info [here](https://aws.amazon.com/pt/premiumsupport/knowledge-center/create-access-key/).
- An empty S3 bucket.

## Usage

Set up the following AWS credentials as secrets in your repository, `AWS_ACCESS_KEY_ID` and `AWS_ACCESS_KEY_ID`

S3 Cache for GitHub Actions supports builds on Linux, Windows and MacOS.

### Archiving artifacts

```yml
- name: Save cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: put
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1 # Or whatever region your bucket was created
    bucket: your-bucket
    s3-class: ONEZONE_IA # It's STANDARD by default. It can be either STANDARD, 
    # REDUCED_REDUDANCY, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER, DEEP_ARCHIVE or STANDARD_IA.
    key: ${{ hashFiles('yarn.lock') }}
    artifacts: |
      node_modules*
```

### Retrieving artifacts

```yml
- name: Retrieve cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: get
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
```

### Clear cache

```yml
- name: Clear cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: delete
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
```

## Example

The following example shows a simple pipeline using S3 Cache GitHub Action:


```yml
- name: Checkout
  uses: actions/checkout@v2

- name: Retrieve cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: get
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}

- name: Install dependencies
  run: yarn

- name: Save cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: put
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    s3-class: STANDARD_IA
    key: ${{ hashFiles('yarn.lock') }}
    artifacts: |
      node_modules/*
```

Example For S3 Server with Custom Endpoint(Minio, etc)
```yml
- name: Checkout
  uses: actions/checkout@v2

- name: Retrieve cache
  uses: leroy-merlin-br/action-s3-cache@v1
  with:
    action: get
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
    endpoint: https://my-s3-server.example.com # don't forget to add your s3 endpoint here
```