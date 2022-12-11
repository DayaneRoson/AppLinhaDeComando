# AppLinhaDeComando

## Esta aplicação simples é executada por linha de comando, com duas operações:
<br/>
* Exibe o IP público de um endereço<br/>
* Exibir o nome do servidor onde o serviço passado está hospedado<br/>
<br/>

## Setup para uso

<br/>
<p>Para utilização desta aplicação é necessário gerar o arquivo executável do código, da seguinte forma:<p>
<br/>

### Download e instalação do Go
<br/>

  * Para instalar o Go na sua máquina, clique [aqui](https://go.dev/doc/install); escolha seu sistema operacional e siga os passos de instalação
  * Após os passos seguidos, abra seu terminal em qualquer diretório e execute o seguinte comando: 
    * ```go version```
  * Após a confirmação da versão, baixe o conteúdo deste repositório em um diretório.
  * Se possível, utilize uma ferramenta para compilação do código, eu aconselho o vscode.

<br/>

### Geração do arquivo executável
<br/>

  * Abra com o vscode este diretório
  * Verifique se possui a extensão do go instalado no seu vscode
  * Abra o terminal dentro do seu vscode
  * Execute o seguinte comando:
    * ```go build```

<br/>

### Execução
<br/>

  * Se preferir, coloque o diretório onde está o executável no PATH de seu sistema operacional
  * Caso contrário, também é possível executá-lo abrindo um terminal no dirétorio onde está seu executável

<br/>

  #### Comandos
<br/>

    *O comando executado pode variar de acordo com seu sistema operacional, exemplos abaixo foram executados no Windows*

  * ```./linha-de-comando ip --host (endereço web desejado)``` <br/>
    --- Retorna os IPs públicos pertencentes a este endereço<br/><br/>
  * ```./linha-de-comando servidores --host (endereço web desejado)```<br/>
    --- Retorna os nomes dos servidores onde está hospedado este endereço
  