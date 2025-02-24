export interface Task {
  id: string;
    title: string;
    description: string;
    status: 'todo' | 'in_progress' | 'completed';
    priority: 'low' | 'medium' | 'high';
    created_at?: string;
  due_date?: string;
  assigned_to?: string;
  created_by: string;
 
  updated_at: string;
  tags?: string[];
}

export interface User {
  id: string;
  email: string;
}

  export interface AIResponse {
    suggestions: string;
    generated_at: string;
  }

  