import { Roles } from './constants/roles';
import { writable, type Writable } from 'svelte/store';

interface AuthState {
  userRole: Roles;
  isLoggedIn: boolean;
}

const initialAuthState: AuthState = {
  userRole: Roles.Guest,
  isLoggedIn: false,
};

export const authStore: Writable<AuthState> = writable(initialAuthState);

export function login(userRole: Roles): void {
  authStore.set({ userRole, isLoggedIn: true });
}

export function logout(): void {
  authStore.set({ userRole: Roles.Guest, isLoggedIn: false });
}