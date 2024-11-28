<template>
  <div>
    <h1>Registration</h1>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="form.username" required />
      </div>

      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="form.email" required />
      </div>

      <div>
        <label for="phone_number">Phone Number:</label>
        <input type="text" id="phone_number" v-model="form.phone_number" required />
      </div>

      <div>
        <label for="role">Role:</label>
        <select id="role" v-model="form.role" required>
          <option value="ADMIN">ADMIN</option>
          <option value="USER">USER</option>
          <option value="AUTHUSER">AUTHUSER</option>
          <option value="AUTHOR">AUTHOR</option>
        </select>
      </div>

      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="form.password" required />
      </div>

      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { ref } from 'vue';

export default {
  name: 'RegisterPage',
  setup() {
    const form = ref({
      username: '',
      email: '',
      phone_number: '',
      role: 'USER', // Значение по умолчанию
      password: '',
    });

    const handleSubmit = async () => {
      try {
        const response = await axios.post('http://localhost:8080/auth/sign-in', form.value);
        console.log('User registered', response.data);
        // Возможно, редирект или уведомление
      } catch (error) {
        console.error('Registration failed:', error);
      }
    };

    return { form, handleSubmit };
  }
};
</script>

<style scoped>
/* Ваши стили для формы */
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
input, select {
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
