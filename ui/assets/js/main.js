function UpdateDescription(item) {
    document.getElementById('taskTitle').innerText = item.Title;
    document.getElementById('taskDescription').innerText = item.Description;
    document.getElementById('taskID').value = item.Id;
    document.getElementById('contentDescription').style.display = "block";
    document.getElementById('notification').style.display = "none";
}
