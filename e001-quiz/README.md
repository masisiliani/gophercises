# Exercício 001 - QUIZ


## Passos de resolução
### Parte 1
- [x] Criar um programa que consome uma flag, que deve ser o nome do arquivo csv
- [x] Abrir um arquivo CSV
- [x] Ler um CSV
- [x] Manter os desafios do CSV na memória
- [x] Retornar um dos desafios para o usuário
- [x] Marcar quando o placar do usuário. Somar um para cada acerto.
- [x] Retornar um novo desafio para o usuário logo em seguida, independente se ele acertou ou errou
- [x] O arquivo deve estar formatado da seguinte forma: `pergunta, resposta`, onde a primeira coluna antes da vírgula é a operação, e a segunda coluna é a resposta.
- [x] Mostrar totalizadores de respostas, na seguinte formatação: `respondidas: x | corretas: y | erradas: z | total: xx`
- [x] No final o programa deve retornar a quantidade de questões corretas, e quantas foram perguntadas.
- [ ] Cobrir com testes
- [ ] Os testes devem cobrir o seguinte caso `"what 2+2, sir?",4` como uma operação válida.


### Parte 2
- [ ] Adicionar um timer, de inicialmente 30 segundos.
- [ ] Alterar o timer para ser customizado via flag
- [ ] O quizz deve parar se o exceder o tempo limite do timer.
- [ ] Solicitar que o usuário aperte a tecla de barra invertida `\` para inicializar o timer  


Exercícios Bonus 
- [ ] Selecionar um item da lista de desafios randomicamente 
- [x] Permitir que o usuário escolha quando parar 
- [x] Mostrar totalizadores após usuários escolher quando parar
