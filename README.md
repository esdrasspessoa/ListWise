# ListWise

ListWise é uma aplicação simples que permite aos usuários criar listas de compras e obter uma estimativa de custo para suas compras. Com ListWise, você pode planejar suas compras com antecedência, adicionar itens à sua lista e obter uma ideia aproximada de quanto gastará ao final. Este projeto é escrito em Go e usa um banco de dados MySQL para armazenar informações sobre preços de itens.

## Funcionalidades

- **Criação de Listas**: Você pode criar listas de compras com facilidade, adicionando os itens que deseja comprar.

- **Estimativa de Custos**: ListWise se conecta a um banco de dados de preços de produtos e fornece uma estimativa aproximada de quanto você gastará com base nos preços médios dos produtos em sua lista.

## Como Usar

1. **Clonar o Repositório**: Clone este repositório em seu sistema local.

2. **Configurar o Banco de Dados**: Certifique-se de ter um banco de dados configurado com preços de produtos. Você pode personalizar o banco de dados para suas necessidades.

3. **Instalar Dependências**: Execute `go get` para instalar todas as dependências do projeto.

4. **Executar o Aplicativo**: Execute o aplicativo com `go run main.go`. O aplicativo irá solicitar que você insira os detalhes da compra e os itens da lista.

5. **Obter a Estimativa de Custos**: Após adicionar todos os itens à sua lista, o aplicativo fornecerá uma estimativa aproximada de custos com base nos preços médios dos produtos.

6. **Gerenciar suas Listas**: Você pode criar e gerenciar várias listas de compras, fornecendo informações sobre diferentes compras.

## Contribuição

Se você deseja contribuir para o projeto ListWise, sinta-se à vontade para criar um fork do repositório, fazer melhorias e enviar um pull request. Estamos abertos a sugestões e colaborações para tornar o aplicativo ainda melhor.

## Agradecimentos

Agradecemos por usar o ListWise. Esperamos que este aplicativo ajude você a planejar suas compras com mais eficiência e a economizar tempo e dinheiro.

## Licença

Este projeto é distribuído sob a [Licença MIT](LICENSE).
