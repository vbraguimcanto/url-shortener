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

