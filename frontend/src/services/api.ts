import axios from 'axios';
import {
    User,
    Note,
    AuthResponse,
    LoginData,
    RegisterData,
    CreateNoteData,
    UpdateNoteData
} from '../types';

// Configure axios instance
const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
    timeout: 10000,
});

// Debug: Log API configuration
console.log('🔧 API Configuration:', {
    baseURL: api.defaults.baseURL,
    timeout: api.defaults.timeout,
    env: process.env.REACT_APP_API_URL
});

// Auth endpoints
export const authAPI = {
    login: async (data: LoginData): Promise<AuthResponse> => {
        console.log('📡 API Login Request:', {
            url: '/api/auth/login',
            data: data,
            method: 'POST'
        });

        const response = await api.post<AuthResponse>('/api/auth/login', data);

        console.log('📡 API Login Response:', {
            status: response.status,
            data: response.data
        });

        return response.data;
    },

    register: async (data: RegisterData): Promise<AuthResponse> => {
        console.log('📡 API Register Request:', {
            url: '/api/auth/register',
            data: data,
            method: 'POST'
        });

        const response = await api.post<AuthResponse>('/api/auth/register', data);

        console.log('📡 API Register Response:', {
            status: response.status,
            data: response.data
        });

        return response.data;
    },
};

// Notes endpoints
export const notesAPI = {
    getAll: async (): Promise<Note[]> => {
        const response = await api.get<{ notes: Note[] }>('/api/notes');
        return response.data.notes;
    },

    create: async (data: CreateNoteData): Promise<Note> => {
        const response = await api.post<Note>('/api/notes', data);
        return response.data;
    },

    update: async (id: number, data: UpdateNoteData): Promise<Note> => {
        const response = await api.put<Note>(`/api/notes/${id}`, data);
        return response.data;
    },

    delete: async (id: number): Promise<void> => {
        await api.delete(`/api/notes/${id}`);
    },
};

// Axios request interceptor to add Authorization header
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if (token && config.headers) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Axios response interceptor for error handling
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            // Token expired or invalid
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export default api;
