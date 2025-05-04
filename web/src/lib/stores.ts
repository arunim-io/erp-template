import { createStore } from "zustand";

interface AuthStore {
  isAuthenticated: boolean;
}

export const authStore = createStore<AuthStore>()(() => ({
  isAuthenticated: false,
}));
