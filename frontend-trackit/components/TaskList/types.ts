
import { Task } from '@/types';

export interface TaskListProps {
    tasks: Task[];
    onUpdateTask: (taskId: string, updates: Partial<Task>) => Promise<void>;
    onDeleteTask: (taskId: string) => Promise<void>;
}