export interface Chat {
	id: string;
	type: string;
	name?: string;
	avatar?: string;
	last_message?: string;
	last_message_at?: string;
	unread_count?: number;
	created_at: string;
}

export interface ChatResponse {
	id: string;
	type: string;
	name?: string;
	avatar?: string;
	last_message?: string;
	last_message_at?: string;
	unread_count?: number;
	created_at: string;
}

export interface ChatMember {
	id: string;
	chat_id: string;
	user_id: string;
	role: string;
	joined_at: string;
	user_name: string;
	user_login: string;
	user_avatar?: string;
}

export interface ChatMessage {
	id: string;
	chat_id: string;
	user_id: string;
	user_name: string;
	user_avatar?: string;
	type: string;
	content: string;
	metadata?: Record<string, any>;
	created_at: string;
}

export interface CreateChatRequest {
	type: string;
	name?: string;
	users: string[];
}

export interface SendMessageRequest {
	type: string;
	content: string;
	metadata?: Record<string, any>;
}

export interface WebSocketMessage {
	type: string;
	payload: any;
}

export interface SearchUser {
	user_id: string;
	email: string;
	login: string;
	role: string;
	avatar?: string;
	full_name?: string;
	profile?: {
		specialty?: string;
		grade?: number;
	};
}
