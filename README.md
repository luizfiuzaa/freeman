# Freeman

> Ferramenta de limpeza automatizada para projetos Flutter — cross-platform, sem dependências externas.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey)
![License](https://img.shields.io/github/license/luizfiuzaa/freeman)

---

## O que é o Freeman?

Ambientes Flutter acumulam caches corrompidos, arquivos de lock desatualizados e artefatos de build que causam erros difíceis de diagnosticar. O Freeman automatiza a limpeza completa desse ambiente em um único comando, economizando o tempo gasto em sequências manuais de `flutter clean`, remoção de pastas e reinstalação de dependências.

---

## Pré-requisitos

| Requisito | Versão mínima | Link |
|---|---|---|
| Go | 1.21 | [go.dev/dl](https://go.dev/dl/) |
| Flutter | qualquer | [flutter.dev](https://flutter.dev/docs/get-started/install) |
| FVM *(opcional)* | qualquer | [fvm.app](https://fvm.app/) |

---

## Instalação

### Compilando a partir do código-fonte

```bash
# Clone o repositório
git clone https://github.com/luizfiuzaa/freeman.git
cd freeman

# Compile o binário
go build -o freeman .        # macOS / Linux
go build -o freeman.exe .    # Windows
```

### Download direto

Baixe o binário pré-compilado para o seu sistema na [página de releases](https://github.com/luizfiuzaa/freeman/releases).

### Tornar global (opcional)

Adicione o binário ao `PATH` para chamá-lo de qualquer diretório:

**macOS / Linux**

```bash
# Opção 1 — mover para um diretório já no PATH
mv freeman /usr/local/bin/freeman

# Opção 2 — adicionar diretório personalizado ao PATH
mv freeman ~/scripts/freeman
echo 'export PATH="$HOME/scripts:$PATH"' >> ~/.zshrc   # ou ~/.bashrc
source ~/.zshrc
```

**Windows**

1. Mova `freeman.exe` para um diretório fixo, ex.: `C:\scripts\`.
2. Abra **Configurações do Sistema → Variáveis de Ambiente**.
3. Em **Variáveis do sistema**, edite `Path` e adicione `C:\scripts\`.
4. Abra um novo terminal e chame `freeman` diretamente.

---

## Uso rápido

Execute na raiz do seu projeto Flutter:

```bash
./freeman        # macOS / Linux
freeman.exe      # Windows
```

Para ver todos os comandos e flags disponíveis:

```bash
freeman --help
```

---

## Como funciona

O Freeman executa as seguintes etapas em sequência:

```
1. (opcional) Remove o pub cache local da máquina   --clean-cache
2. flutter clean
3. flutter pub cache repair
4. flutter pub cache clean
5. Remove diretórios e arquivos de build
6. flutter pub get
```

### Diretórios e arquivos removidos na etapa 5

| Tipo | Caminho |
|---|---|
| Diretório | `build/` |
| Diretório | `.dart_tool/` |
| Diretório | `.gradle/` |
| Diretório | `.idea/` |
| Diretório | `.packages/` |
| Diretório | `ios/Pods/` |
| Diretório | `ios/.symlinks/` |
| Diretório | `ios/Flutter/Flutter.framework/` |
| Diretório | `ios/Flutter/Flutter.podspec/` |
| Diretório | `ios/Flutter/App.framework/` |
| Diretório | `android/.gradle/` |
| Diretório | `android/.idea/` |
| Diretório | `android/build/` |
| Arquivo | `pubspec.lock` |

---

## Arquitetura

O projeto é um único binário Go (`main.go`) sem dependências externas — usa apenas a biblioteca padrão.

```
freeman/
├── main.go          # Toda a lógica da aplicação
├── go.mod           # Módulo Go
└── ~/.freeman/
    └── config.json  # Configuração global persistente (criado em tempo de execução)
```

### Fluxo interno

```
os.Args
  │
  ├── "--help" / "-h"   → printHelp()  → exit
  ├── "config"          → handleConfig() → saveConfig() → exit
  │
  └── flags de execução
        │
        ├── shouldUseFVM()     detecta FVM (flag > config > .fvm/)
        ├── cleanPubCache()    remove ~/.pub-cache ou %LOCALAPPDATA%\Pub\Cache
        ├── runFlutter()       executa flutter ou fvm flutter
        └── os.RemoveAll()     remove diretórios e arquivos de build
```

### Funções principais

| Função | Responsabilidade |
|---|---|
| `main()` | Parse de flags e orquestração do fluxo |
| `printHelp()` | Exibe a mensagem de ajuda |
| `handleConfig()` | Lê e persiste configurações globais |
| `loadConfig()` / `saveConfig()` | Serialização JSON da config em `~/.freeman/config.json` |
| `shouldUseFVM()` | Decide se usa FVM com base em flag, config e presença de `.fvm/` |
| `runFlutter()` | Executa comandos flutter ou fvm flutter |
| `pubCachePath()` | Retorna o caminho do pub cache conforme o SO |
| `cleanPubCache()` | Remove o pub cache local da máquina |
| `logRemove()` | Imprime item removido (ou simulado em dry-run) |

---

## Flags de execução

### Referência completa

| Flag | Comportamento |
|---|---|
| `--safe` | Somente `flutter clean` + `flutter pub get` |
| `--no-repair` | Pula `flutter pub cache repair` |
| `--no-cache-clean` | Pula `flutter pub cache clean` |
| `--keep-lockfile` | Preserva o `pubspec.lock` |
| `--clean-cache` | Remove o pub cache local da máquina |
| `--dry-run` | Simula a execução sem alterar nada |
| `--verbose` | Exibe cada arquivo e diretório removido |
| `--fvm` / `--use-fvm` | Força o uso do FVM nesta execução |
| `--help` / `-h` | Exibe a mensagem de ajuda |

As flags podem ser combinadas livremente:

```bash
freeman --safe --fvm
freeman --no-repair --keep-lockfile --verbose
freeman --dry-run --verbose
freeman --clean-cache --fvm
```

---

## Safe Mode (`--safe`)

Executa apenas `flutter clean` + `flutter pub get`, sem tocar em caches globais ou remover diretórios. Ideal para limpezas rápidas ou pipelines de CI/CD onde reconstruir o cache é custoso.

```bash
freeman --safe
```

---

## Dry Run (`--dry-run`)

Mostra tudo que seria removido e executado, sem realizar nenhuma alteração.

```bash
freeman --dry-run
freeman --dry-run --verbose   # com listagem de cada item individualmente
```

---

## Limpeza do Pub Cache Local (`--clean-cache`)

Remove o diretório do pub cache da máquina diretamente no sistema de arquivos, sem depender do Flutter CLI. Útil quando erros persistem mesmo após `flutter pub cache clean`.

```bash
freeman --clean-cache
```

| SO | Caminho removido |
|---|---|
| Windows | `%LOCALAPPDATA%\Pub\Cache` |
| macOS / Linux | `~/.pub-cache` |

---

## Suporte ao FVM (Flutter Version Management)

O Freeman detecta e utiliza o [FVM](https://fvm.app/) automaticamente. Se o FVM não estiver instalado, o Flutter global é usado como fallback.

### Ordem de prioridade

| Condição | Resultado |
|---|---|
| Flag `--fvm` ou `--use-fvm` passada | Usa FVM |
| Config global `prioritize_fvm: true` | Usa FVM |
| Diretório `.fvm/` presente no projeto | Usa FVM |
| FVM não instalado (qualquer caso acima) | Fallback para Flutter global |
| Nenhuma das condições acima | Usa Flutter global |

### Configuração global

```bash
# Ativar FVM para todos os projetos
freeman config --prioritize-fvm true

# Desativar
freeman config --prioritize-fvm false
```

A configuração é salva em `~/.freeman/config.json`.

---

## Exemplos de uso

```bash
# Limpeza completa padrão
freeman

# Limpeza rápida sem cache ops (boa para CI/CD)
freeman --safe

# Ver o que seria feito antes de executar
freeman --dry-run

# Ver cada arquivo sendo removido
freeman --verbose

# Pular etapas demoradas
freeman --no-repair
freeman --no-cache-clean
freeman --no-repair --no-cache-clean

# Preservar o pubspec.lock
freeman --keep-lockfile

# Limpar também o pub cache local da máquina
freeman --clean-cache

# Usar FVM explicitamente
freeman --fvm
freeman --clean-cache --fvm --verbose
```

---

## Observações

- Execute sempre na raiz do projeto Flutter.
- O Freeman remove arquivos sem confirmação — use `--dry-run` para revisar antes de executar em projetos críticos.
- A partir da v2.0.0 o Freeman é **cross-platform**: Windows, macOS e Linux.
