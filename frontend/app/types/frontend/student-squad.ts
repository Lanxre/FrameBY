export interface SquadOrganizer {
  id: string;
  name: string;
  role: string;
  position: string;
  phone: string;
}

export interface StudentSquad {
  id: string;
  organizer: SquadOrganizer;
  title: string;
  description: string | null;
  profile: string | null;
  max_participants: number;
  current_count: number;
  status: string;
  status_id: number;
  created_at: string;
  updated_at: string;
}

export interface StudentSquadDetail extends StudentSquad {
  participant_ids: string[];
}

export interface CreateStudentSquadData {
  title: string;
  description?: string;
  profile?: string;
  max_participants: number;
}

export interface SquadStatus {
  id: number;
  name: string;
  description: string;
}
