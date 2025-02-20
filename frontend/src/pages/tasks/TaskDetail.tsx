import React, { useEffect, useState } from 'react';
import { Task } from '../../services/tasks/task';
import { useParams } from 'react-router-dom';
import Loading from '../../components/loading';
import { taskStatus } from '../../components/interfaces';
import './TaskDetails.css';

const TaskDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [task, setTask] = useState<Task>();

  useEffect(() => {
    const fetchData = async (id: string) => {
      try {
        const res = await Task.getTaskById(id);
        setTask(res);
      } catch (error) {
        console.log(error);
      }
    };
    if (id) {
      fetchData(id);
    }
  }, [id]);

  return (
    <>
    {!task ? (
        <Loading data="task" />
      ) : (
        <>
          <h1>{task.name}</h1>
          <p>{task.category} - created on: {task.createdAt!.toDateString()}</p>
          <p>{task.description}</p>
          <div>
            {task.status === taskStatus.completed ? (
              <p>{task.status} on {task.completedAt!.toDateString()}</p>
            ) : (
              <p>{task.status} </p>
            )}
          </div>
          <p>{task.assignedTo}</p>
        </>
      )}
    </>
  );
};

export default TaskDetail;
