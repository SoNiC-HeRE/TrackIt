import { motion } from 'framer-motion';
import { Task } from '@/types';
import { taskApi } from '@/services/api';
import { toast } from 'react-hot-toast';
import { 
  CheckCircleIcon, 
  ClockIcon, 
  ExclamationCircleIcon 
} from '@heroicons/react/24/outline';

interface TaskListProps {
  tasks: Task[];
  onTaskUpdate: () => void;
}

export default function TaskList({ tasks, onTaskUpdate }: TaskListProps) {
  const handleStatusChange = async (taskId: string, newStatus: string) => {
    try {
      await taskApi.updateTask(taskId, { status: newStatus });
      onTaskUpdate();
      toast.success('Task status updated');
    } catch (error) {
      toast.error('Failed to update task status');
    }
  };

  const getPriorityColor = (priority: string) => {
    switch (priority) {
      case 'high':
        return 'text-red-500';
      case 'medium':
        return 'text-yellow-500';
      default:
        return 'text-green-500';
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'completed':
        return <CheckCircleIcon className="h-5 w-5 text-green-500" />;
      case 'in_progress':
        return <ClockIcon className="h-5 w-5 text-yellow-500" />;
      default:
        return <ExclamationCircleIcon className="h-5 w-5 text-gray-500" />;
    }
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {tasks.map((task, index) => (
        <motion.div
          key={task.id}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.3, delay: index * 0.1 }}
          className="bg-white rounded-xl shadow-sm p-6 hover:shadow-md transition-shadow"
        >
          <div className="flex items-start justify-between mb-4">
            <h3 className="text-lg font-semibold text-gray-900">{task.title}</h3>
            <span className={`${getPriorityColor(task.priority)} text-sm font-medium`}>
              {task.priority}
            </span>
          </div>

          <p className="text-gray-600 mb-4 line-clamp-2">{task.description}</p>

          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              {getStatusIcon(task.status)}
              <select
                value={task.status}
                onChange={(e) => handleStatusChange(task.id, e.target.value)}
                className="text-sm border-none bg-transparent focus:ring-0"
              >
                <option value="todo">Todo</option>
                <option value="in_progress">In Progress</option>
                <option value="completed">Completed</option>
              </select>
            </div>

            <span className="text-sm text-gray-500">
              {new Date(task.created_at).toLocaleDateString()}
            </span>
          </div>
        </motion.div>
      ))}
    </div>
  );
}