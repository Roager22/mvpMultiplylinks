import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { Bar, Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  LineElement,
  PointElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

// Регистрируем необходимые компоненты Chart.js
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  LineElement,
  PointElement,
  Title,
  Tooltip,
  Legend
);

export default function LinkMetrics() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [multiLink, setMultiLink] = useState(null);
  const [buttons, setButtons] = useState([]);
  const [metrics, setMetrics] = useState([]);
  const [dailyClicks, setDailyClicks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [dateRange, setDateRange] = useState('week'); // 'week', 'month', 'year'

  useEffect(() => {
    // Проверка аутентификации
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
      return;
    }

    // Установка токена для запросов
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    // Загрузка данных
    fetchData();
  }, [id, dateRange, navigate]);

  const fetchData = async () => {
    try {
      setLoading(true);
      
      // Получаем информацию о мультиссылке
      const multiLinkResponse = await axios.get(`/api/multilinks/${id}`);
      setMultiLink(multiLinkResponse.data.multilink);
      
      // Получаем кнопки мультиссылки
      const buttonsResponse = await axios.get(`/api/multilinks/${id}/buttons`);
      setButtons(buttonsResponse.data.buttons || []);
      
      // Получаем метрики для всех кнопок
      const metricsResponse = await axios.get(`/api/multilinks/${id}/metrics`);
      setMetrics(metricsResponse.data.metrics || []);
      
      // Получаем статистику по дням
      const dailyResponse = await axios.get(`/api/multilinks/${id}/metrics/daily?range=${dateRange}`);
      setDailyClicks(dailyResponse.data.dailyClicks || []);
      
      setError('');
    } catch (err) {
      setError(err.response?.data?.error || 'Произошла ошибка при загрузке данных. Пожалуйста, попробуйте снова.');
    } finally {
      setLoading(false);
    }
  };

  // Подготовка данных для графика по кнопкам
  const prepareButtonsChartData = () => {
    const labels = buttons.map(button => button.title);
    const data = buttons.map(button => {
      const buttonMetrics = metrics.find(m => m.linkButtonId === button.id);
      return buttonMetrics ? buttonMetrics.clicks : 0;
    });
    
    return {
      labels,
      datasets: [
        {
          label: 'Количество кликов',
          data,
          backgroundColor: buttons.map(button => button.color || '#3B82F6'),
          borderColor: buttons.map(button => button.color || '#3B82F6'),
          borderWidth: 1,
        },
      ],
    };
  };

  // Подготовка данных для графика по дням
  const prepareDailyChartData = () => {
    const labels = dailyClicks.map(item => item.date);
    const data = dailyClicks.map(item => item.clicks);
    
    return {
      labels,
      datasets: [
        {
          label: 'Клики по дням',
          data,
          fill: false,
          backgroundColor: 'rgba(75, 192, 192, 0.2)',
          borderColor: 'rgba(75, 192, 192, 1)',
          tension: 0.1,
        },
      ],
    };
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'Статистика кликов',
      },
    },
  };

  if (loading && !multiLink) {
    return (
      <div className="py-10">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <p className="text-gray-500">Загрузка...</p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="py-10">
      <header>
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <h1 className="text-3xl font-bold leading-tight tracking-tight text-gray-900">Статистика мультиссылки</h1>
          {multiLink && (
            <p className="mt-2 text-gray-500">{multiLink.title}</p>
          )}
        </div>
      </header>
      <main>
        <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
          {error && (
            <div className="mb-4 rounded-md bg-red-50 p-4">
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
          
          {/* Фильтры по времени */}
          <div className="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6 mb-5">
            <div className="flex items-center justify-between">
              <h2 className="text-lg font-medium text-gray-900">Период</h2>
              <div className="flex space-x-2">
                <button
                  onClick={() => setDateRange('week')}
                  className={`px-3 py-2 text-sm font-medium rounded-md ${dateRange === 'week' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'}`}
                >
                  Неделя
                </button>
                <button
                  onClick={() => setDateRange('month')}
                  className={`px-3 py-2 text-sm font-medium rounded-md ${dateRange === 'month' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'}`}
                >
                  Месяц
                </button>
                <button
                  onClick={() => setDateRange('year')}
                  className={`px-3 py-2 text-sm font-medium rounded-md ${dateRange === 'year' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'}`}
                >
                  Год
                </button>
              </div>
            </div>
          </div>
          
          {/* Общая статистика */}
          <div className="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6 mb-5">
            <h2 className="text-lg font-medium text-gray-900 mb-4">Общая статистика</h2>
            <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
              <div className="bg-gray-50 overflow-hidden shadow rounded-lg">
                <div className="px-4 py-5 sm:p-6">
                  <dt className="text-sm font-medium text-gray-500 truncate">Всего кликов</dt>
                  <dd className="mt-1 text-3xl font-semibold text-gray-900">
                    {metrics.reduce((sum, metric) => sum + metric.clicks, 0)}
                  </dd>
                </div>
              </div>
              <div className="bg-gray-50 overflow-hidden shadow rounded-lg">
                <div className="px-4 py-5 sm:p-6">
                  <dt className="text-sm font-medium text-gray-500 truncate">Активных кнопок</dt>
                  <dd className="mt-1 text-3xl font-semibold text-gray-900">
                    {buttons.filter(button => button.isActive).length}
                  </dd>
                </div>
              </div>
              <div className="bg-gray-50 overflow-hidden shadow rounded-lg">
                <div className="px-4 py-5 sm:p-6">
                  <dt className="text-sm font-medium text-gray-500 truncate">Последний клик</dt>
                  <dd className="mt-1 text-3xl font-semibold text-gray-900">
                    {metrics.some(m => m.lastClickAt) 
                      ? new Date(Math.max(...metrics.filter(m => m.lastClickAt).map(m => new Date(m.lastClickAt)))).toLocaleDateString() 
                      : 'Нет данных'}
                  </dd>
                </div>
              </div>
            </div>
          </div>
          
          {/* Графики */}
          <div className="grid grid-cols-1 gap-5 lg:grid-cols-2">
            {/* График по кнопкам */}
            <div className="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6">
              <h2 className="text-lg font-medium text-gray-900 mb-4">Клики по кнопкам</h2>
              {buttons.length > 0 ? (
                <Bar data={prepareButtonsChartData()} options={options} />
              ) : (
                <p className="text-gray-500 text-center py-10">Нет данных для отображения</p>
              )}
            </div>
            
            {/* График по дням */}
            <div className="bg-white shadow px-4 py-5 sm:rounded-lg sm:p-6">
              <h2 className="text-lg font-medium text-gray-900 mb-4">Динамика кликов</h2>
              {dailyClicks.length > 0 ? (
                <Line data={prepareDailyChartData()} options={options} />
              ) : (
                <p className="text-gray-500 text-center py-10">Нет данных для отображения</p>
              )}
            </div>
          </div>
          
          {/* Таблица с детальной статистикой по кнопкам */}
          <div className="bg-white shadow sm:rounded-lg mt-5">
            <div className="px-4 py-5 sm:p-6">
              <h2 className="text-lg font-medium text-gray-900 mb-4">Детальная статистика по кнопкам</h2>
              <div className="overflow-x-auto">
                <table className="min-w-full divide-y divide-gray-300">
                  <thead className="bg-gray-50">
                    <tr>
                      <th scope="col" className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Название</th>
                      <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">URL</th>
                      <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Клики</th>
                      <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Последний клик</th>
                      <th scope="col" className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Статус</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-gray-200 bg-white">
                    {buttons.length > 0 ? (
                      buttons.map(button => {
                        const buttonMetrics = metrics.find(m => m.linkButtonId === button.id) || { clicks: 0 };
                        return (
                          <tr key={button.id}>
                            <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">{button.title}</td>
                            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                              <a href={button.url} target="_blank" rel="noopener noreferrer" className="text-indigo-600 hover:text-indigo-900">
                                {button.url.length > 30 ? button.url.substring(0, 30) + '...' : button.url}
                              </a>
                            </td>
                            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{buttonMetrics.clicks}</td>
                            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                              {buttonMetrics.lastClickAt ? new Date(buttonMetrics.lastClickAt).toLocaleString() : 'Нет данных'}
                            </td>
                            <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                              <span className={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${button.isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}`}>
                                {button.isActive ? 'Активна' : 'Неактивна'}
                              </span>
                            </td>
                          </tr>
                        );
                      })
                    ) : (
                      <tr>
                        <td colSpan="5" className="px-3 py-4 text-sm text-gray-500 text-center">Нет данных для отображения</td>
                      </tr>
                    )}
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}