'use client'

import { motion } from 'framer-motion'
import KanbanCard from './KanbanCard'
import { Task } from '@/types'

const statuses = ['todo', 'in_progress', 'completed']

const statusLabels = {
  todo: 'To Do',
  in_progress: 'In Progress',
  completed: 'Completed'
}

const statusColors = {
  todo: 'bg-blue-100',
  in_progress: 'bg-yellow-100',
  completed: 'bg-green-100'
}

const KanbanBoard = ({ tasks, onUpdateTask, onDeleteTask }: {
  tasks: Task[]
  onUpdateTask: (taskId: string, updates: Partial<Task>) => void
  onDeleteTask: (taskId: string) => void
}) => {
  return (
    <div className="flex gap-4 overflow-x-auto pb-4 min-h-[32rem]">
      {statuses.map((status) => {
        const statusTasks = tasks.filter(task => task.status === status)
        return (
          <div 
            key={status}
            className={`flex-1 min-w-[300px] rounded-lg p-4 ${statusColors[status]}`}
          >
            <h3 className="font-semibold text-lg mb-4 text-gray-700">
              {statusLabels[status]} ({statusTasks.length})
            </h3>
            <div className="space-y-4">
              {statusTasks.map(task => (
                <motion.div
                  key={task.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                >
                  <KanbanCard
                    task={task}
                    onUpdate={onUpdateTask}
                    onDelete={onDeleteTask}
                  />
                </motion.div>
              ))}
            </div>
          </div>
        )
      })}
    </div>
  )
}

export default KanbanBoard
