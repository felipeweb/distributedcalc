# Calculadora distribuida

## Funcionalidades Implementadas

- expressões matemáticas simples
- expressões matemáticas com variáveis

## Serviços internos

- **Add**: Serviço GRPC responsável pela operação de soma
- **Sub**: Serviço GRPC responsável pela operação de subtração
- **Mul**: Serviço GRPC responsável pela operação de multiplicação
- **Quo**: Serviço GRPC responsável pela operação de divisão

## Serviços publicos

- Calc: Serviço REST responsável por receber as expressões e as variaveis fazer o parse distribuir as operações e devolver o resultado em formato JSON

## Executando o sistema

```
docker-compose up
```

## Exemplo de chamada

```
curl -i -X PUT \
   -H "Content-Type:application/json" \
   -d \
'{"expression": "X+(2Y+(X/Y))","variables":{"X":8,"Y":4}} ' \
 'http://localhost:8080'
```

## Rodando os testes

Os testes do pacote `parser` são de unidade e rodam sem o sistema estar de pé porém os resto dos testes simulam o uso da API.
Por isso para rodar todos os testes precisamos dos sistemas rodando para podemos executar o comando de testes com sucesso

```
API_ADDR=http://localhost:8080 go test ./... 
```
