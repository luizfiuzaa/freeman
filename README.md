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

---

## Suporte ao FVM (Flutter Version Management)

O Freeman detecta e utiliza o [FVM](https://fvm.app/) automaticamente quando disponível. Se o FVM não estiver instalado, o Flutter global é usado como fallback sem interrupção.

### Detecção automática

Se o seu projeto possuir o diretório `.fvm/`, o Freeman já usa o FVM automaticamente — nenhuma configuração necessária.

### Via flag

Force o uso do FVM em uma execução específica:

```bash
freeman --fvm
# ou
freeman --use-fvm
```

### Via configuração global

Ative o FVM para todos os projetos de forma persistente:

```bash
freeman config --prioritize-fvm true
```

Para desativar:

```bash
freeman config --prioritize-fvm false
```

A configuração é salva em `~/.freeman/config.json`.

### Prioridade de decisão

| Condição | Resultado |
|---|---|
| Flag `--fvm` ou `--use-fvm` passada | Usa FVM |
| Config global `prioritize_fvm: true` | Usa FVM |
| Diretório `.fvm/` presente no projeto | Usa FVM |
| FVM não instalado (qualquer caso acima) | Fallback para Flutter global |
| Nenhuma das condições acima | Usa Flutter global |

---

## Observações

- Certifique-se de ter o Flutter instalado e configurado corretamente antes de executar.
- O Freeman pode remover arquivos importantes, por isso é recomendável fazer backup do projeto antes de executá-lo.
- A partir da v2.0.0 o Freeman é **cross-platform** — funciona em Windows, macOS e Linux.

---

## Colocar Freeman como comando global (variável de ambiente)

Após compilar, você pode adicionar o binário ao `PATH` do sistema para chamá-lo de qualquer diretório sem precisar informar o caminho completo.

### macOS / Linux

**Opção 1 — mover para um diretório já no PATH:**

```bash
mv freeman /usr/local/bin/freeman
```

**Opção 2 — adicionar um diretório personalizado ao PATH:**

1. Mova o binário para o diretório desejado, ex.: `~/scripts/`:
   ```bash
   mv freeman ~/scripts/freeman
   ```

2. Adicione o diretório ao `PATH` no seu arquivo de perfil (`~/.bashrc`, `~/.zshrc`, etc.):
   ```bash
   export PATH="$HOME/scripts:$PATH"
   ```

3. Recarregue o perfil:
   ```bash
   source ~/.zshrc   # ou source ~/.bashrc
   ```

Agora basta digitar `freeman` na raiz de qualquer projeto Flutter.

### Windows

1. Mova o `freeman.exe` para um diretório fixo, ex.: `C:\scripts\`.

2. Adicione esse diretório ao `PATH` do sistema:
   - Abra **Configurações do Sistema → Variáveis de Ambiente**
   - Em **Variáveis do sistema**, selecione `Path` e clique em **Editar**
   - Adicione `C:\scripts\` e confirme

3. Abra um novo terminal e chame diretamente:
   ```bat
   freeman
   ```
