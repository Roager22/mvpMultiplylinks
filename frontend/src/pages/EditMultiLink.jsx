import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import LinkButtonList from '../components/LinkButtonList';
import LinkButtonForm from '../components/LinkButtonForm';

export default function EditMultiLink() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [multiLink, setMultiLink] = useState(null);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [slug, setSlug] = useState('');
  const [isActive, setIsActive] = useState(true);
  const [buttons, setButtons] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [showButtonForm, setShowButtonForm] = useState(false);
  const [editingButton, setEditingButton] = useState(null);

  useEffect(() => {
    // Проверка аутентификации
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
      return;
    }

    // Установка токена для запросов
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    // Загрузка данных мультиссылки
    fetchMultiLink();
  }, [id, navigate]);

  const fetchMultiLink = async () => {
    try {
      setLoading(true);
      const response = await axios.get(`/api/multilinks/${id}`);
      const multiLinkData = response.data.multilink;
      
      setMultiLink(multiLinkData);
      setTitle(multiLinkData.title);
      setDescription(multiLinkData.description || '');
      setSlug(multiLinkData.slug);
      setIsActive(multiLinkData.isActive);

      // Загрузка кнопок для мультиссылки
      const buttonsResponse = await axios.get(`/api/multilinks/${id}/buttons`);
      setButtons(buttonsResponse.data.buttons || []);
      
      setError('');
    } catch (err) {
      setError(err.response?.data?.error || 'Произошла ошибка при загрузке данных. Пожалуйста, попробуйте снова.');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await axios.put(`/api/multilinks/${id}`, {
        title,
        description,
        slug,
        isActive
      });
      
      // Обновляем данные после успешного сохранения
      fetchMultiLink();
    } catch (err) {
      setError(err.response?.data?.error || 'Произошла ошибка при обновлении мультиссылки. Пожалуйста, попробуйте снова.');
    } finally {
      setLoading(false);
    }
  };

  const handleAddButton = () => {
    setEditingButton(null);
    setShowButtonForm(true);
  };

  const handleEditButton = (button) => {
    setEditingButton(button);
    setShowButtonForm(true);
  };

  const handleButtonFormClose = () => {
    setShowButtonForm(false);
    setEditingButton(null);
  };

  const handleButtonSave = async (buttonData) => {
    try {
      if (editingButton) {
        // Обновление существующей кнопки
        await axios.put(`/api/multilinks/${id}/buttons/${editingButton.id}`, buttonData);
      } else {
        // Создание новой кнопки
        await axios.post(`/api/multilinks/${id}/buttons`, buttonData);
      }
      
      // Обновляем список кнопок
      const response = await axios.get(`/api/multilinks/${id}/buttons`);
      setButtons(response.data.buttons || []);
      
      // Закрываем форму
      setShowButtonForm(false);
      setEditingButton(null);
    } catch (err) {
      setError(err.response?.data?.error || 'Произошла ошибка при сохранении кнопки. Пожалуйста, попробуйте снова.');
    }
  };

  const handleDeleteButton = async (buttonId) => {
    if (!confirm('Вы уверены, что хотите удалить эту кнопку?')) {
      return;
    }
    
    try {
      await axios.delete(`/api/multilinks/${id}/buttons/${buttonId}`);
      
      // Обновляем список кнопок
      setButtons(buttons.filter(button => button.id !== buttonId));
    } catch (err) {
      setError(err.response?.data?.error || 'Произошла ошибка при удалении кнопки. Пожалуйста, попробуйте снова.');
    }
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
          <h1 className="text-3xl font-bold leading-tight tracking-tight text-gray-900">Редактирование мультиссылки</h1>
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
          
          <div className="px-4 py-5 sm:p-6 bg-white shadow sm:rounded-lg mt-5">
            <form onSubmit={handleSubmit}>
              <div className="space-y-6">
                <div>
                  <label htmlFor="title" className="block text-sm font-medium text-gray-700">
                    Название
                  </label>
                  <div className="mt-1">
                    <input
                      id="title"
                      name="title"
                      type="text"
                      required
                      value={title}
                      onChange={(e) => setTitle(e.target.value)}
                      className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                  </div>
                </div>

                <div>
                  <label htmlFor="description" className="block text-sm font-medium text-gray-700">
                    Описание
                  </label>
                  <div className="mt-1">
                    <textarea
                      id="description"
                      name="description"
                      rows={3}
                      value={description}
                      onChange={(e) => setDescription(e.target.value)}
                      className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                  </div>
                </div>

                <div>
                  <label htmlFor="slug" className="block text-sm font-medium text-gray-700">
                    Slug (уникальный идентификатор для URL)
                  </label>
                  <div className="mt-1">
                    <input
                      id="slug"
                      name="slug"
                      type="text"
                      value={slug}
                      onChange={(e) => setSlug(e.target.value)}
                      className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                  </div>
                </div>

                <div className="relative flex items-start">
                  <div className="flex h-5 items-center">
                    <input
                      id="isActive"
                      name="isActive"
                      type="checkbox"
                      checked={isActive}
                      onChange={(e) => setIsActive(e.target.checked)}
                      className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                  </div>
                  <div className="ml-3 text-sm">
                    <label htmlFor="isActive" className="font-medium text-gray-700">
                      Активна
                    </label>
                    <p className="text-gray-500">Мультиссылка будет доступна для просмотра</p>
                  </div>
                </div>

                <div className="flex justify-end">
                  <button
                    type="button"
                    onClick={() => navigate('/dashboard')}
                    className="rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 mr-3"
                  >
                    Отмена
                  </button>
                  <button
                    type="submit"
                    disabled={loading}
                    className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    {loading ? 'Сохранение...' : 'Сохранить'}
                  </button>
                </div>
              </div>
            </form>
          </div>
          
          {/* Секция с кнопками-ссылками */}
          <div className="px-4 py-5 sm:p-6 bg-white shadow sm:rounded-lg mt-5">
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-lg font-medium text-gray-900">Кнопки-ссылки</h2>
              <button
                type="button"
                onClick={handleAddButton}
                className="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
              >
                Добавить кнопку
              </button>
            </div>
            
            {/* Список кнопок */}
            <LinkButtonList 
              buttons={buttons} 
              onEdit={handleEditButton} 
              onDelete={handleDeleteButton} 
            />
            
            {/* Форма добавления/редактирования кнопки */}
            {showButtonForm && (
              <LinkButtonForm 
                button={editingButton} 
                onSave={handleButtonSave} 
                onCancel={handleButtonFormClose} 
              />
            )}
          </div>
        </div>
      </main>
    </div>
  );
}