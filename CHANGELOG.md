# Changelog

Todas as mudanças notáveis deste projeto serão documentadas aqui.

O formato segue [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/).

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
