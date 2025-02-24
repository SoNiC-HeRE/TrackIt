'use client'

import { useState } from 'react'
import { Task } from '@/types'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/outline'

const KanbanCard = ({ task, onUpdate, onDelete }: {
  task: Task
  onUpdate: (taskId: string, updates: Partial<Task>) => void
  onDelete: (taskId: string) => void
}) => {
  const [isEditing, setIsEditing] = useState(false)
  const [editedTitle, setEditedTitle] = useState(task.title)
  const [editedDescription, setEditedDescription] = useState(task.description)

  const handleStatusChange = (newStatus: string) => {
    onUpdate(task.id, { status: newStatus })
  }

  const handleSave = () => {
    onUpdate(task.id, { 
      title: editedTitle,
      description: editedDescription
    })
    setIsEditing(false)
  }

  return (
    <div className="bg-white rounded-lg shadow-sm p-4 border border-gray-200">
      {isEditing ? (
        <div className="space-y-2">
          <input
            type="text"
            value={editedTitle}
            onChange={(e) => setEditedTitle(e.target.value)}
            className="w-full p-2 border rounded text-sm"
          />
          <textarea
            value={editedDescription}
            onChange={(e) => setEditedDescription(e.target.value)}
            className="w-full p-2 border rounded text-sm"
            rows={3}
          />
          <div className="flex gap-2 justify-end">
            <button
              onClick={() => setIsEditing(false)}
              className="px-3 py-1 text-sm text-gray-600 hover:bg-gray-100 rounded"
            >
              Cancel
            </button>
            <button
              onClick={handleSave}
              className="px-3 py-1 text-sm bg-blue-600 text-white rounded hover:bg-blue-700"
            >
              Save
            </button>
          </div>
        </div>
      ) : (
        <>
          <div className="flex justify-between items-start">
            <h4 className="font-medium text-gray-900 text-sm">{task.title}</h4>
            <div className="flex gap-2">
              <button
                onClick={() => setIsEditing(true)}
                className="text-gray-400 hover:text-blue-600"
              >
                <PencilIcon className="h-4 w-4" />
              </button>
              <button
                onClick={() => onDelete(task.id)}
                className="text-gray-400 hover:text-red-600"
              >
                <TrashIcon className="h-4 w-4" />
              </button>
            </div>
          </div>
          {task.description && (
            <p className="text-gray-600 mt-2 text-xs">{task.description}</p>
          )}
          <div className="mt-4">
            <select
              value={task.status}
              onChange={(e) => handleStatusChange(e.target.value)}
              className="w-full p-1 border rounded text-xs bg-gray-50"
            >
              <option value="todo">To Do</option>
              <option value="in_progress">In Progress</option>
              <option value="completed">Completed</option>
            </select>
          </div>
        </>
      )}
    </div>
  )
}

export default KanbanCard