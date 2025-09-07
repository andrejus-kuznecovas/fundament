export interface User {
    id: number;
    email: string;
    created_at: string;
}

export interface Note {
    id: number;
    user_id?: number;
    content: string;
    created_at: string;
    updated_at: string;
}

export interface AuthResponse {
    token: string;
    user: User;
}

export interface LoginData {
    email: string;
    password: string;
}

export interface RegisterData {
    email: string;
    password: string;
}

export interface CreateNoteData {
    content: string;
}

export interface UpdateNoteData {
    content: string;
}

export interface ApiError {
    message: string;
    status?: number;
}
