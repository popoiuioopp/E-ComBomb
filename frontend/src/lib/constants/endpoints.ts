const API_BASE_URL = import.meta.env.VITE_API_URL;

export const ENDPOINTS = {
    // Authentication 
    registerUser: `${API_BASE_URL}/users/signup`,
    loginUser: `${API_BASE_URL}/users/login`
};