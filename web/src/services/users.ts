export default async function api<T>(url: string): Promise<T> {
  const response = await fetch(url, {
    credentials: 'include',
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  })
  if (!response.ok) {
    return;
  }
  return await response.json();
}

export const login =  async ({usernameOrEmail, password}): Promise<User> => {
  const res = await fetch("http://localhost:1323/login", {
    method: "post",
    body: JSON.stringify({
      usernameOrEmail: usernameOrEmail,
      password: password,
    }),
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message);
  }

  return await res.json();
}

export const register = async ({username, email, password}): Promise<User> => {
  const res = await fetch("http://localhost:1323/signup", {
    method: "POST",
    body: JSON.stringify({
      username: username,
      email: email,
      password: password,
    }),
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
 
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message)
  }

  return await res.json();
}

// handle users forgot password request
export const forgotPassword = async (email: string) => {
  const res = await fetch("http://localhost:1323/forgot-password", {
    method: "post",
    body: JSON.stringify({
      email: email
    }),
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      },
    credentials: "include",
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message)
  }

  return await res.json();
}

// handles users change password request
export const changePassword = async (
  token: string, 
  newPassword: string
  ): Promise<User> => {
  const res = await fetch("http://localhost:1323/change-password", {
    method: "POST",
    body: JSON.stringify({
      token: token,
      newPassword: newPassword
    }),
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        },
    credentials: "include",
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.message)
  }

  return await res.json();
}

export interface User {
  id: number;
  createdAt: string;
  updatedAt: string;
  username: string;
  email: string;
}

export interface Expense {
  id: number;
  createdAt: string;
  updatedAt: string;
  description: string;
  amount: number;
  category: string;
  user_id: number;
}

export interface ExpensesResponse {
  expenses: Expense[];
  hasMore: boolean;
}

export interface Category {
  id: number;
  name: string;
}

