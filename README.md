# Aplicação - Encurtador de URLs

A aplicação tem como objetivo realizar o encurtamento de URLs. Assim, é possível por meio da aplicação realizar o redirecionamento de URLs.

## Requisitos da Aplicação

Antes de executar a aplicação, certifique-se de ter os seguintes pré-requisitos instalados em seu sistema:

- Go (versão 1.22 ou superior)
- Docker
- docker-compose

### Executando via Makefile

Para executar a aplicação com Docker via Makefile, você deverá executar o seguinte comando:

```bash
make docker-run
```

Caso a sua execução seja sem Docker, você poderá executar com o seguinte comando:


```bash
make run
```
Para executar os testes unitários, você poderá executar com o seguinte comando:

```bash
make test
```


### Testes da Aplicação

1. Para testes na aplicação, você pode fazer requisições via cURL:

```base
curl --location 'localhost:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "http://www.google.com"
}'
```

2. Para acessar a URL que será redirecionada, você pode acessar via navegador no endereço: localhost:8080/COD_URL_ENCURTADA

