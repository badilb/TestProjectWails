import { CreateTask, GetTasks, UpdateTask, DeleteTask } from './wailsjs/go/main/App';

document.addEventListener('DOMContentLoaded', () => {
    const taskInput = document.getElementById('task-input');
    const addTaskBtn = document.getElementById('add-task-btn');
    const taskList = document.getElementById('task-list');

    // Load tasks on page load
    loadTasks();

    // Add new task
    addTaskBtn.addEventListener('click', async () => {
        const taskName = taskInput.value.trim();
        if (taskName) {
            await CreateTask(taskName);
            taskInput.value = '';
            loadTasks();
        }
    });

    // Load tasks from the API
    async function loadTasks() {
        const tasks = await GetTasks();
        taskList.innerHTML = '';

        tasks.forEach(task => {
            const taskItem = document.createElement('li');
            taskItem.className = 'task-item';

            taskItem.innerHTML = `
                <input type="text" value="${task.name}" ${task.completed ? 'disabled' : ''}/>
                <button class="edit-btn" onclick="editTask(${task.id}, this)">✎</button>
                <button class="delete-btn" onclick="deleteTask(${task.id})">✘</button>
            `;

            taskList.appendChild(taskItem);
        });
    }

    // Edit task
    window.editTask = async (id, btn) => {
        const taskInput = btn.parentElement.querySelector('input');
        if (!taskInput.disabled) {
            const newName = taskInput.value.trim();
            await UpdateTask(id, newName);
            loadTasks();
        }
    };

    // Delete task
    window.deleteTask = async (id) => {
        await DeleteTask(id);
        loadTasks();
    };
});
