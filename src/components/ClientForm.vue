<template>
  <div>
    <h2>{{ isEditMode ? "Edit Client" : "Add Client" }}</h2>
    <form @submit.prevent="submitForm">
      <input v-model="client.username" placeholder="Username" required />
      <input v-model="client.email" placeholder="Email" type="email" required />
      <input v-model="client.phone_number" placeholder="Phone Number" required />
      <input v-model="client.role" placeholder="Role" required />
      <input v-model="client.password" placeholder="Password" type="password" required />
      <button type="submit">{{ isEditMode ? "Update" : "Add" }}</button>
      <button type="button" @click="clearForm">Cancel</button>
    </form>
  </div>
</template>

<script>
import api from "../utils/api";

export default {
  props: {
    clientToEdit: {
      type: Object,
      default: null,
    },
  },
  data() {
    return {
      client: {
        username: "",
        email: "",
        phone_number: "",
        role: "",
        password: "",
      },
    };
  },
  computed: {
    isEditMode() {
      return !!this.clientToEdit;
    },
  },
  watch: {
    clientToEdit: {
      immediate: true,
      handler(newVal) {
        if (newVal) {
          this.client = { ...newVal };
        }
      },
    },
  },
  methods: {
    async submitForm() {
      try {
        if (this.isEditMode) {
          await api.put(`/api/client/${this.client.id}`, this.client);
        } else {
          await api.post("/api/client", this.client);
        }
        this.$emit("form-submitted");
        this.clearForm();
      } catch (error) {
        console.error("Failed to submit form:", error);
      }
    },
    clearForm() {
      this.client = {
        username: "",
        email: "",
        phone_number: "",
        role: "",
        password: "",
      };
      this.$emit("clear-form");
    },
  },
};
</script>
