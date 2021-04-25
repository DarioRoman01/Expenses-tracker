<script lang="ts">
  import { onMount } from "svelte";
  import api from "../services/users";
  import type { ExpensesResponse, Expense } from "../services/expenses";
  import Modal from "./Modal.svelte";
  import ExpenseCard from "./ExpenseCard.svelte";
  import {
    ProgressCircular,
    MaterialApp,
    Overlay,
    Icon,
    Button,
  } from "svelte-materialify";
  import { mdiPlusBox } from "@mdi/js";

  let active = false;
  let showedExpense: Expense;
  let expenses: ExpensesResponse;

  onMount(async () => {
    expenses = await api<ExpensesResponse>(
      "http://localhost:1323/expenses?limit=10"
    );
  });

  const deleteExpene = (expense: Expense) => {
    const index = expenses.expenses.indexOf(expense);
    expenses.expenses = [
      ...expenses.expenses.slice(0, index),
      ...expenses.expenses.slice(index + 1),
    ];
  };

  const handleUpdate = () => {
    active = false;
    const res = api<ExpensesResponse>(
      "http://localhost:1323/expenses?limit=10"
    );
    res.then((newExpenses) => (expenses = newExpenses));
  };

  const loadMore = (lastExpense: Expense) => {
    const moreExpenses = api<ExpensesResponse>(
      `http://localhost:1323/expenses?limit=10&cursor=${lastExpense.createdAt}`
    );

    moreExpenses.then((e) => {
      expenses.hasMore = e.hasMore;
      e.expenses.forEach(ex => expenses.expenses.push(ex));
    });
  };
</script>

<MaterialApp>
  {#if expenses}
    {#each expenses.expenses as expense}
      <div class="d-flex flex-column">
        <ExpenseCard
          {expense}
          on:remove={(_e) => deleteExpene(expense)}
          on:update={(_e) => ((active = true), (showedExpense = expense))}
        />
        <Overlay {active}>
          <Modal
            expense={showedExpense}
            on:updates={(_e) => handleUpdate()}
            on:close={(_e) => (active = false)}
          />
        </Overlay>
      </div>
    {/each}
    {#if expenses.hasMore == true}
      <div class="d-flex justify-center mt-4 mb-3">
        <Button
          class="primary-color"
          on:click={() =>
            loadMore(expenses.expenses[expenses.expenses.length - 1])}
        >
          <Icon path={mdiPlusBox} />
        </Button>
      </div>
    {/if}
  {:else}
    <ProgressCircular size={50} indeterminate color="primary" />
  {/if}
</MaterialApp>
