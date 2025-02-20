import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import { Task } from '../../services/tasks/task';
import Loading from '../../components/loading';

const TaskList = () => {
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await Task.getTasks();
        setTasks(res);
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: true,
      render: (text: string, task: Task) => (
        <a href={`/task/${task.id}`}>{text}</a>
      ),
    },
    {
      title: 'Description',
      dataIndex: 'description',
      key: 'description',
      sorter: false,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      sorter: (a: Task, b: Task) => a.status.localeCompare(b.status),
    },
    {
      title: 'Assignee To',
      dataIndex: 'assignedTo',
      key: 'assignedTo',
      sorter: (a: Task, b: Task) => {
        if (a.assignedTo && b.assignedTo) {
          return a.assignedTo.localeCompare(b.assignedTo);
        } else if (a.assignedTo) return -1;
        else if (b.assignedTo) return 1;
        return 0;
      }
    },
  ];

  return (
    <>
      <h1>Tasks List</h1>
      {tasks.length === 0 ? (
        <Loading data="tasks" />
      ) : (
        <Table columns={columns} dataSource={tasks} rowKey="ID"/>
      )}
    </>
  );
};
export default TaskList;
