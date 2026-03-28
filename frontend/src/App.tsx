import { useState } from "react";

export default function App() {
  const [text, setText] = useState("");
  const [todos, setTodos] = useState<Todo[]>([]);

  type Todo = {
    readonly key: number;
    title: string;
    completed: boolean;
    removed: boolean;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setText(e.target.value);
  };

  const handleSubmit = () => {
    if (!text) return;
    const newTodo: Todo = {
      key: new Date().getTime(),
      title: text,
      completed: false,
      removed: false,
    };
    setTodos((todos) => [newTodo, ...todos]);
    setText("");
  };

  const patchTodo = (
    key: number,
    patches: Partial<Pick<Todo, "title" | "completed" | "removed">>,
  ) => {
    setTodos((todos) =>
      todos.map((todo) => {
        if (todo.key === key) {
          return { ...todo, ...patches };
        }
        return todo;
      }),
    );
  };

  return (
    <div>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          handleSubmit();
        }}
      >
        <input type="text" value={text} onChange={(e) => handleChange(e)} />
        <input type="submit" value="追加" onSubmit={handleSubmit} />
      </form>
      <ul>
        {todos.map((todo) => {
          return (
            <li key={todo.key}>
              <input
                type="checkbox"
                checked={todo.completed}
                disabled={todo.removed}
                onChange={() =>
                  patchTodo(todo.key, { completed: !todo.completed })
                }
              ></input>
              <input
                type="text"
                disabled={todo.completed || todo.removed}
                value={todo.title}
                onChange={(e) => patchTodo(todo.key, { title: e.target.value })}
              />
              <button
                onClick={() => patchTodo(todo.key, { removed: !todo.removed })}
              >
                {todo.removed ? "復元" : "削除"}
              </button>
            </li>
          );
        })}
      </ul>
    </div>
  );
}
