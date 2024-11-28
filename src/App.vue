<template>
  <div id="app">
    <h1>Welcome to the App</h1>
    <button @click="goToLogin">Go to Login</button>
    <button @click="goToRegister">Go to Register</button>

    <!-- Вы можете добавить сюда другие компоненты или контейнеры -->
  </div>
</template>

<script>
import { defineComponent } from 'vue';  // Импортируем для использования Composition API
import { useRouter } from 'vue-router'; // Подключаем useRouter

export default defineComponent({
  name: 'App',
  setup() {
    const router = useRouter(); // Получаем доступ к роутеру

    // Функции для перехода по маршруту
    const goToLogin = () => {
      router.push("/login"); // Перенаправление на страницу логина
    };

    const goToRegister = () => {
      router.push("/register"); // Перенаправление на страницу регистрации
    };

    return {
      goToLogin,
      goToRegister
    };
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
}

button {
  margin: 10px;
  padding: 10px;
  font-size: 16px;
  cursor: pointer;
}
</style>
