<template>
  <div>
    <h1>Clients</h1>
    <button @click="fetchClients">Refresh List</button>
    <ul>
      <li v-for="client in clients" :key="client.id">
        {{ client.username }} ({{ client.email }})
        <button @click="deleteClient(client.id)">Delete</button>
        <button @click="$emit('edit-client', client)">Edit</button>
      </li>
    </ul>
  </div>
</template>

<script>
import api from "../utils/api";

export default {
  data() {
    return {
      clients: [],
    };
  },
  methods: {
    async fetchClients() {
      try {
        const response = await api.get("/api/client");
        this.clients = response.data;
      } catch (error) {
        console.error("Failed to fetch clients:", error);
      }
    },
    async deleteClient(id) {
      try {
        await api.delete(`/api/client/${id}`);
        this.fetchClients(); // Refresh list after deletion
      } catch (error) {
        console.error("Failed to delete client:", error);
      }
    },
  },
  mounted() {
    this.fetchClients();
  },
};
</script>
