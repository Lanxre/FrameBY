export const SQUAD_STATUS_LABELS: Record<string, string> = {
  pending: 'Ожидает подтверждения',
  recruitment_open: 'Идёт набор',
  rejected: 'Отклонена',
  closed: 'Закрыта'
}

export const SQUAD_STATUS_COLORS: Record<string, { bg: string; text: string }> = {
  pending: { bg: 'bg-amber-100', text: 'text-amber-600' },
  recruitment_open: { bg: 'bg-green-100', text: 'text-green-600' },
  rejected: { bg: 'bg-red-100', text: 'text-red-600' },
  closed: { bg: 'bg-gray-100', text: 'text-gray-600' }
}

export const SQUAD_STATUS_OPTIONS = [
  { id: 'all', name: 'Все' },
  { id: 'pending', name: 'Ожидает' },
  { id: 'recruitment_open', name: 'Набор' },
  { id: 'rejected', name: 'Отклонена' },
  { id: 'closed', name: 'Закрыта' }
]
