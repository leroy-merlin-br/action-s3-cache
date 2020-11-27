# action-s3-cache

Github action to cache dependencies and build outputs to s3 bucket

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
