import React from 'react';
import TaskList from './pages/tasks/TaskList';
import {
  Route,
  BrowserRouter as Router,
  Routes, Navigate,
} from 'react-router-dom';
import TaskDetail from './pages/tasks/TaskDetail';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<TaskList />} />
        <Route path="/task/:id" element={<TaskDetail />} />
        <Route path="*" element={<Navigate to="/" replace={true}/>} />
      </Routes>
    </Router>
  );
};
export default App;
