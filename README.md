# primitive-obsession

# descrição
O que o app faz?
    O app serve como apoio para validar que não é uma boa ideia trabalhar com tipos primitivos em propriedades que requer validações.
Com o que é construído?
    Para esse estudo, escolhemos a linguagem Golang
Qual foi a motivação?
    A ideia aqui é abordar o conceito explorado pelo Martin Flower que visa trabalhar tipos alto nível ao invés de tipos comuns para representar atributos.

## Object Calisthenics (9 regras)
1. Apenas um nível de identação
2. Nunca use else
3. Encapsule todos os tipos primitivos e strings (obsessão por tipos primitivos)
- Baixa reutilização de código
- Impossível escrever testes unitários
- Provoca duplicação de código
- Alto acoplamento
- Tipos primitivos não tem comportamento
- Impossibilidade de utilizar a modelagem tática do DDD
- Value Objects
4. Envolva coleções em classes
5. Use apenas um ponto por linha (Lei de Demeter)
6. Não abrevie nomes de variáveis e métodos
7. Mantenha todas as classes pequenas (máx. 50 linhas)
8. Nenhuma classe com mais de duas variáveis de instância
9. Não use Getters ou Setters

para iniciar um projeto vazio

$ go mod init (nome-da-pasta-do-projeto)

para executar o projeto sem compilar

$ go run main.go

go build
