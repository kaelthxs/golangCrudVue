<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="form.username" required />
      </div>

      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="form.password" required />
      </div>

      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { ref } from 'vue';

export default {
  name: 'LoginPage',
  setup() {
    const form = ref({
      username: '',
      password: '',
    });

    const handleSubmit = async () => {
      try {
        const response = await axios.post('http://localhost:8080/auth/sign-in', form.value);
        console.log('Login successful', response.data);
        // Сохранение токена и редирект
      } catch (error) {
        console.error('Login failed:', error);
      }
    };

    return { form, handleSubmit };
  }
};
</script>

<style scoped>
/* Стили для формы логина */
form {
  max-width: 400px;
  margin: 0 auto;
}
div {
  margin-bottom: 10px;
}
label {
  display: block;
  margin-bottom: 5px;
}
input {
  width: 100%;
  padding: 8px;
  font-size: 14px;
}
button {
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}
</style>
