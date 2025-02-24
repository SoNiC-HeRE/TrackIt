
import { Task } from '@/types';

export interface CreateTaskModalProps {
    isOpen: boolean;
    onClose: () => void;
    onCreateTask: (taskData: Partial<Task>) => Promise<void>;
}