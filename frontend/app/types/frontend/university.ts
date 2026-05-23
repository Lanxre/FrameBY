export interface StudentSquadInfo {
	id: string;
	name: string;
	description?: string;
	joined_at: string;
}

export interface WorkInfo {
	id: string;
	title: string;
	company: string;
	position: string;
	status: string;
	started_at: string;
}

export interface StudentEmploymentInfo {
	student: {
		user_id: string;
		email: string;
		login: string;
		role: string;
		avatar?: string;
		profile: {
			full_name?: string;
			specialty?: string;
			grade?: number;
			position?: string;
			phone?: string;
		};
		updated_at: string;
	};
	student_squads: StudentSquadInfo[];
	works: WorkInfo[];
}

export interface StudentEmploymentListResponse {
	students: StudentEmploymentInfo[];
	total: number;
	limit: number;
	offset: number;
}

export interface UniversityDepartment {
	id: string;
  university_name: string;
  department_name: string;
  address?: string;
}