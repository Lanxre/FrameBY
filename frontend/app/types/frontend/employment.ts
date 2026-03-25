export interface EmploymentRequest {
  id: string
  enterprise_id: string
  enterprise_name: string
  university_department_id: string
  university_department_name: string
  title: string
  description: string | null
  requirements: string | null
  salary: string | null
  schedule: string | null
  location: string | null
  employment_type: string
  status: string
  status_id: number
  created_at: string
  updated_at: string
}

export interface EmploymentRequestDetail extends EmploymentRequest {
  participant_ids: string[]
  participants_count: number
}

export interface EmploymentParticipant {
  id: string
  request_id: string
  user_id: string
  user_name: string
  user_role: string
  status: string
  status_id: number
  applied_at: string
  updated_at: string
}

export interface CreateEmploymentRequestData {
  enterprise_id: string
  university_department_id: string
  title: string
  description?: string
  requirements?: string
  salary?: string
  schedule?: string
  location?: string
  employment_type: string
}

export interface ApplyToRequestData {
  request_id: string
}

export interface EmploymentStatus {
  id: number
  name: string
  description: string
}

export interface ParticipantStatus {
  id: number
  name: string
  description: string
}

export const EMPLOYMENT_STATUS_LABELS: Record<string, string> = {
  pending: 'Ожидает подтверждения',
  approved: 'Подтверждена',
  closed: 'Закрыта'
}

export const EMPLOYMENT_STATUS_COLORS: Record<string, { bg: string; text: string }> = {
  pending: { bg: 'bg-amber-100', text: 'text-amber-600' },
  approved: { bg: 'bg-green-100', text: 'text-green-600' },
  closed: { bg: 'bg-gray-100', text: 'text-gray-600' }
}

export const PARTICIPANT_STATUS_LABELS: Record<string, string> = {
  applied: 'Подана',
  accepted: 'Принята',
  rejected: 'Отклонена',
  contracted: 'Трудоустроен'
}

export const PARTICIPANT_STATUS_COLORS: Record<string, { bg: string; text: string }> = {
  applied: { bg: 'bg-blue-100', text: 'text-blue-600' },
  accepted: { bg: 'bg-green-100', text: 'text-green-600' },
  rejected: { bg: 'bg-red-100', text: 'text-red-600' },
  contracted: { bg: 'bg-emerald-100', text: 'text-emerald-600' }
}
