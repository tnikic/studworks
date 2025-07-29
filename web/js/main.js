const apiBase = "http://studworks.devops.csdc.fh/api";

const searchForm = document.getElementById('searchForm');
const classInput = document.getElementById('classInput');
const resultDiv = document.getElementById('result');

let lastClassName = "";
let lastStudents = [];

searchForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const className = classInput.value.trim();
    if (!className) return;
    showLoading();
    await fetchStudents(className);
});

async function fetchStudents(className) {
    resultDiv.innerHTML = '';
    lastClassName = className;
    lastStudents = [];
    try {
        const res = await fetch(`${apiBase}/students/${encodeURIComponent(className)}`);
        if (!res.ok) throw new Error("No such class was found.");
        const students = await res.json();
        if (!Array.isArray(students) || students.length === 0) {
            showError("No such class was found.");
            return;
        }
        lastStudents = students;
        showClassBox(className, students);
    } catch (err) {
        showError("No such class was found.");
    }
}

function showLoading() {
    resultDiv.innerHTML = `
        <div class="class-box" style="text-align:center;">
            <span>Loading...</span>
        </div>
    `;
}

function showError(msg) {
    resultDiv.innerHTML = `<div class="alert">${msg}</div>`;
}

function showClassBox(className, students) {
    const ul = students.map(stu => 
        `<li class="student-item">${stu.first_name} ${stu.last_name} <span class="student-uid">${stu.uid}</span></li>`
    ).join('');
    resultDiv.innerHTML = `
        <div class="class-box">
            <div class="class-box-header">
                <span class="class-title">${className}</span>
                <button class="import-btn" id="importBtn">Import</button>
            </div>
            <ul class="students-list">
                ${ul}
            </ul>
        </div>
    `;
    document.getElementById('importBtn').onclick = importStudents;
}

async function importStudents() {
    const btn = document.getElementById('importBtn');
    btn.disabled = true;
    btn.textContent = "Importing...";
    try {
        // 1. Create the class (no body)
        const resClass = await fetch(`${apiBase}/classes/${encodeURIComponent(lastClassName)}`, {
            method: 'POST'
        });
        if (!resClass.ok) throw new Error("Failed to create class.");
        // 2. POST students list
        const resStudents = await fetch(`${apiBase}/students/create`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(lastStudents)
        });
        if (!resStudents.ok) throw new Error("Failed to import students.");
        // 3. Success: redirect to index.html (reload page)
        window.location = "index.html";
    } catch (err) {
        btn.disabled = false;
        btn.textContent = "Import";
        showError("Import failed. Please try again.");
    }
}
