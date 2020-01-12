function UpdateDescription(item) {
    document.getElementById('taskTitle').innerText = item.Title;
    document.getElementById('taskDescription').innerText = item.Description;
    document.getElementById('taskID').value = item.Id;
    document.getElementById('contentDescription').style.display = "block";
    document.getElementById('notification').style.display = "none";
}

function EditPost(task) {
    document.getElementById('editTitle').value = task.Title;
    document.getElementById('editTitle').placeholder = task.Title;
    document.getElementById('editDate').value = task.CreatedAt;
    document.getElementById('editDate').placeholder = task.CreatedAt;
    document.getElementById('editDescription').value = task.ShortDescription;
    document.getElementById('editDescription').placeholder = task.ShortDescription;
}

function openFetchPost() {
    window.location.href = '/teacher/fetchPosts'
}

function makeNewPost() {
    window.location.href = '/teacher/makeNewPost'
}

function DeletePost(task) {
    window.location.href = '/teacher/removeTask?id=' + task.Id
}

function getParentId() {
    id = document.getElementById("parentId").innerText;
    document.getElementById("formParentId").value = id
}