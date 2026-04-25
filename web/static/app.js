const API_BASE = '';

async function fetchStatus() {
    try {
        const res = await fetch(`${API_BASE}/health`);
        const data = await res.json();
        document.getElementById('status-indicator').textContent = '● ' + data.status;
        document.getElementById('status-indicator').className = 'status-healthy';
    } catch (e) {
        document.getElementById('status-indicator').textContent = '● Disconnected';
        document.getElementById('status-indicator').className = 'status-critical';
    }
}

async function analyzeAlerts() {
    const btn = document.getElementById('analyze-btn');
    btn.disabled = true;
    btn.textContent = 'Analyzing...';

    try {
        const res = await fetch(`${API_BASE}/api/analyze`, { method: 'POST' });
        const incidents = await res.json();
        renderIncidents(incidents);
    } catch (e) {
        alert('Failed to analyze alerts: ' + e.message);
    } finally {
        btn.disabled = false;
        btn.textContent = '🔍 Analyze Active Alerts';
    }
}

function renderIncidents(incidents) {
    const container = document.getElementById('incidents-list');
    
    if (!incidents || incidents.length === 0) {
        container.innerHTML = '<div class="empty-state">No active incidents detected</div>';
        return;
    }

    container.innerHTML = incidents.map(inc => `
        <div class="incident-card ${inc.severity === 'critical' ? 'critical' : ''}">
            <div class="incident-header">
                <strong>${inc.alert.name}</strong>
                <span class="severity severity-${inc.severity}">${inc.severity}</span>
            </div>
            <div class="diagnosis">${inc.diagnosis}</div>
            <ul class="remediation-list">
                ${(inc.remediation || []).map(r => `<li>${r}</li>`).join('')}
            </ul>
            <small style="color: var(--muted); margin-top: 0.5rem; display: block;">
                ${new Date(inc.createdAt).toLocaleString()}
            </small>
        </div>
    `).join('');
}

document.getElementById('analyze-btn').addEventListener('click', analyzeAlerts);
document.getElementById('refresh-btn').addEventListener('click', fetchStatus);

fetchStatus();
