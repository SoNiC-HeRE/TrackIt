package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Email    string            `bson:"email" json:"email"`
    Password string            `bson:"password" json:"-"`
}

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string            `bson:"title" json:"title"`
    Description string            `bson:"description" json:"description"`
    Status      string            `bson:"status" json:"status"`
    Priority    string            `bson:"priority" json:"priority"`
    DueDate     *time.Time        `bson:"due_date,omitempty" json:"due_date,omitempty"`
    AssignedTo  primitive.ObjectID `bson:"assigned_to,omitempty" json:"assigned_to,omitempty"`
    CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
    CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time         `bson:"updated_at" json:"updated_at"`
    Tags        []string          `bson:"tags,omitempty" json:"tags,omitempty"`
}

type AITaskSuggestion struct {
    TaskID      primitive.ObjectID `bson:"task_id" json:"task_id"`
    Suggestion  string            `bson:"suggestion" json:"suggestion"`
    GeneratedAt time.Time         `bson:"generated_at" json:"generated_at"`
}


type TaskTemplate struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string            `bson:"name" json:"name"`
    Description string            `bson:"description" json:"description"`
    Priority    string            `bson:"priority" json:"priority"`
    Tags        []string          `bson:"tags" json:"tags"`
    CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
    CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
}

type RecurringTask struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    TaskTemplate primitive.ObjectID `bson:"task_template" json:"task_template"`
    Frequency    string            `bson:"frequency" json:"frequency"` // daily, weekly, monthly
    NextDue      time.Time         `bson:"next_due" json:"next_due"`
    LastCreated  time.Time         `bson:"last_created" json:"last_created"`
    Active       bool              `bson:"active" json:"active"`
}