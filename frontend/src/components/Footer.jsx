import { Link } from 'react-router-dom';

export default function Footer() {
  return (
    <footer className="bg-gray-800 text-white py-6">
      <div className="container mx-auto px-4">
        <div className="flex flex-col md:flex-row justify-between items-center">
          <div className="mb-4 md:mb-0">
            <Link to="/" className="text-xl font-bold">MultyLink</Link>
            <p className="text-gray-400 text-sm mt-1">Все ваши ссылки в одном месте</p>
          </div>
          
          <div className="flex flex-col md:flex-row space-y-2 md:space-y-0 md:space-x-6">
            <Link to="/" className="text-gray-300 hover:text-white text-sm">
              Главная
            </Link>
            <Link to="/about" className="text-gray-300 hover:text-white text-sm">
              О сервисе
            </Link>
            <Link to="/privacy" className="text-gray-300 hover:text-white text-sm">
              Политика конфиденциальности
            </Link>
            <Link to="/terms" className="text-gray-300 hover:text-white text-sm">
              Условия использования
            </Link>
          </div>
        </div>
        
        <div className="mt-6 pt-6 border-t border-gray-700 text-center text-gray-400 text-sm">
          &copy; {new Date().getFullYear()} MultyLink. Все права защищены.
        </div>
      </div>
    </footer>
  );
}