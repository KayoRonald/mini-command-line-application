# Aplicação de Linha de Comando em Go

Este é um projeto simples que implementa uma aplicação de linha de comando em Go, permitindo a busca de endereços IP e nomes de servidores na internet.

### Funcionalidades

- **Buscar IPs**: Esta funcionalidade permite buscar endereços IP associados a um determinado nome de host na internet.
- **Buscar Servidores**: Esta funcionalidade permite buscar o nome de servidores associados a um determinado nome de host na internet.

### Instalação

1. Certifique-se de ter o Go instalado em seu sistema. Caso contrário, você pode baixá-lo e instalá-lo a partir do [site oficial do Go](https://golang.org/dl/).

2. Clone este repositório para o seu ambiente local:

   ```bash
   git clone https://github.com/KayoRonald/mini-command-line-application.git
   ```

3. Navegue até o diretório clonado:

   ```bash
   cd seu-repositorio
   ```

4. Execute o seguinte comando para compilar e instalar a aplicação:

   ```bash
   go tidy
   ```

### Uso

Depois de instalar a aplicação, você pode usá-la a partir da linha de comando. Eis como você pode usar as funcionalidades disponíveis:

- **Buscar IPs**:
  
  ```bash
  nome-da-aplicacao ip --host <nome-do-host>
  ```

  Substitua `github.com` pelo nome do host que você deseja buscar os IPs.

- **Buscar Servidores**:

  ```bash
  nome-da-aplicacao servidores --host <nome-do-host>
  ```

  Substitua `github.com` pelo nome do host que você deseja buscar os nomes dos servidores.

### Exemplo

Para buscar os IPs associados ao host "golang.org", você pode executar o seguinte comando:

```bash
nome-da-aplicacao ip --host golang.org
```

### Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
