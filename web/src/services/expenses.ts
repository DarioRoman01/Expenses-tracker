// expense data that is return from the request
export interface Expense {
  id: number;
  createdAt: string;
  updatedAt: string;
  description: string;
  amount: number;
  category: string;
  user_id: number;
}

// expense response paginated
export interface ExpensesResponse {
  expenses: Expense[];
  hasMore: boolean;
}

// category data that is return from the server
export interface Category {
  id: number;
  name: string;
}

// handle send create expense request
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

// send delete expense request
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
};

// handle expense updates requests
export const updateExpense = async (
  id: number,
  description: string,
  amount: number,
  category: string
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
};

// send get expenses by category request
export const getByCategory = async (
  category: string
): Promise<ExpensesResponse> => {
  const res = await fetch(`http://localhost:1323/expenses/${category}`, {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!res.ok) {
    const err: string = await res.json();
    throw new Error(err);
  }

  return await res.json();
};

// send list expenses request with cursor
export const getByDate = async (date: string): Promise<ExpensesResponse> => {
  const res = await fetch(
    `http://localhost:1323/expenses?limit=10&cursor=${date}`,
    {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      credentials: "include",
    }
  );

  if (!res.ok) {
    const err: string = await res.json();
    throw new Error(err);
  }

  return await res.json();
};
