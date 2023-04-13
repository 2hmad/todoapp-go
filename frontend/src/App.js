import "./App.css";
import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [tasks, setTasks] = useState([]);

  const [newTaskTitle, setNewTaskTitle] = useState("");

  const fetchData = async () => {
    const response = await axios.get("http://localhost:8085/api/v1/todos");
    if (response.data) {
      setTasks((prevTasks) => [...prevTasks, response.data]);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleAddTask = async () => {
    const newTask = {
      id: tasks.length + 1,
      title: newTaskTitle,
      done: false,
    };
    setTasks([...tasks, newTask]);
    setNewTaskTitle("");

    await axios.post(`http://localhost:8085/api/v1/todos`, {
      title: newTaskTitle,
      done: false,
    });
  };

  const handleMarkTaskAsDone = async (taskId) => {
    const updatedTasks = tasks.map((task) => {
      if (task.id === taskId) {
        task.done = true;
      }
      return task;
    });
    setTasks(updatedTasks);

    await axios.put(`http://localhost:8085/api/v1/todos/${taskId}`, {
      done: true,
    });
  };

  const handleMarkTaskAsUndone = async (taskId) => {
    const updatedTasks = tasks.map((task) => {
      if (task.id === taskId) {
        task.done = false;
      }
      return task;
    });
    setTasks(updatedTasks);
    await axios.put(`http://localhost:8085/api/v1/todos/${taskId}`, {
      done: false,
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            gap: "20px",
          }}
        >
          <input
            type="text"
            name="title"
            placeholder="Task title"
            style={{
              padding: "10px",
              width: "300px",
              borderRadius: "5px",
              border: "1px solid #ccc",
              outline: "none",
            }}
            value={newTaskTitle}
            onChange={(e) => setNewTaskTitle(e.target.value)}
          />
          <button
            style={{
              padding: "10px",
              width: "300px",
              borderRadius: "5px",
              border: "1px solid #ccc",
              outline: "none",
              cursor: "pointer",
            }}
            onClick={() => handleAddTask()}
          >
            Add Task
          </button>
          <ul
            style={{
              display: "flex",
              flexDirection: "column",
              gap: "20px",
              width: "100%",
            }}
          >
            {tasks.map((task) => (
              <li
                key={task.id}
                style={{
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "space-between",
                  width: "300px",
                }}
              >
                <span
                  style={{
                    textDecoration: task.done ? "line-through" : "none",
                  }}
                >
                  {task.title}
                </span>
                <button
                  onClick={() => {
                    task.done
                      ? handleMarkTaskAsUndone(task.id)
                      : handleMarkTaskAsDone(task.id);
                  }}
                >
                  {task.done ? "Mark as undone" : "Mark as done"}
                </button>
              </li>
            ))}
          </ul>
        </div>
      </header>
    </div>
  );
}

export default App;
