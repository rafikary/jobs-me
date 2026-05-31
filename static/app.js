// API Base URL
const API_URL = '/api/jobs';

// Open Add Modal
function openAddModal() {
    document.getElementById('modalTitle').textContent = 'Add New Application';
    document.getElementById('jobForm').reset();
    document.getElementById('jobId').value = '';
    document.getElementById('jobModal').classList.remove('hidden');
}

// Close Modal
function closeModal() {
    document.getElementById('jobModal').classList.add('hidden');
}

// Save Job (Create or Update)
async function saveJob(event) {
    event.preventDefault();
    
    const id = document.getElementById('jobId').value;
    const data = {
        company: document.getElementById('company').value,
        position: document.getElementById('position').value,
        platform: document.getElementById('platform').value,
        status: document.getElementById('status').value,
        job_link: document.getElementById('jobLink').value,
        notes: document.getElementById('notes').value
    };

    try {
        const url = id ? `${API_URL}/${id}` : API_URL;
        const method = id ? 'PUT' : 'POST';
        
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            closeModal();
            window.location.reload();
        } else {
            alert('Gagal menyimpan data');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Terjadi kesalahan');
    }
}

// Edit Job
async function editJob(id) {
    try {
        const response = await fetch(`${API_URL}/${id}`);
        const job = await response.json();

        document.getElementById('modalTitle').textContent = 'Edit Application';
        document.getElementById('jobId').value = job.id;
        document.getElementById('company').value = job.company;
        document.getElementById('position').value = job.position;
        document.getElementById('platform').value = job.platform;
        document.getElementById('status').value = job.status;
        document.getElementById('jobLink').value = job.job_link || '';
        document.getElementById('notes').value = job.notes || '';
        
        document.getElementById('jobModal').classList.remove('hidden');
    } catch (error) {
        console.error('Error:', error);
        alert('Gagal mengambil data');
    }
}

// Delete Job
async function deleteJob(id) {
    if (!confirm('Are you sure you want to delete this application?')) {
        return;
    }

    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            window.location.reload();
        } else {
            alert('Gagal menghapus data');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Terjadi kesalahan');
    }
}

// Filter Jobs
async function filterJobs(status, platform, search) {
    const params = new URLSearchParams();
    if (status) params.append('status', status);
    if (platform) params.append('platform', platform);
    if (search) params.append('search', search);

    const url = `${API_URL}?${params.toString()}`;
    
    try {
        const response = await fetch(url);
        const jobs = await response.json();
        
        // Update table (you can implement this with HTMX or reload)
        window.location.href = `/?${params.toString()}`;
    } catch (error) {
        console.error('Error:', error);
    }
}

// Close modal when clicking outside
document.getElementById('jobModal')?.addEventListener('click', function(event) {
    if (event.target === this) {
        closeModal();
    }
});

// Close modal on ESC key
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        closeModal();
    }
});
