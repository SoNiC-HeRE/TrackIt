'use client';

import { useState } from 'react';
import { motion } from 'framer-motion';
import { useAuth } from '@/contexts/AuthContext';
import { taskApi } from '@/services/api';
import { toast } from 'react-hot-toast';
import {
    PaperAirplaneIcon,
    SparklesIcon,
} from '@heroicons/react/24/outline';
import LoadingSpinner from '@/components/LoadingSpinner';
import Sidebar from '@/components/SideBar';

interface Message {
    id: string;
    content: string;
    isAI: boolean;
    timestamp: Date;
}

export default function AIAssistantPage() {
    const { user, isLoading: isAuthLoading } = useAuth();
    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleSendMessage = async () => {
        if (!input.trim() || isLoading) return;
    
        const userMessage: Message = {
            id: Date.now().toString(),
            content: input,
            isAI: false,
            timestamp: new Date(),
        };
    
        setMessages(prev => [...prev, userMessage]);
        setInput('');
        setIsLoading(true);
    
        try {
            const response = await taskApi.getAISuggestions(input);
            
            const aiMessage: Message = {
                id: (Date.now() + 1).toString(),
                content: response.suggestions,
                isAI: true,
                timestamp: new Date(response.timestamp),
            };
    
            setMessages(prev => [...prev, aiMessage]);
        } catch (error: any) {
            console.error('Error getting AI response:', error);
            const errorMessage = error.response?.data?.error || 'Failed to get AI response';
            
            const aiMessage: Message = {
                id: (Date.now() + 1).toString(),
                content: `Error: ${errorMessage}. Please try again.`,
                isAI: true,
                timestamp: new Date(),
            };
    
            setMessages(prev => [...prev, aiMessage]);
            toast.error(errorMessage);
        } finally {
            setIsLoading(false);
        }
    };

    if (isAuthLoading) {
        return (
            <div className="flex justify-center items-center h-screen">
                <LoadingSpinner size="large" />
            </div>
        );
    }

    if (!user) {
        return null;
    }

    return (
        <div className="flex h-screen bg-gray-50">
            <Sidebar />
            <main className="flex-1 overflow-hidden flex flex-col">
                <div className="p-8 border-b">
                    <h1 className="text-3xl font-bold text-gray-900">AI Assistant</h1>
                    <p className="text-gray-600 mt-2">
                        Get AI-powered suggestions and help with your tasks
                    </p>
                </div>

              
                <div className="flex-1 overflow-y-auto p-8 space-y-6">
                    {messages.map((message) => (
                        <motion.div
                            key={message.id}
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            className={`flex ${message.isAI ? 'justify-start' : 'justify-end'}`}
                        >
                            <div
                                className={`max-w-2xl rounded-lg p-4 ${
                                    message.isAI
                                        ? 'bg-white shadow-sm'
                                        : 'bg-blue-600 text-white'
                                }`}
                            >
                                {message.isAI && (
                                    <div className="flex items-center gap-2 mb-2 text-blue-600">
                                        <SparklesIcon className="h-5 w-5" />
                                        <span className="font-medium">AI Assistant</span>
                                    </div>
                                )}
                                <p className="whitespace-pre-wrap">{message.content}</p>
                                <div className={`text-xs mt-2 ${message.isAI ? 'text-gray-500' : 'text-blue-200'}`}>
                                    {message.timestamp.toLocaleTimeString()}
                                </div>
                            </div>
                        </motion.div>
                    ))}
                    {isLoading && (
                        <div className="flex justify-start">
                            <div className="bg-white rounded-lg p-4 shadow-sm">
                                <LoadingSpinner size="small" />
                            </div>
                        </div>
                    )}
                </div>

                
                <div className="p-4 border-t bg-white">
                    <div className="max-w-4xl mx-auto flex gap-4">
                        <input
                            type="text"
                            value={input}
                            onChange={(e) => setInput(e.target.value)}
                            onKeyPress={(e) => e.key === 'Enter' && handleSendMessage()}
                            placeholder="Ask for task suggestions..."
                            className="flex-1 border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                            disabled={isLoading}
                        />
                        <button
                            onClick={handleSendMessage}
                            disabled={isLoading || !input.trim()}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors disabled:bg-blue-300 disabled:cursor-not-allowed"
                        >
                            <PaperAirplaneIcon className="h-5 w-5" />
                        </button>
                    </div>
                </div>
            </main>
        </div>
    );
}