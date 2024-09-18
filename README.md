# go-expert-labs-cloud-run

Programa para consultar a temperatura de previsão do tempo em celsius, fahrenheit e kelvin.

## Imagem docker

Para utilizar a imagem é necessário ter em mãos uma API KEY do https://www.weatherapi.com/.

Executando a imagem do docker hub localmente (porta 8080):

```bash
KEY="...."
docker run -e WEATHER_API_KEY="$KEY" -e PORT=8080 -p 8080:8080 betonetotbo/consulta-clima:latest
```

## Variáveis de ambiente existentes

| Variável | Descrição | Valor padrão |
|----------|-----------|--------------|
| PORT     | Porta do servidor HTTP | 8080         |
| WEATHER_API_KEY | Chave de API do https://www.weatherapi.com/ |              |

## Rodando localmente

Existe um `docker-compose.yaml` em `deployments/`. Para executá-lo basta:

```bash
cd deployments/
docker compose up -d
```

Este compose irá iniciar a aplicação localmente na porta 8080.

# Testando

Para testar, você pode:

* Utilizar o script `scripts/api.http`
* ou, via curl:
    * `curl -v http://localhost:8080/?cep=01001-000`  

## Google Cloud Run

*pedding*