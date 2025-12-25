import { z } from 'zod';

// ===================
// USER SCHEMAS
// ===================

export const UserSchema = z.object({
    id: z.string().uuid(),
    email: z.string().email(),
    name: z.string().min(1),
    avatar_url: z.string().url().optional(),
    created_at: z.string().datetime(),
});

export const CreateUserSchema = z.object({
    email: z.string().email(),
    password: z.string().min(8),
    name: z.string().min(1),
});

// ===================
// ORGANIZATION SCHEMAS
// ===================

export const OrganizationSchema = z.object({
    id: z.string().uuid(),
    name: z.string().min(1),
    slug: z.string().min(1),
    plan: z.enum(['free', 'pro', 'enterprise']),
    created_at: z.string().datetime(),
});

export const CreateOrganizationSchema = z.object({
    name: z.string().min(1),
});

// ===================
// AGENT SCHEMAS
// ===================

export const VoiceConfigSchema = z.object({
    provider: z.enum(['elevenlabs', 'cartesia', 'openai']),
    voice_id: z.string(),
    speed: z.number().min(0.5).max(2.0).default(1.0),
    stability: z.number().min(0).max(1).default(0.5),
});

export const LLMConfigSchema = z.object({
    provider: z.enum(['openai', 'anthropic', 'groq']),
    model: z.string(),
    temperature: z.number().min(0).max(2).default(0.7),
    max_tokens: z.number().min(1).max(4096).default(500),
});

export const STTConfigSchema = z.object({
    provider: z.enum(['deepgram', 'assemblyai', 'whisper']),
    model: z.string().default('nova-2'),
    language: z.string().default('en-US'),
});

export const AgentSchema = z.object({
    id: z.string().uuid(),
    org_id: z.string().uuid(),
    name: z.string().min(1),
    system_prompt: z.string(),
    first_message: z.string().optional(),
    voice: VoiceConfigSchema,
    llm: LLMConfigSchema,
    stt: STTConfigSchema,
    interruption_sensitivity: z.number().min(0).max(1).default(0.8),
    end_call_phrases: z.array(z.string()).default([]),
    is_active: z.boolean().default(true),
    created_at: z.string().datetime(),
});

export const CreateAgentSchema = z.object({
    name: z.string().min(1),
    system_prompt: z.string(),
    first_message: z.string().optional(),
    voice: VoiceConfigSchema.optional(),
    llm: LLMConfigSchema.optional(),
    stt: STTConfigSchema.optional(),
});

// ===================
// CALL SCHEMAS
// ===================

export const CallDirectionEnum = z.enum(['inbound', 'outbound']);
export const CallStatusEnum = z.enum([
    'queued',
    'ringing',
    'in_progress',
    'completed',
    'failed',
    'no_answer',
    'busy',
]);

export const CallSchema = z.object({
    id: z.string().uuid(),
    org_id: z.string().uuid(),
    agent_id: z.string().uuid(),
    phone_number: z.string(),
    direction: CallDirectionEnum,
    status: CallStatusEnum,
    started_at: z.string().datetime(),
    ended_at: z.string().datetime().optional(),
    duration_seconds: z.number().int().min(0),
    cost_cents: z.number().int().min(0),
    transcript_url: z.string().url().optional(),
    recording_url: z.string().url().optional(),
    summary: z.string().optional(),
    sentiment: z.enum(['positive', 'negative', 'neutral']).optional(),
});

// Type exports
export type User = z.infer<typeof UserSchema>;
export type CreateUser = z.infer<typeof CreateUserSchema>;
export type Organization = z.infer<typeof OrganizationSchema>;
export type CreateOrganization = z.infer<typeof CreateOrganizationSchema>;
export type Agent = z.infer<typeof AgentSchema>;
export type CreateAgent = z.infer<typeof CreateAgentSchema>;
export type Call = z.infer<typeof CallSchema>;
