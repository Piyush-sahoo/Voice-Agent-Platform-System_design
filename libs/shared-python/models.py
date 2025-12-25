"""
Pydantic models for API schemas
"""

from pydantic import BaseModel, Field
from typing import Optional, List
from datetime import datetime
from enum import Enum


class CallDirection(str, Enum):
    INBOUND = "inbound"
    OUTBOUND = "outbound"


class CallStatus(str, Enum):
    QUEUED = "queued"
    RINGING = "ringing"
    IN_PROGRESS = "in_progress"
    COMPLETED = "completed"
    FAILED = "failed"
    NO_ANSWER = "no_answer"
    BUSY = "busy"


class AgentConfig(BaseModel):
    """Voice agent configuration"""
    id: str
    name: str
    system_prompt: str
    first_message: Optional[str] = None
    
    # Voice settings
    voice_provider: str = "elevenlabs"
    voice_id: str = "rachel"
    voice_speed: float = 1.0
    
    # LLM settings
    llm_provider: str = "openai"
    llm_model: str = "gpt-4o-mini"
    llm_temperature: float = 0.7
    llm_max_tokens: int = 500
    
    # STT settings
    stt_provider: str = "deepgram"
    stt_model: str = "nova-2"
    stt_language: str = "en-US"
    
    # Behavior
    interruption_sensitivity: float = 0.8
    end_call_phrases: List[str] = []


class CallDetailRecord(BaseModel):
    """Call Detail Record"""
    id: str
    org_id: str
    agent_id: str
    phone_number: str
    direction: CallDirection
    status: CallStatus
    started_at: datetime
    ended_at: Optional[datetime] = None
    duration_seconds: int = 0
    cost_cents: int = 0
    transcript_url: Optional[str] = None
    recording_url: Optional[str] = None
    summary: Optional[str] = None
    sentiment: Optional[str] = None


class BillingLease(BaseModel):
    """Billing lease for active call"""
    lease_token: str
    call_id: str
    org_id: str
    reserved_cents: int
    expires_at: datetime
