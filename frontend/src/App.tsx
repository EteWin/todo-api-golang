import { useState } from "react";

export default function App() {
  const [text, setText] = useState("");
  const [todos, setTodos] = useState<Todo[]>([]);
  const [openKey, setOpenKey] = useState<number | null>(null);

  type Todo = {
    readonly key: number;
    title: string;
    completed: boolean;
    removed: boolean;
    memo?: string;
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
    patches: Partial<Pick<Todo, "title" | "completed" | "removed" | "memo">>,
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
    <div className="container">
      <form
        onSubmit={(e) => {
          e.preventDefault();
          handleSubmit();
        }}
      >
        <input
          type="text"
          value={text}
          onChange={(e) => handleChange(e)}
          placeholder="タスクを入力"
        />
        <input type="submit" value="追加" onSubmit={handleSubmit} />
      </form>
      <ul>
        {todos.map((todo) => {
          return (
            <li className="todo-item" key={todo.key}>
              <div className="todo-row">
                <label className="switch">
                  <input
                    type="checkbox"
                    checked={todo.completed}
                    disabled={todo.removed}
                    onChange={() =>
                      patchTodo(todo.key, { completed: !todo.completed })
                    }
                  />
                  <span className="slider"></span>
                </label>
                <input
                  type="text"
                  disabled={todo.completed || todo.removed}
                  value={
                    todo.removed
                      ? "(削除済み)"
                      : todo.completed
                        ? `${todo.title}（完了）`
                        : todo.title
                  }
                  maxLength={16}
                  onChange={(e) =>
                    patchTodo(todo.key, { title: e.target.value })
                  }
                  placeholder="タスクを編集"
                />
                <button
                  className={`memo-toggle ${openKey === todo.key ? "open" : ""}`}
                  disabled={todo.removed}
                  onClick={() =>
                    setOpenKey(openKey === todo.key ? null : todo.key)
                  }
                >
                  ◀
                </button>

                <button
                  onClick={() => {
                    const isRemoved = !todo.removed;
                    patchTodo(todo.key, { removed: isRemoved });
                    if (isRemoved && openKey == todo.key) {
                      setOpenKey(null);
                    }
                  }}
                >
                  {todo.removed ? "復元" : "削除"}
                </button>
              </div>
              {openKey === todo.key && (
                <div className="todo-memo-detail">
                  <textarea
                    disabled={todo.removed}
                    value={todo.memo || ""}
                    placeholder="メモを入力..."
                    onChange={(e) =>
                      patchTodo(todo.key, { memo: e.target.value })
                    }
                  />
                </div>
              )}
            </li>
          );
        })}
      </ul>
    </div>
  );
}
