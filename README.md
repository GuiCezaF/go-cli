# Scaff CLI
Bem-vindo ao Scaff CLI! Esta ferramenta de linha de comando (CLI), desenvolvida em Go, permite a criação rápida e eficiente de scaffolds de projetos. Com ela, você pode gerar estruturas de projeto padronizadas, economizando tempo e garantindo consistência no desenvolvimento.

## Índice

- [Instalação](#instalação)
- [Uso](#uso)
  - [Comandos Disponíveis](#comandos-disponíveis)
  - [Opções](#opções)
- [Tecnlogias suportadas](#tecnlogias-suportadas)



## Instalação
Primeiro, certifique-se de ter o Go instalado em sua máquina. Para instalar a CLI, clone o repositório e compile o código:
```bash
git clone https://github.com/GuiCezaF/scaff-cli.git
cd scaff-cli
go build -o scaff 
```
Você pode mover o binário scaffold gerado para um diretório incluído em seu PATH, para facilitar o uso:
```bash
mv scaff /usr/local/bin/
```

## Uso
Após a instalação, você pode criar scaffolds de projetos utilizando os comandos da CLI.

### Comandos Disponíveis
- scaffold create --name <nome-do-projeto> --path <caminho>: Cria um novo projeto baseado em um scaffold padrão.
### Opções
- '--name, -n <nome-do-projeto>:' Especifica o nome do projeto que será criado.
- '--path, -p <caminho>:' Define o caminho onde o projeto será criado. Se não for especificado, o projeto será criado no diretório atual.

## Exemplos
Criando um novo projeto com o nome meu-projeto no diretório /home/usuario/projetos:
```bash
scaff create --name meu-projeto --path /home/usuario/projetos
```
Criando um novo projeto no diretório atual:
```bash
scaffold create --name meu-projeto --path .
```

## Tecnlogias suportadas 
  - [x] Golang
  - [ ] NodeJs
  - [ ] Python

