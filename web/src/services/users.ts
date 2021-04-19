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
