import { createApp } from 'vue';
import App from './App.vue';
import router from './router';  // Импортируем маршрутизатор

// Создаем приложение, подключаем маршрутизатор и монтируем его
createApp(App)
    .use(router)  // Подключаем маршрутизатор
    .mount('#app');
