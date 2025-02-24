
'use client';

import { motion, AnimatePresence } from 'framer-motion';

import { Task } from '@/types';

interface TaskListProps {
    tasks: Task[];
    onUpdateTask: (taskId: string, updates: Partial<Task>) => Promise<void>;
    onDeleteTask: (taskId: string) => Promise<void>;
}



const TaskList: React.FC<TaskListProps> = ({ tasks, onUpdateTask, onDeleteTask }) => {
    return (
        <div className="space-y-4">
            <AnimatePresence>
                {tasks.map((task) => (
                    <motion.div
                        key={task.id}
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        exit={{ opacity: 0, y: -20 }}
                        className="bg-white rounded-lg shadow p-4"
                    >
                        
                        <div className="flex justify-between items-center">
                            <div>
                                <h3 className="text-lg font-medium">{task.title}</h3>
                                <p className="text-gray-600">{task.description}</p>
                            </div>
                            <div className="flex space-x-2">
                                <select
                                    value={task.status}
                                    onChange={(e) => onUpdateTask(task.id, { status: e.target.value as Task['status'] })}
                                    className="rounded border border-gray-300 px-2 py-1"
                                >
                                    <option value="todo">Todo</option>
                                    <option value="in_progress">In Progress</option>
                                    <option value="completed">Completed</option>
                                </select>
                                <button
                                    onClick={() => onDeleteTask(task.id)}
                                    className="text-red-600 hover:text-red-800"
                                >
                                    Delete
                                </button>
                            </div>
                        </div>
                    </motion.div>
                ))}
            </AnimatePresence>
        </div>
    );
};

export default TaskList;