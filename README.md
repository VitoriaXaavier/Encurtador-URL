### Desafio de Código: Encurtador de URLs

Desenvolver uma aplicação encurtadora de URLs para a WEB. Utilize a linguagem de programação, frameworks, bibliotecas e ferramentas open-source da sua preferência.

===============================================
#### Considerações Gerais
- Você deverá usar este repositório para desenvolver o desafio.
- Todos os seus commits devem estar registrados aqui, pois queremos ver como você trabalha.
- É preciso conseguir rodar seu código em um Mac OS X ou no Ubuntu.
- Deve ser possível de executar o seu desafio em uma VM ou máquina local com os seguintes comandos:

```shell
git clone {{seu-repositorio}}
cd {{seu-repositorio}}
make setup
make run
```
- Documente as instruções de uso da sua solução no [INSTRUCTIONS.md](INSTRUCTIONS.md).
- Registre todo o processo de desenvolvimento no [HISTORY.md](HISTORY.md). Registre as decisões que foram tomadas e seus porquês, o que seria feito se tivesse mais tempo (explique como você as resolveria).

=====================
#### Encurtador de URLs

**Objetivo**: Dado o acesso a uma URL encurtada, o usuário deverá ser redirecionado para a URL original.

**Requisitos**:
- Prover interface para listagem, criação e exclusão das URLs encurtadas.
- Exibir URLs encurtadas após sua criação.
- Listar URLs encurtadas com suas respectivas URLs e acessos.
- Redirecionar os acessos às URLs encurtadas para as sua respectivas URLs.
- Contabilizar os acessos às URLs encurtadas.
- Ignorar os acessos às URLs excluídas ou inexistentes.

Sugestão de interface para ser utilizada está no próprio repositório:
- [Encurtador de URL](exemplo-encurtador-input.png)
- [Exibição da URL encurtada](exemplo-encurtador-output.png)
- [Listagem das URLs encurtadas](exemplo-encurtador-list.png)

**Atenção**: A solução desenvolvida deve ser capaz de servir um volume de requisições considerável. Se possível, leve em conta questões de escalabilidade e testes com relação a este aspecto.

===============================================
#### O que será avaliado na sua solução?

Seu código será analisado por desenvolvedoras/es que avaliarão: Simplicidade e clareza da solução, estilo de código, arquitetura, testes unitários, testes funcionais, nível de automação dos testes, design da interface e documentação.

=============
#### Dicas

- Use ferramentas e bibliotecas open-source.
- Documente as decisões tomadas e seus porquês.
- Automatize o máximo possível.
- Em caso de dúvidas, pergunte.
