# Leroy Merlin Action s3 cache

Github action to cache artifacts to s3 bucket

## Usage

### Saving

```yml
- name: Save cache
  uses: ./
  with:
    action: put
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
    artifacts: |
      yourartifacts/*
      separated.by
      newline/*
```

### Retrieving

```yml
- name: Retrieve cache
  uses: ./
  with:
    action: get
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
```

### Cleaning

```yml
- name: Clear cache
  uses: ./
  with:
    action: delete
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
```

### Sample pipeline

```yml
- name: Checkout
  uses: actions/checkout@v2

- name: Retrieve cache
  uses: ./
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
  uses: ./
  with:
    action: put
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
    aws-region: us-east-1
    bucket: your-bucket
    key: ${{ hashFiles('yarn.lock') }}
    artifacts: |
      yourartifacts/*
      separated.by
      newline/*
```

## Running locally

### Copie o arquivo .env.example e preencha com suas informações:

```zsh
cp .env.example .env
```

Os artefatos que serão zipados deverão ficar na pasta `temp/` do projeto(a pasta é ignorada por `.gitignore`). Crie os arquivos de teste, e adicione na variável `ARTIFACTS` em seu arquivo `.env`

_Obs.: os artefatos deverão ser separados por `/n`_

Execute o script `run.sh` passando a action desejada(put|get|delete):

```zsh
./run.sh put
```
