const API_BASE_URL = import.meta.env.VITE_API_URL;

export const ENDPOINTS = {
    // Authentication 
    registerUser: `${API_BASE_URL}/api/users/signup`,
    loginUser: `${API_BASE_URL}/api/users/login`,

    // Products
    getAllProducts: `${API_BASE_URL}/api/product`,
};