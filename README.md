## API Todo List em memória usando o pacote net/http nativo do golang


| Método | Endpoint       | Descrição                     | Request Body                              | Response                       |
|--------|----------------|-------------------------------|-------------------------------------------|--------------------------------|
| GET    | `/todos`       | Lista todas as tarefas        | –                                         | Array de tarefas JSON          |
| GET    | `/todo/{id}`  | Retorna tarefa específica     | –                                         | Objeto JSON da tarefa          |
| POST   | `/todos/create`       | Cria uma nova tarefa          | `{ "name": "string", "done": bool }`   | Objeto JSON da tarefa criada   |
| PUT    | `/todos/update`  | Atualiza uma tarefa existente | `{ "name": "string", "done": bool }` | Objeto JSON da tarefa atualizada |
| DELETE | `/todos/delete`  | Remove uma tarefa             | –                                         | Array de tarefas JSON atualizado         |