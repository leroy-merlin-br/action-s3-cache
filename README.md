# S3 Cache for GitHub Actions

GitHub Action that allows you to cache build artifacts to S3

Fork from [leroy-merlin-br/action-s3-cache](https://github.com/leroy-merlin-br/action-s3-cache).

Changes:
- Use tar and untar instead of zip
- Upgrade golang version

## Prerequisites
- An AWS account. [Sign up here](https://aws.amazon.com/pt/resources/create-account/).

## Usage


### Archiving artifacts

```yml
- name: Save cache
  uses: everest/action-s3-cache@v2
  with:
    action: put
    aws-region: us-east-1 # Or whatever region your bucket was created
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
    artifacts: |
      node_modules*
```

### Retrieving artifacts

```yml
- name: Retrieve cache
  uses: everest/action-s3-cache@v2
  with:
    action: get
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
```

### Clear cache

```yml
- name: Clear cache
  uses: everest/action-s3-cache@v2
  with:
    action: delete
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
  uses: everest/action-s3-cache@v2
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
  uses: everest/action-s3-cache@v2
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

## Contributing

This action uses pre-compiled Go binaries that are built and committed to version control automatically by CI.

- [ ] TODO Change to uploading built binaries to GitHub release assets instead of committing to version control (this would prevent bloating the repo size and history)

1. Create a branch and make your code changes to the Go source files
2. Open a PR targeting `develop`
3. Once merged, CI creates a release PR with built binaries for all platforms
4. Review and merge the release PR
5. This triggers tag and GitHub release creation (`v1`, `v2`, `v3`, etc.)
