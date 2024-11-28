import { createRouter, createWebHistory } from 'vue-router';
import LoginPage from '../components/Login.vue';
import RegisterPage from '../components/Register.vue';
import DashboardPage from '../components/Dashboard.vue';  // Пример компонента для после логина

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: LoginPage,
    },
    {
        path: '/register',
        name: 'Register',
        component: RegisterPage,
    },
    {
        path: '/dashboard',  // Пример маршрута для главной страницы
        name: 'Dashboard',
        component: DashboardPage,
    },
    // Добавьте другие маршруты по необходимости
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
