## API Todo List em memória usando o pacote net/http nativo do golang


| Método | Endpoint       | Descrição                     | Request Body                              | Response                       |
|--------|----------------|-------------------------------|-------------------------------------------|--------------------------------|
| GET    | `/todos`       | Lista todas as tarefas        | –                                         | Array de tarefas JSON          |
| GET    | `/todos/{id}`  | Retorna tarefa específica     | –                                         | Objeto JSON da tarefa          |
| POST   | `/todos`       | Cria uma nova tarefa          | `{ "title": "string", "done": bool }`   | Objeto JSON da tarefa criada   |
| PUT    | `/todos/{id}`  | Atualiza uma tarefa existente | `{ "title": "string", "done": bool }` | Objeto JSON da tarefa atualizada |
| DELETE | `/todos/{id}`  | Remove uma tarefa             | –                                         | Array de tarefas JSON atualizado         |