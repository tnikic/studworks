const API_BASE = "http://studworks.devops.csdc.fh/api";

const searchInput = document.getElementById('class-search-input');
const searchButton = document.getElementById('search-button');
const importButton = document.getElementById('import-button');
const classInfoSection = document.getElementById('class-info');
const classNameDisplay = document.getElementById('class-name');
const studentCountDisplay = document.getElementById('student-count');

let currentClassName = "";
let currentStudentCount = 0;

// --- Search logic ---
searchButton.addEventListener('click', () => {
    const className = searchInput.value.trim();
    if (className) {
        searchForClass(className);
    }
});

searchInput.addEventListener('keyup', (event) => {
    if (event.key === 'Enter') {
        searchButton.click();
    }
});

// --- Import logic with new endpoint ---
importButton.addEventListener('click', async () => {
    if (!currentClassName) return;

    try {
        // 1. Create class (ignore errors if exists)
        await axios.post(`${API_BASE}/classes/${encodeURIComponent(currentClassName)}`);

        // 2. Trigger backend to create all students for the class
        await axios.post(`${API_BASE}/students/${encodeURIComponent(currentClassName)}`);

        // 3. Refresh
        await searchForClass(currentClassName);
    } catch (error) {
        console.error("Import failed:", error);
    }
});

async function searchForClass(className) {
    hideClassInfo();
    try {
        const response = await axios.get(`${API_BASE}/students/${encodeURIComponent(className)}`);
        const students = response.data;
        currentClassName = className;
        currentStudentCount = Array.isArray(students) ? students.length : 0;
        showClassInfo(className, currentStudentCount);
    } catch (err) {
        currentClassName = className;
        currentStudentCount = 0;
        showClassInfo(className, 0);
    }
}

function showClassInfo(className, studentCount) {
    classNameDisplay.textContent = className;
    studentCountDisplay.textContent = `Students: ${studentCount}`;
    classInfoSection.style.display = "block";
}

function hideClassInfo() {
    classInfoSection.style.display = "none";
    classNameDisplay.textContent = "";
    studentCountDisplay.textContent = "";
}
