export interface LoadingProps {
  data: string;
}

export enum taskStatus {
  toDo = 'To Do',
  inProgress = 'In Progress',
  completed = 'Completed',
}

export enum categories {
  content = 'Content',
  design = 'Design',
  bug = 'Bug',
  develop = 'Development',
}