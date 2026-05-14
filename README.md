# Freeman

O **Freeman** é uma ferramenta de limpeza automatizada para projetos Flutter. Ele remove arquivos e diretórios desnecessários e corrige problemas de cache, preparando o ambiente para uma nova build.

## Como Funciona

1. **Limpeza Geral**
   - `flutter clean`: Remove arquivos temporários de builds anteriores.
   - `flutter pub cache repair`: Repara o cache do pub.
   - `flutter pub cache clean`: Limpa o cache do pub.

2. **Remoção de Diretórios e Arquivos**
   - Remove diretórios e arquivos específicos do Flutter e Gradle que podem causar problemas em builds subsequentes.

3. **Recuperação e Atualização de Dependências**
   - Executa `flutter pub get` para garantir que todas as dependências estejam atualizadas.

## Instalação

Certifique-se de ter o [Go](https://go.dev/dl/) instalado, depois compile o binário:

```bash
go build -o freeman .
```

## Uso

Execute o binário na raiz do seu projeto Flutter:

```bash
./freeman        # macOS / Linux
freeman.exe      # Windows
```

## Observações

- Certifique-se de ter o Flutter instalado e configurado corretamente antes de executar.
- O Freeman pode remover arquivos importantes, por isso é recomendável fazer backup do projeto antes de executá-lo.
- A partir da v2.0.0 o Freeman é **cross-platform** — funciona em Windows, macOS e Linux.

---

## Colocar Freeman como alias

### macOS / Linux

Mova o binário para um diretório no seu `PATH`:

```bash
mv freeman /usr/local/bin/freeman
```

Depois é só chamar `freeman` de qualquer lugar.

### Windows

Mova o `freeman.exe` para um diretório já no `PATH` (ex.: `C:\scripts\`) e adicione esse diretório ao `PATH` do sistema, ou crie um alias no seu terminal preferido.
