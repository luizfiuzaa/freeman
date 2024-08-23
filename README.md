

# Freeman

O **Freeman Script** é um script de limpeza automatizado projetado para ambientes Flutter. Ele remove arquivos e diretórios desnecessários e corrige problemas de cache, preparando o ambiente para uma nova construção.

## Como Funciona

1. **Limpeza Geral**
   - `flutter clean`: Remove arquivos temporários de builds anteriores.
   - `flutter pub cache repair`: Repara o cache do pub.
   - `flutter pub cache clean`: Limpa o cache do pub.

2. **Remoção de Diretórios e Arquivos**
   - Remove diretórios e arquivos específicos do Flutter e gradle que podem causar problemas em builds subsequentes.

3. **Recuperação e Atualização de Dependências**
   - Executa `flutter pub get` para garantir que todas as dependências estejam atualizadas.

## Uso

1. Salve o script como um arquivo `.bat`, por exemplo, `freeman_cleanup.bat`.
2. Execute o script no prompt de comando do Windows para iniciar o processo de limpeza.

```batch
freeman.bat
```

## Observações

- Certifique-se de ter o Flutter instalado e configurado corretamente antes de executar o script.
- O script pode remover arquivos importantes, por isso é recomendável fazer backup do projeto antes de executá-lo.
- O script é projetado para ser executado em um ambiente Windows.

---

Você pode ajustar o conteúdo conforme necessário para refletir detalhes específicos ou adicionar informações extras.
