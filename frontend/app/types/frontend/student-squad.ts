export interface StudentSquad {
  id: string
  organizer_id: string
  organizer_name: string
  organizer_role: string
  title: string
  description: string | null
  profile: string | null
  max_participants: number
  current_count: number
  status: string
  status_id: number
  created_at: string
  updated_at: string
}

export interface StudentSquadDetail {
  id: string
  organizer_id: string
  organizer_name: string
  organizer_role: string
  title: string
  description: string | null
  profile: string | null
  max_participants: number
  current_count: number
  status: string
  status_id: number
  created_at: string
  updated_at: string
  participant_ids: string[]
}

export interface CreateStudentSquadData {
  title: string
  description?: string
  profile?: string
  max_participants: number
}

export interface SquadStatus {
  id: number
  name: string
  description: string
}

export const SQUAD_STATUS_LABELS: Record<string, string> = {
  pending: 'Ожидает подтверждения',
  recruitment_open: 'Идёт набор',
  closed: 'Закрыта'
}

export const SQUAD_STATUS_COLORS: Record<string, { bg: string; text: string }> = {
  pending: { bg: 'bg-amber-100', text: 'text-amber-600' },
  recruitment_open: { bg: 'bg-green-100', text: 'text-green-600' },
  closed: { bg: 'bg-gray-100', text: 'text-gray-600' }
}
