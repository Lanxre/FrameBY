export interface UserEntityResponse {
	message: string;
	user: UserEntity;
}

export interface EnterpriseEntity {
	id: number;
	name: string;
	address: string;
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
	enterprise?: EnterpriseEntity;
}
