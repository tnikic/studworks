// API endpoints
const API_BASE = "http://studworks.devops.csdc.fh/api";
const searchForm = document.getElementById("searchForm");
const searchInput = document.getElementById("searchInput");
const resultRoot = document.getElementById("resultRoot");
const errorText = document.getElementById("errorText");

let importInProgress = false;

searchForm.addEventListener("submit", async function(e) {
  e.preventDefault();
  errorText.style.display = "none";
  resultRoot.innerHTML = "";
  const className = searchInput.value.trim();
  if (!className) return;

  showLoading();
  try {
    const studentsRes = await fetch(`${API_BASE}/students/${encodeURIComponent(className)}`);
    if (!studentsRes.ok)
      throw new Error("Class not found or server error.");
    const students = await studentsRes.json();

    showResultCard(className, students);
  } catch (err) {
    showError(err.message || "Unknown error.");
  }
});

function showLoading() {
  resultRoot.innerHTML = `<div class="result-card"><div class="empty-text">Searching...</div></div>`;
}

function showError(msg) {
  resultRoot.innerHTML = "";
  errorText.textContent = msg;
  errorText.style.display = "";
}

function showResultCard(className, students) {
  errorText.style.display = "none";
  importInProgress = false;
  let studentsListHTML = "";
  if (students && students.length > 0) {
    studentsListHTML = `
      <ul class="students-list">
        ${students.map(s => `
          <li>
            <span class="student-name">${escapeHTML(s.firstName)} ${escapeHTML(s.lastName)}</span>
            <span class="student-uid">${escapeHTML(s.uid)}</span>
          </li>
        `).join("")}
      </ul>
    `;
  } else {
    studentsListHTML = `<div class="empty-text">No students found in this class.</div>`;
  }

  resultRoot.innerHTML = `
    <div class="result-card" id="classResultCard">
      <div class="class-row">
        <span class="class-title">${escapeHTML(className)}</span>
        <button class="import-btn" id="importBtn">Import</button>
        <span class="status-text" id="importStatus" style="display:none"></span>
      </div>
      ${studentsListHTML}
    </div>
  `;

  const importBtn = document.getElementById("importBtn");
  const importStatus = document.getElementById("importStatus");

  importBtn.addEventListener("click", async function() {
    if (importInProgress) return;
    importInProgress = true;
    importBtn.disabled = true;
    importStatus.textContent = "Importing...";
    importStatus.style.display = "";

    try {
      // 1. Create class
      let resp = await fetch(`${API_BASE}/classes/${encodeURIComponent(className)}`, {
        method: "POST",
      });
      if (!resp.ok) throw new Error("Failed to create class");

      // 2. Import each student
      for (const student of students) {
        let respS = await fetch(`${API_BASE}/students`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            class_name: className,
            full_name: student.fullName,
            uid: student.uid
          })
        });
        if (!respS.ok)
          throw new Error(`Failed to import student: ${student.fullName}`);
      }
      importStatus.textContent = "Imported!";
      importStatus.style.color = "#22c55e";
    } catch (err) {
      importStatus.textContent = err.message || "Import failed";
      importStatus.style.color = "#dc2626";
    } finally {
      importInProgress = false;
    }
  });
}

// Simple HTML escape to prevent XSS, if any
function escapeHTML(str) {
  return String(str).replace(/[&<>"']/g, function(m) {
    return ({
      "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;"
    })[m];
  });
}
