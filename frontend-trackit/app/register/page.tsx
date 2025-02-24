'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { motion } from 'framer-motion';
import { toast } from 'react-hot-toast';
import { useAuth } from '@/contexts/AuthContext';
import { 
    EnvelopeIcon, 
    LockClosedIcon, 
    UserIcon,
    EyeIcon,
    EyeSlashIcon,
    ArrowLeftIcon
} from '@heroicons/react/24/outline';

interface RegisterFormData {
    name: string;
    email: string;
    password: string;
}

export default function RegisterPage() {
    const router = useRouter();
    const { register, user, isLoading: authLoading } = useAuth();
    const [isLoading, setIsLoading] = useState(false);
    const [showPassword, setShowPassword] = useState(false);
    const [formData, setFormData] = useState<RegisterFormData>({
        name: '',
        email: '',
        password: '',
    });
    const [errors, setErrors] = useState<Partial<RegisterFormData>>({});

    
    useEffect(() => {
        
        if (user && !authLoading) {
            router.push('/dashboard');
        }
    }, [user, authLoading, router]);

    const validateForm = (): boolean => {
        const newErrors: Partial<RegisterFormData> = {};

        
        if (!formData.name.trim()) {
            newErrors.name = 'Name is required';
        } else if (formData.name.length < 2) {
            newErrors.name = 'Name must be at least 2 characters';
        }

        
        if (!formData.email.trim()) {
            newErrors.email = 'Email is required';
        } else if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i.test(formData.email)) {
            newErrors.email = 'Invalid email address';
        }

        
        if (!formData.password) {
            newErrors.password = 'Password is required';
        } else if (formData.password.length < 8) {
            newErrors.password = 'Password must be at least 8 characters';
        } else if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/.test(formData.password)) {
            newErrors.password = 'Password must contain uppercase, lowercase, and numbers';
        }

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };


   

const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
        return;
    }

    setIsLoading(true);

    try {
        await register(formData.name, formData.email, formData.password);
        
    } catch (error: any) {
        console.error('Registration error:', error);
        toast.error(error.response?.data?.error || 'Registration failed');
    } finally {
        setIsLoading(false);
    }
};
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value,
        }));
       
        if (errors[name as keyof RegisterFormData]) {
            setErrors(prev => ({
                ...prev,
                [name]: '',
            }));
        }
    };

   
    if (authLoading) {
        return (
            <div className="min-h-screen flex items-center justify-center">
                <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
            </div>
        );
    }

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-50 py-12 px-4 sm:px-6 lg:px-8">
            <motion.div
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5 }}
                className="max-w-md w-full space-y-8 bg-white p-8 rounded-xl shadow-lg"
            >
                
                <div className="text-center">
                    <motion.div
                        initial={{ scale: 0 }}
                        animate={{ scale: 1 }}
                        transition={{ delay: 0.2 }}
                        className="mx-auto h-12 w-12 bg-blue-600 text-white flex items-center justify-center rounded-full"
                    >
                        <UserIcon className="h-6 w-6" />
                    </motion.div>
                    <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
                        Create your account
                    </h2>
                    <p className="mt-2 text-sm text-gray-600">
                        Already have an account?{' '}
                        <Link 
                            href="/login"
                            className="font-medium text-blue-600 hover:text-blue-500 transition-colors"
                        >
                            Sign in
                        </Link>
                    </p>
                </div>

                
                <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
                    
                    <div>
                        <label htmlFor="name" className="block text-sm font-medium text-gray-700">
                            Full Name
                        </label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <UserIcon className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                id="name"
                                name="name"
                                type="text"
                                autoComplete="name"
                                required
                                value={formData.name}
                                onChange={handleChange}
                                className={`block w-full pl-10 pr-3 py-2 border ${
                                    errors.name ? 'border-red-300 ring-red-500' : 'border-gray-300'
                                } rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500`}
                                placeholder="John Doe"
                            />
                        </div>
                        {errors.name && (
                            <motion.p
                                initial={{ opacity: 0, y: -10 }}
                                animate={{ opacity: 1, y: 0 }}
                                className="mt-2 text-sm text-red-600"
                            >
                                {errors.name}
                            </motion.p>
                        )}
                    </div>

                    
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                            Email address
                        </label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <EnvelopeIcon className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                id="email"
                                name="email"
                                type="email"
                                autoComplete="email"
                                required
                                value={formData.email}
                                onChange={handleChange}
                                className={`block w-full pl-10 pr-3 py-2 border ${
                                    errors.email ? 'border-red-300 ring-red-500' : 'border-gray-300'
                                } rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500`}
                                placeholder="you@example.com"
                            />
                        </div>
                        {errors.email && (
                            <motion.p
                                initial={{ opacity: 0, y: -10 }}
                                animate={{ opacity: 1, y: 0 }}
                                className="mt-2 text-sm text-red-600"
                            >
                                {errors.email}
                            </motion.p>
                        )}
                    </div>

                    
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <LockClosedIcon className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                id="password"
                                name="password"
                                type={showPassword ? "text" : "password"}
                                autoComplete="new-password"
                                required
                                value={formData.password}
                                onChange={handleChange}
                                className={`block w-full pl-10 pr-10 py-2 border ${
                                    errors.password ? 'border-red-300 ring-red-500' : 'border-gray-300'
                                } rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500`}
                                placeholder="••••••••"
                            />
                            <button
                                type="button"
                                onClick={() => setShowPassword(!showPassword)}
                                className="absolute inset-y-0 right-0 pr-3 flex items-center"
                            >
                                {showPassword ? (
                                    <EyeSlashIcon className="h-5 w-5 text-gray-400 hover:text-gray-600" />
                                ) : (
                                    <EyeIcon className="h-5 w-5 text-gray-400 hover:text-gray-600" />
                                )}
                            </button>
                        </div>
                        {errors.password && (
                            <motion.p
                                initial={{ opacity: 0, y: -10 }}
                                animate={{ opacity: 1, y: 0 }}
                                className="mt-2 text-sm text-red-600"
                            >
                                {errors.password}
                            </motion.p>
                        )}
                    </div>

                    
                    <div>
                        <button
                            type="submit"
                            disabled={isLoading}
                            className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-blue-300 disabled:cursor-not-allowed transition-colors"
                        >
                            {isLoading ? (
                                <motion.div
                                    initial={{ rotate: 0 }}
                                    animate={{ rotate: 360 }}
                                    transition={{ duration: 1, repeat: Infinity, ease: "linear" }}
                                    className="h-5 w-5 border-2 border-white border-t-transparent rounded-full"
                                />
                            ) : (
                                'Create Account'
                            )}
                        </button>
                    </div>
                </form>

                
                <div className="mt-6">
                    <p className="text-xs text-gray-600 text-center">
                        By creating an account, you agree to our{' '}
                        <Link href="/terms" className="text-blue-600 hover:text-blue-500">
                            Terms of Service
                        </Link>{' '}
                        and{' '}
                        <Link href="/privacy" className="text-blue-600 hover:text-blue-500">
                            Privacy Policy
                        </Link>
                    </p>
                </div>

               
                <div className="mt-4 text-center">
                    <Link
                        href="/"
                        className="inline-flex items-center text-sm text-gray-500 hover:text-gray-700"
                    >
                        <ArrowLeftIcon className="h-4 w-4 mr-1" />
                        Back to Home
                    </Link>
                </div>
            </motion.div>
        </div>
    );
}