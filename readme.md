# Arquitetura Hexagonal - Ports and Adapters

## Introdução

- Mudar completamente a forma de desenvolver software
- Desenvolver software de qualidade

### Negócios

Complexidade de negócio, que é o problema que eu estou sendo pago para resolver

### Complexidade técnica

Complexidade técnica, que é a complexidade que eu tenho para resolver a regra de negócio. O que eu vou usar de aparato técnico para resolver o problema de negócio

### Separação das complexidades

Não se deve misturar a complexidade técnica com a de negócios, para que a complexidade técnica não invada as regras de negócio.

## O Ciclo de um projeto

### Pontos importantes sobre arquitetura

- Crescimento sustentável
- Software precisa se pagar ao passar do tempo
- Software deve ser desenhado por você e não pelo seu framework
- O software deve ser igual um lego, com peças que precisam se encaixar e eventualmente serem substituídos

**Arquitetura diz respeito ao futuro do software. CRUD qualquer um faz.**

### Ciclo de vida de muitos projetos

- Fase 1
  - Banco de dados
  - Cadastros
  - Validações
  - Servidor Web
  - Controllers
  - Views
  - Autenticação
  - Upload de arquivos
- Fase 2
  - Regras de negócio
  - Criação de APIs
  - Consumo de APIs
  - Autorização
  - Relatórios
  - Logs
- Fase 3
  - Mais acessos
  - Upgrade de hardware
  - Cache
  - API parceiros
  - Regras parceiros
  - Relatórios
- Fase 4
  - Mais acessos
  - Upgrade de hardware
  - BD relatórios
  - Comandos
  - V2 da API
- Fase 5
  - Escala horizontal
  - Sessões
  - Uploads
  - Refatoração
  - Autoscaling
  - CI/CD
- Fase 6
  - GraphQL
  - Bugs constantes
  - Logs? Ops
  - Integração CRM
  - Migração para React
- Fase 7
  - Inconsistência CRM
  - Containers
  - CI/CD
  - Memória
  - Logs
  - Se livrar do legado
- Fase 8
  - Microsserviços
  - DB compartilhado
  - Problemas de tracing
  - Lentidão
  - Custo elevado
- Fase 9
  - Kubernetes
  - CI/CD
  - Mensageria
  - Perda de mensagens
  - Consultorias
- Fase 10
  - Use a imaginação

## Reflexões importantes

- Visão de futuro
- Limites bem definidos
- Troca e adição de componentes
- Escala
- Otimizações frequentes
- Preparado para mudanças bruscas

### Reflexões

- Está sendo doloroso para o developer?
- Poderia ter sido evitado?
- O software está se pagando?
- Será que a relação com o cliente está boa?
- O cliente terá prejuízo com a brusca mudança arquitetural?
- Em qual momento tudo se perdeu?
- Se você fosse novo na equipe, você julgaria os desenvolvedores que fizera tudo isso?

## Arquitetura vs Design de Software

"Atividades relacionadas a arquitetura de software são sempre de design. Entretanto, nem todas atividades de design são sobre arquitetura. O objetivo primário da arquitetura de software é garantir que os atributos de negócio sejam atendidos pelo sistema. Qualquer decisão de design que não tenha relação com este objetivo não é arquitetural. Todas as decisões de design para um componente que não sejam visíveis fora dele, geralmente, também não são.^Elemar Jr.

## Apresentando arquitetura Hexagonal

"Allow an application to e equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases." Cockburn

### Arquitetura hexagonal ou ports and adapters

- Aplicação no centro
- Vários adaptadores com conexões externas
  - Nenhuma aplicação externa acessa a aplicação diretamente, mas sim os adaptadores.

### Dinâmica da arquitetura

- Definição de limites e proteção nas regras da aplicação
- Componentização e desacoplamento
  - Logs
  - Cache
  - Upload
  - Banco de dados
  - Comandos
  - Filas
  - HTTP / APIs / GraphQL
- Facilidade na quebra para microsserviços

### Lógica Básica

- Aplicação no centro
- Cliente no lado esquerdo (qualquer coisa que vai acessar minha aplicação)
  - REST
  - CLI
  - RPC
  - GraphQL
  - UI
- Servidor no lado direito (tudo que minha aplicação precisa acessar)
  - DB
  - REDIS
  - Filesystem
  - Lambda

## Hexagonal vs Clean vs Onion

- Não há padrão estabelecido de como o código deve ser organizado
- Hexagonal e Onion seguem o mesmo princípio de separação, porém definem as camadas que devo utilizar
- Quanto mais desacoplado for o código, melhor.
