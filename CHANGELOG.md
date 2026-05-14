# Changelog

Todas as mudanças notáveis deste projeto serão documentadas aqui.

O formato segue [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/).

---

## [2.3.1]

### Adicionado

- **Flag de ajuda (`--help` / `-h`)** — resolve [#8](https://github.com/luizfiuzaa/freeman/issues/8)
  - Exibe todos os comandos, flags, modos de execução, opções de configuração e exemplos de uso
  - Pode ser chamada a qualquer momento: `freeman --help` ou `freeman -h`

### Documentação

- **README.md reformulado** — resolve [#5](https://github.com/luizfiuzaa/freeman/issues/5)
  - Seção de pré-requisitos com versões mínimas de Go e Flutter
  - Guia de instalação completo: compilação a partir do fonte, download de binário e configuração global no PATH (macOS, Linux e Windows)
  - Seção de arquitetura: estrutura de arquivos, fluxo interno e tabela de funções principais
  - Tabela de diretórios e arquivos removidos pelo Freeman
  - Referência completa de todas as flags com exemplos combinados
  - Seções dedicadas para Safe Mode, Dry Run, `--clean-cache` e FVM

---

## [2.3.0]

### Adicionado

- **Safe Mode (`--safe`)** — resolve [#7](https://github.com/luizfiuzaa/freeman/issues/7)
  - Executa apenas `flutter clean` e `flutter pub get`, sem remover caches globais ou diretórios do projeto
  - Ideal para limpezas rápidas ou ambientes CI/CD onde a perda de cache é custosa

- **Flags de controle seletivo de limpeza**
  - `--no-repair`: pula o `flutter pub cache repair`
  - `--no-cache-clean`: pula o `flutter pub cache clean`
  - `--keep-lockfile`: preserva o `pubspec.lock` durante a limpeza

- **Dry Run (`--dry-run`)**: exibe tudo que seria removido e executado sem realizar nenhuma alteração

- **Verbose Mode (`--verbose`)**: exibe cada diretório e arquivo removido durante a limpeza específica

- **Contador de steps dinâmico**: o progresso `N/N` agora reflete o número real de comandos flutter que serão executados conforme as flags ativas

---

## [2.2.0]

### Adicionado

- **Flag `--clean-cache`** para limpeza do pub cache local da máquina — resolve [#6](https://github.com/luizfiuzaa/freeman/issues/6)
  - Remove o diretório de cache do pub diretamente no sistema operacional, sem depender do Flutter CLI
  - Caminhos por SO:
    - **Windows:** `%LOCALAPPDATA%\Pub\Cache`
    - **macOS / Linux:** `~/.pub-cache`
  - Pode ser combinada com as demais flags: `freeman --clean-cache --fvm`

---

## [2.1.0]

### Adicionado

- **Suporte ao FVM (Flutter Version Management)** — resolve [#4](https://github.com/luizfiuzaa/freeman/issues/4)
  - Detecção automática do diretório `.fvm/` no projeto: o Freeman usa o FVM sem nenhuma configuração adicional
  - Flag `--fvm` / `--use-fvm` para forçar o uso do FVM em uma execução específica
  - Subcomando `config` com opção `--prioritize-fvm true/false` para configuração global persistente
  - Fallback automático para o Flutter global caso o FVM não esteja instalado
  - Configuração salva em `~/.freeman/config.json`
- Documentação do FVM adicionada ao `README.md` com tabela de prioridade de decisão

---

## [2.0.0]

### Adicionado

- Reescrita completa em **Go** — substituindo os scripts `.bat` anteriores
- Suporte **cross-platform**: Windows, macOS e Linux
- Remoção automática de diretórios e arquivos problemáticos (`build`, `.dart_tool`, `.gradle`, `ios/Pods`, etc.)
- Execução sequencial de `flutter clean`, `flutter pub cache repair`, `flutter pub cache clean` e `flutter pub get`

### Removido

- Scripts `.bat` (Windows-only) da versão anterior

---

## [1.0.0]

### Adicionado

- Versão inicial em script `.bat` para Windows
- Limpeza de cache Flutter e remoção de diretórios de build
