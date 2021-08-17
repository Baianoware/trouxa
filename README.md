# Trouxa

Instale os seus pacotes de maneira fácil e rápida, com apenas um comando.

## O que é?

Inspirado pelo projeto do Gustavo do script para instalação de pacotes comuns no seu sistema, tive um *insigh*. Um
pequeno programa, CLI mesmo, para fazer algo parecido, um pouco mais customizado, utilizando; vou tentar exemplificar.

Imagine a seguinte situação: você acabou de instalar o seu sistema, e precisa dos programas que utiliza no dia a dia,
você poderia simplesmente carregar um arquivo com a lista dos programas e versões que necessita e requerer a instalação
de todos eles, independentemente da sua distribuição; essa é a ideia do **Trouxa**, simplificar essa instalação através
de uma interface única, independente da distribuição.

### Porque do nome Trouxa?

No dicionário, trouxa significa: Embrulho, geralmente feito com pano, para guardar ou transportar objetos; trouxo.

# Proposta

- Utilizar a linguagem **Go**
    - Simplicidade
    - Velocidade
    - Ferramentas
    - Compilada

## Como utilizar?

Você pode rodar o projeto direto com:

```sh
go run ./cmd/ 
```

Ou construir o binário para a pasta `build` e utilizá-lo com mais liberdade, como foi pensando:

```sh
go build -o build/trouxa ./cmd/
```

Ele funcionará conforme o gerenciador de pacotes da distribuição.

Por exemplo, num sistema que faz o uso do *pacman* a utilização ficaria assim:

```sh
trouxa pacman -p packages.txt
```

Ou sem especificar o nome do arquivo, se o mesmo já existir com o nome `packages.txt` no mesmo diretório:
```sh
trouxa pacman -p
```

## Características

- CLI
- Portável
- Loggável (produzirá *logs* durante a instalação)
- Arquivo de configuração simples, para ser inscrito a qualquer momento.

    ```
    python
    pycharm
    neovim
    node
    ```

## Informações extras úteis

- https://github.com/golang-standards/project-layout
