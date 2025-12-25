"""
Agent Worker - The AI Brain

This worker connects to LiveKit rooms and runs the VAD → STT → LLM → TTS pipeline.
It handles:
- Voice Activity Detection (Silero)
- Speech-to-Text (Deepgram)
- LLM Response (OpenAI/Groq)
- Text-to-Speech (ElevenLabs/Cartesia)
- Barge-in (interruption handling)
"""

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import os
import asyncio
from typing import Optional

app = FastAPI(
    title="Agent Worker",
    description="Voice AI Agent Worker - VAD → STT → LLM → TTS pipeline",
    version="1.0.0"
)

# CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Worker state
class WorkerState:
    def __init__(self):
        self.accepting_jobs: bool = True
        self.active_jobs: int = 0
        self.total_jobs_processed: int = 0

state = WorkerState()


# Health check
@app.get("/health")
async def health_check():
    return {
        "status": "healthy",
        "service": "agent-worker",
        "accepting_jobs": state.accepting_jobs,
        "active_jobs": state.active_jobs
    }


# Worker status
@app.get("/status")
async def get_status():
    """Get worker status and statistics"""
    return {
        "accepting_jobs": state.accepting_jobs,
        "active_jobs": state.active_jobs,
        "total_jobs_processed": state.total_jobs_processed
    }


# Graceful shutdown (drain)
@app.post("/drain")
async def drain():
    """
    Stop accepting new jobs and wait for active jobs to complete.
    Used for graceful pod termination in Kubernetes.
    """
    state.accepting_jobs = False
    return {
        "message": "Draining started",
        "active_jobs": state.active_jobs,
        "accepting_jobs": state.accepting_jobs
    }


# Prometheus metrics
@app.get("/metrics")
async def metrics():
    """Prometheus metrics endpoint"""
    # TODO: Expose Prometheus metrics
    metrics_text = f"""
# HELP agent_worker_active_jobs Number of active jobs
# TYPE agent_worker_active_jobs gauge
agent_worker_active_jobs {state.active_jobs}

# HELP agent_worker_total_jobs Total jobs processed
# TYPE agent_worker_total_jobs counter
agent_worker_total_jobs {state.total_jobs_processed}

# HELP agent_worker_accepting_jobs Whether worker is accepting jobs
# TYPE agent_worker_accepting_jobs gauge
agent_worker_accepting_jobs {1 if state.accepting_jobs else 0}
"""
    return metrics_text


# ===================
# LIVEKIT AGENT (Stub)
# ===================

async def entrypoint():
    """
    Main entrypoint for the LiveKit agent.
    This would normally connect to LiveKit and wait for job assignments.
    
    Pseudo-code:
    1. Connect to LiveKit server
    2. Wait for job assignment
    3. Join room
    4. Setup VAD on incoming audio track
    5. On speech detected: STT → LLM → TTS
    6. On barge-in: cancel current TTS, clear LLM queue
    7. Loop until call ends
    """
    pass


# ===================
# BARGE-IN HANDLER (Stub)
# ===================

async def handle_barge_in():
    """
    Handle user interruption (barge-in).
    
    When user speaks while AI is talking:
    1. Kill current TTS stream immediately
    2. Clear LLM generation queue
    3. Mark interruption timestamp in memory
    4. Wait for speech-end to process new input
    """
    pass


if __name__ == "__main__":
    import uvicorn
    port = int(os.getenv("PORT", 8006))
    uvicorn.run(app, host="0.0.0.0", port=port)
