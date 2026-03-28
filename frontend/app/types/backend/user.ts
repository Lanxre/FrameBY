export interface UserEntityResponse {
	message: string;
	user: UserEntity;
}

export interface UserEntity {
	id: string;
	email: string;
	login: string;
	role: string;
	created_at: Date;
	updated_at: Date;

	full_name?: string;
	avatar?: string;
	subrole?: string;
	enterprises_id?: number;
}
