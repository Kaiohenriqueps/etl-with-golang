# etl-with-golang
An ETL project with Golang!

## Pré requisitos:
1) Docker instalado na máquina;
2) Criar um arquivo .env na pasta raíz do projeto.

## Criando o arquivo .env
Dentro do arquivo, é preciso ter as seguintes informações:
```
HOST=<hostname_postgres>
PORT=<port_postgres>
USER=<user_postgres>
PASS=<pass_postgres>
DBNAME=<database_name_postgres>
DATA_PATH=<caminho_ate_o_arquivo_de_teste>
```

## Subindo os serviços do Docker
- Após o Docker ter sido iniciado, dentro da pasta raiz do projeto rodar o seguinte comando:
```
$ docker-compose up -d
```
- Para verificar se o postgres e o pgadmin estão de pé, rodar o seguinte comando:
```
$ docker-compose logs -f postgres
$ docker-compose logs -f pgadmin
```
- Para verificar quanto tempo levou para a inserção dos dados no postgres, verifique os logs do serviço *etl* com o seguinte comando:
```
$ docker-compose logs -f etl
```

## Conferindo o resultado no PgAdmin
* Para verificar os dados na tabela do Postgres, é preciso a conexão com o PgAdmin pela seguinte url: *http://localhost:16543/browser/*.
* Utilizar o *PGADMIN_DEFAULT_EMAIL* e o *PGADMIN_DEFAULT_PASSWORD* que estão presentes no docker-compose.
* Sinta-se a vontade para mudar, caso queira.
* Para conectar o database do postgres ao PgAdmin basta seguir o passo a passo no [link](https://renatogroffe.medium.com/postgresql-docker-executando-uma-inst%C3%A2ncia-e-o-pgadmin-4-a-partir-de-containers-ad783e85b1a4).

## Higienização dos dados
Foi feito um trabalho de higienização dos dados da seguinte forma:
1) Campo de cpf: sem caracteres especiais;
2) Campo de cnpj: sem caracteres especiais;
3) Campo do tipo data: caso a data não esteja preenchida, foi colocado um valor padrão de "1111-01-01";
4) Campos do tipo numérico: caso o valor não esteja preenchido, foi colocado um valor padrão de "0";
5) Campos do tipo texto: caso o valor não esteja preenchido, foi colocado um valor padrão de "NA".

## Criação da tabela no Postgres
Foi criado uma tabela com o nome *compras*, com o campo cpf como primary key, no Postgres com o seguinte esquema:

| nome_coluna | tipo_coluna |
| ----------- | ----------- |
| cpf | TEXT |
| private | NUMERIC |
| incompleto | NUMERIC |
| dataultimacompra | DATE |
| ticketmedio | NUMERIC |
| ticketultimacompra | NUMERIC |
| lojamaisfrequente | TEXT |
| lojaultimacompra | TEXT |
| flagcpf | TEXT |
| flagcnpjfrequente | TEXT |
| flagcnpjultima | TEXT |