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

export const addExpense = async (
  description: string,
  amount: number,
  category: string
): Promise<ExpensesResponse> => {
  const res = await fetch("http://localhost:1323/expenses", {
    method: "POST",
    body: JSON.stringify({
      description: description,
      amount: amount,
      category: category,
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
};

export const deleteExpense = async (id: number) => {
  const res = await fetch(`http://localhost:1323/expenses/${id}`, {
    method: "DELETE",
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

export const updateExpense = async (
  id: number,
  description: string,
  amount: number,
  category: string,
): Promise<Expense> => {
  const res = await fetch(`http://localhost:1323/expenses/${id}`, {
    method: "PATCH",
    body: JSON.stringify({
      description: description,
      amount: amount,
      category: category,
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
