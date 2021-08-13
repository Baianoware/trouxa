# Trouxa

Instale seus pocotes de maneria fácil e rápida, com apenas um comando independente de sua distribuição.

## O que é?

Inspirado pelo projeto de script para instalação de pacotes comuns em seu sistema, tive um *insigh*. Um pequeno programa, CLI mesmo, para fazer algo parecido, um pouco mais customizado, utilizando *scirpts* específicos para cada pacote; vou tentar exemplificar.

Imagine a seguinte situação: você acabou de instalar o seu sistema, e precisa dos programas que utiliza no dia-a-dia, você poderia simplesmente carregar um arquivo com a lista dos programas e versões que necessita e requerer a instalação de todos eles, independentemente da sua distribuição; essa é a ideia do **Trouxa**, simplificar essa instalação através de uma inferface única, independente da distribuição.

### Porque do nome Trouxa?

- Embrulho, geralmente feito com pano, para guardar ou transportar objetos; trouxo.

Se guarda pano/objetos, porque não guardar *Softwares?*

## Proposta

- Utilizar a linguagem **Go**
    - Simplicidade
    - Velocidade
    - Ferramentas
    - Compilada
- Criar *scripts* específicas para cada programa em cada gerenciador de pacotes de cada distribuição
    - Liberar a criação de *Scripts* de instalação para a comunidade, facilitando a adoção

## Histórias dos usuários

"Acabei de instalar um sistema novo, tenho a lista dos programas que vou utilizar, ao invés de instalá-los um a um, posso simplesmente criar um aquivo simples, com o nome de todos e a versão, *latest*, e pedir para o **Trouxa** fazer essas operações para mim." ~ Jhon Doe

## Caracterísitcas

- CLI
- Portável
- Loggável (produzirá *logs* durante a instação)
- Arquivo de configuração simples, para ser inscrito a qualquer momento

    ```
    python /3.0.0
    pycharm /latest
    vim /lastes
    node /16.2.2
    ```
