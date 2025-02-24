
'use client';

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { taskApi } from '@/services/api';
import { Task } from '@/types';
import { useAuth } from '@/contexts/AuthContext';
import { toast } from 'react-hot-toast';
import {
    PlusIcon,
    FunnelIcon,
    MagnifyingGlassIcon,
} from '@heroicons/react/24/outline';
import LoadingSpinner from '@/components/LoadingSpinner';
import CreateTaskModal from '@/components/CreateTaskModal';
import Sidebar from '@/components/SideBar';

export default function TasksPage() {
    const { user, isLoading: isAuthLoading } = useAuth();
    const [tasks, setTasks] = useState<Task[]>([]);
    const [filteredTasks, setFilteredTasks] = useState<Task[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [searchQuery, setSearchQuery] = useState('');
    const [filters, setFilters] = useState({
        status: 'all',
        priority: 'all',
    });

    useEffect(() => {
        if (!isAuthLoading && user) {
            fetchTasks();
        }
    }, [isAuthLoading, user]);

    useEffect(() => {
        if (!tasks) return;
        
        let filtered = [...tasks];

        
        if (filters.status !== 'all') {
            filtered = filtered.filter(task => task.status === filters.status);
        }

        
        if (filters.priority !== 'all') {
            filtered = filtered.filter(task => task.priority === filters.priority);
        }

        
        if (searchQuery) {
            const query = searchQuery.toLowerCase();
            filtered = filtered.filter(task => 
                task.title.toLowerCase().includes(query) ||
                task.description.toLowerCase().includes(query)
            );
        }

        setFilteredTasks(filtered);
    }, [tasks, filters, searchQuery]);

    const fetchTasks = async () => {
        try {
            setIsLoading(true);
            const fetchedTasks = await taskApi.getTasks();
            setTasks(fetchedTasks || []);
        } catch (error) {
            console.error('Error fetching tasks:', error);
            toast.error('Failed to fetch tasks');
            setTasks([]);
        } finally {
            setIsLoading(false);
        }
    };

    const handleCreateTask = async (taskData: Partial<Task>) => {
        try {
            await taskApi.createTask(taskData);
            await fetchTasks();
            setIsCreateModalOpen(false);
            toast.success('Task created successfully');
        } catch (error) {
            console.error('Error creating task:', error);
            toast.error('Failed to create task');
        }
    };

    const handleUpdateTask = async (taskId: string, updates: Partial<Task>) => {
        try {
            await taskApi.updateTask(taskId, updates);
            await fetchTasks();
            toast.success('Task updated successfully');
        } catch (error) {
            console.error('Error updating task:', error);
            toast.error('Failed to update task');
        }
    };

    const handleDeleteTask = async (taskId: string) => {
        try {
            await taskApi.deleteTask(taskId);
            await fetchTasks();
            toast.success('Task deleted successfully');
        } catch (error) {
            console.error('Error deleting task:', error);
            toast.error('Failed to delete task');
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
            <main className="flex-1 overflow-y-auto p-8">
                <div className="max-w-7xl mx-auto">
                    <div className="flex justify-between items-center mb-8">
                        <h1 className="text-3xl font-bold text-gray-900">Tasks</h1>
                        <button
                            onClick={() => setIsCreateModalOpen(true)}
                            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center space-x-2"
                        >
                            <PlusIcon className="h-5 w-5" />
                            <span>Create Task</span>
                        </button>
                    </div>

                  
                    <div className="bg-white rounded-lg shadow p-4 mb-6">
                        <div className="flex flex-col md:flex-row gap-4">
                            <div className="flex-1 relative">
                                <MagnifyingGlassIcon className="h-5 w-5 absolute left-3 top-3 text-gray-400" />
                                <input
                                    type="text"
                                    placeholder="Search tasks..."
                                    value={searchQuery}
                                    onChange={(e) => setSearchQuery(e.target.value)}
                                    className="pl-10 pr-4 py-2 w-full border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                />
                            </div>
                            <div className="flex gap-4">
                                <select
                                    value={filters.status}
                                    onChange={(e) => setFilters(prev => ({ ...prev, status: e.target.value }))}
                                    className="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                >
                                    <option value="all">All Status</option>
                                    <option value="todo">Todo</option>
                                    <option value="in_progress">In Progress</option>
                                    <option value="completed">Completed</option>
                                </select>
                                <select
                                    value={filters.priority}
                                    onChange={(e) => setFilters(prev => ({ ...prev, priority: e.target.value }))}
                                    className="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                >
                                    <option value="all">All Priority</option>
                                    <option value="low">Low</option>
                                    <option value="medium">Medium</option>
                                    <option value="high">High</option>
                                </select>
                            </div>
                        </div>
                    </div>

                 
                    {isLoading ? (
                        <div className="flex justify-center items-center h-64">
                            <LoadingSpinner size="large" />
                        </div>
                    ) : filteredTasks.length === 0 ? (
                        <div className="text-center py-12 bg-white rounded-lg shadow">
                            <FunnelIcon className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                            <h3 className="text-lg font-medium text-gray-900 mb-2">
                                No tasks found
                            </h3>
                            <p className="text-gray-600">
                                Try adjusting your filters or create a new task
                            </p>
                        </div>
                    ) : (
                        <div className="grid gap-4">
                            {filteredTasks.map((task) => (
                                <TaskCard
                                    key={task.id}
                                    task={task}
                                    onUpdate={handleUpdateTask}
                                    onDelete={handleDeleteTask}
                                />
                            ))}
                        </div>
                    )}
                </div>

               
                <CreateTaskModal
                    isOpen={isCreateModalOpen}
                    onClose={() => setIsCreateModalOpen(false)}
                    onCreateTask={handleCreateTask}
                />
            </main>
        </div>
    );
}

const TaskCard = ({ task, onUpdate, onDelete }: {
    task: Task;
    onUpdate: (id: string, updates: Partial<Task>) => Promise<void>;
    onDelete: (id: string) => Promise<void>;
}) => {
    const statusColors = {
        todo: 'bg-gray-100 text-gray-800',
        in_progress: 'bg-yellow-100 text-yellow-800',
        completed: 'bg-green-100 text-green-800',
    };

    const priorityColors = {
        low: 'bg-blue-100 text-blue-800',
        medium: 'bg-orange-100 text-orange-800',
        high: 'bg-red-100 text-red-800',
    };

    return (
        <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            className="bg-white rounded-lg shadow p-6"
        >
            <div className="flex justify-between items-start mb-4">
                <div>
                    <h3 className="text-lg font-semibold text-gray-900">{task.title}</h3>
                    <p className="text-gray-600 mt-1">{task.description}</p>
                </div>
                <div className="flex gap-2">
                    <span className={`px-3 py-1 rounded-full text-sm ${priorityColors[task.priority]}`}>
                        {task.priority}
                    </span>
                    <span className={`px-3 py-1 rounded-full text-sm ${statusColors[task.status]}`}>
                        {task.status}
                    </span>
                </div>
            </div>
            <div className="flex justify-between items-center">
                <div className="text-sm text-gray-500">
                    Created: {new Date(task.created_at).toLocaleDateString()}
                </div>
                <div className="flex gap-2">
                    <select
                        value={task.status}
                        onChange={(e) => onUpdate(task.id, { status: e.target.value as Task['status'] })}
                        className="text-sm border rounded px-2 py-1 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    >
                        <option value="todo">Todo</option>
                        <option value="in_progress">In Progress</option>
                        <option value="completed">Completed</option>
                    </select>
                    <button
                        onClick={() => onDelete(task.id)}
                        className="text-red-600 hover:text-red-800 text-sm"
                    >
                        Delete
                    </button>
                </div>
            </div>
        </motion.div>
    );
};