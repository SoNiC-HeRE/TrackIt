
'use client';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import { toast } from 'react-hot-toast';

interface User {
    id: string;
    name: string;
    email: string;
}

interface AuthContextType {
    user: User | null;
    isLoading: boolean;
    login: (email: string, password: string) => Promise<void>;
    register: (name: string, email: string, password: string) => Promise<void>;
    logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        }
        checkAuth();
    }, []);

    const checkAuth = async () => {
        try {
            const token = localStorage.getItem('token');
            if (token) {
                const response = await axios.get(`${API_URL}/api/me`);
                setUser(response.data.user);
                router.push('/dashboard');
            }
        } catch (error) {
            console.error('Auth check error:', error);
            localStorage.removeItem('token');
            setUser(null);
            router.push('/login');
        } finally {
            setIsLoading(false);
        }
    };

    const register = async (name: string, email: string, password: string) => {
        try {
            const response = await axios.post(`${API_URL}/api/register`, {
                name,
                email,
                password
            });

            const { token, user } = response.data;
            localStorage.setItem('token', token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
            setUser(user);
            toast.success('Registration successful!');
            await router.push('/dashboard');
        } catch (error: any) {
            const errorMessage = error.response?.data?.error || 'Registration failed';
            toast.error(errorMessage);
            throw error;
        }
    };

    const login = async (email: string, password: string) => {
        try {
            const response = await axios.post(`${API_URL}/api/login`, {
                email,
                password
            });

            const { token, user } = response.data;
            localStorage.setItem('token', token);
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
            setUser(user);
            toast.success('Login successful!');
            await router.push('/dashboard');
        } catch (error: any) {
            const errorMessage = error.response?.data?.error || 'Login failed';
            toast.error(errorMessage);
            throw error;
        }
    };

    const logout = async () => {
        try {
            await axios.post(`${API_URL}/api/logout`);
            toast.success('Logged out successfully');
        } catch (error) {
            console.error('Logout error:', error);
        } finally {
            localStorage.removeItem('token');
            delete axios.defaults.headers.common['Authorization'];
            setUser(null);
            router.push('/'); // Redirect to home page instead of login
        }
    };
  

    return (
        <AuthContext.Provider value={{ user, isLoading, login, register, logout }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
}