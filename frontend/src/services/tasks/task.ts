import { BaseApi } from '../../api/baseApi';
import axios from 'axios';
import {categories, taskStatus} from "../../components/interfaces";

interface TaskApiResponse {
  id: number;
  name: string;
  description: string;
  status: string;
  assigned_to?: string;
  created_at: string;
  completed_at?: string;
  category: string;
}

export class Task extends BaseApi {
  public id!: number;
  public name!: string;
  public description!: string;
  public status!: taskStatus;
  public assignedTo?: string;
  public createdAt?: Date;
  public completedAt?: Date;
  public category?: categories;

  protected static async getEndpoint(endpoint: string): Promise<string> {
    return this.getUrl(`tasks/${endpoint}`);
  }

  private static mapApiResponseToTask(data: TaskApiResponse): Task {
    return {
      id: data.id,
      name: data.name,
      description: data.description,
      status: data.status as taskStatus,
      assignedTo: data.assigned_to,
      createdAt: new Date(data.created_at),
      completedAt: data.completed_at ? new Date(data.completed_at) : undefined,
      category: data.category as categories,
    };
  }

  public static async getTasks(): Promise<Task[]> {
    const res = await axios.get(await Task.getEndpoint(`all`));

    if (!res.data) throw new Error('unable to get tasks');
    return res.data.map((task: TaskApiResponse) => Task.mapApiResponseToTask(task));
  }

  public static async getTaskById(id: string): Promise<Task> {
    const res = await axios.get(await Task.getEndpoint(id));

    if (!res.data) throw new Error('unable to get tasks');
    return Task.mapApiResponseToTask(res.data);
  }
}
