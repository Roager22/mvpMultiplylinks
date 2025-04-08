import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

export default function Dashboard() {
  const [multiLinks, setMultiLinks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    // Проверка аутентификации
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
      return;
    }

    // Установка токена для запросов
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    // Загрузка мультиссылок пользователя
    fetchMultiLinks();
  }, [navigate]);

  const fetchMultiLinks = async () => {
    try {
      setLoading(true);
      const response = await axios.get('/api/multilinks');
      setMultiLinks(response.data.multilinks || []);
      setError('');
    } catch (err) {
      setError('Не удалось загрузить ваши мультиссылки. Пожалуйста, попробуйте позже.');
      console.error('Error fetching multilinks:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateMultiLink = () => {
    navigate('/dashboard/create');
  };

  return (
    <div className="py-10">
      <header>
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold leading-tight tracking-tight text-gray-900">Личный кабинет</h1>
        </div>
      </header>
      <main>
        <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
          {/* Верхняя панель с кнопкой создания */}
          <div className="px-4 py-5 sm:px-6 flex justify-between items-center">
            <div>
              <h3 className="text-lg font-medium leading-6 text-gray-900">Мои мультиссылки</h3>
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                Управляйте всеми вашими мультиссылками в одном месте
              </p>
            </div>
            <button
              type="button"
              onClick={handleCreateMultiLink}
              className="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
              Создать мультиссылку
            </button>
          </div>

          {/* Сообщение об ошибке */}
          {error && (
            <div className="rounded-md bg-red-50 p-4 mx-4 sm:mx-0 mb-4">
              <div className="flex">
                <div className="ml-3">
                  <h3 className="text-sm font-medium text-red-800">Ошибка</h3>
                  <div className="mt-2 text-sm text-red-700">
                    <p>{error}</p>
                  </div>
                </div>
              </div>
            </div>
          )}

          {/* Список мультиссылок */}
          <div className="mt-4 overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg mx-4 sm:mx-0">
            {loading ? (
              <div className="p-10 text-center">
                <p className="text-gray-500">Загрузка...</p>
              </div>
            ) : multiLinks.length === 0 ? (
              <div className="p-10 text-center">
                <p className="text-gray-500">У вас пока нет мультиссылок. Создайте первую!</p>
              </div>
            ) : (
              <table className="min-w-full divide-y divide-gray-300">
                <thead className="bg-gray-50">
                  <tr>
                    <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">
                      Название
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Slug
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Статус
                    </th>
                    <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">
                      Создана
                    </th>
                    <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-6">
                      <span className="sr-only">Действия</span>
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200 bg-white">
                  {multiLinks.map((link) => (
                    <tr key={link.id}>
                      <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">
                        {link.title}
                      </td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                        {link.slug}
                      </td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                        <span className={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${link.isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}`}>
                          {link.isActive ? 'Активна' : 'Неактивна'}
                        </span>
                      </td>
                      <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                        {new Date(link.createdAt).toLocaleDateString()}
                      </td>
                      <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                        <Link to={`/dashboard/edit/${link.id}`} className="text-indigo-600 hover:text-indigo-900 mr-4">
                          Редактировать
                        </Link>
                        <Link to={`/${link.slug}`} target="_blank" className="text-green-600 hover:text-green-900">
                          Просмотр
                        </Link>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
      </main>
    </div>
  );
}