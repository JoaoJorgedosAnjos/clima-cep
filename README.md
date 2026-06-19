# Desafio: Clima por CEP no Google Cloud Run

Este projeto é a resolução do desafio técnico de Deploy no Cloud Run da pós-graduação Go Expert. Ele recebe um CEP, identifica a cidade via ViaCEP e retorna as temperaturas atuais (Celsius, Fahrenheit e Kelvin) consumindo a WeatherAPI.

## 🚀 URL da Aplicação (Cloud Run)
**Acesse a API em produção:**
`https://clima-cep-702529409828.southamerica-east1.run.app/clima/73150150`
---

## 🛠️ Como rodar localmente com Docker

### Clone o repositório:

    git clone [https://github.com/joaojorgedosanjos/clima-cep.git](https://github.com/joaojorgedosanjos/clima-cep.git)
    cd clima-cep

### Faça o build da imagem Docker:


    docker build -t clima-cep .


### Execute o container (substitua pela sua API Key da WeatherAPI):


    docker run -p 8080:8080 -e WEATHER_API_KEY="sua_chave_aqui" clima-cep

### Faça uma requisição para a API:

      curl http://localhost:8080/clima/73150150

## Como rodar os testes automatizados

Para rodar os testes unitários das lógicas de conversão de temperatura, execute na raiz do projeto:

    go test ./internal/temperature/...


