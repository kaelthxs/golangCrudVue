import axios from "axios";
import { getToken } from "./auth";

// Создаём экземпляр axios
const api = axios.create({
    baseURL: "http://localhost:8080", // базовый URL для всех запросов
});

// Добавляем токен в заголовки, если он есть
api.interceptors.request.use(
    (config) => {
        const token = getToken();
        if (token) {
            config.headers["Authorization"] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default api;
