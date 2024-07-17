document.getElementById('addTaskForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const taskName = document.getElementById('taskName').value;
    const taskTime = document.getElementById('taskTime').value;
    const taskDone = document.getElementById('taskDone').checked;

    const utcTime = new Date(Date.parse(taskTime)).toISOString();

    const response = await fetch('/api/v1/todos/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            Name: taskName, 
            Done: taskDone,
            Time: utcTime
        })
    });

    if (response.ok) {
        location.reload();
    } else {
        console.error('Failed to add task');
    }
});

async function deleteTask(taskId) {
    const response = await fetch(`/api/v1/todos/${taskId}`, {
        method: 'DELETE'
    });

    if (response.ok) {
        document.getElementById(`task-${taskId}`).remove();
    } else {
        console.error('Failed to delete task');
    }
}

function updateTask(taskId, taskName, taskTime, taskDone) {
    const newTaskName = prompt('Update task name:', taskName);
    
    const existingDate = new Date(Date.parse(taskTime));
    const formattedDateTime = existingDate.toISOString().slice(0, 16);
    
    const newTaskTime = prompt('Update task time (YYYY-MM-DDTHH:MM):', formattedDateTime);
    const newTaskDone = confirm('Is the task done?');

    if (newTaskName && newTaskTime) {
        const utcTime = new Date(Date.parse(newTaskTime)).toISOString();

        fetch(`/api/v1/todos/${taskId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 
                Name: newTaskName, 
                Done: newTaskDone,
                Time: utcTime
            })
        }).then(response => {
            if (response.ok) {
                location.reload();
            } else {
                console.error('Failed to update task');
            }
        });
    }
}

function formatTime(timeString) {
    const date = new Date(Date.parse(timeString));
    return date.toLocaleString(); 
}