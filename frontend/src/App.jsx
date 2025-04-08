import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './App.css';

// Заглушки для компонентов, которые будут созданы позже
const Home = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">Главная страница MultyLink</h1></div>;
const Login = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">Вход в систему</h1></div>;
const Register = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">Регистрация</h1></div>;
const Dashboard = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">Личный кабинет</h1></div>;
const PublicPage = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">Публичная страница пользователя</h1></div>;
const NotFound = () => <div className="p-8 text-center"><h1 className="text-3xl font-bold">404 - Страница не найдена</h1></div>;

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  
  // Проверка аутентификации при загрузке приложения
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      // Здесь можно добавить проверку валидности токена на сервере
      setIsAuthenticated(true);
    }
  }, []);

  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        {/* Здесь будет компонент Header */}
        
        <main className="container mx-auto py-4">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/:username" element={<PublicPage />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </main>
        
        {/* Здесь будет компонент Footer */}
      </div>
    </Router>
  );
}

export default App;
