<script lang="ts">
  import { onMount } from "svelte";
  import api from "../services/users";
  import type {ExpensesResponse, Expense} from "../services/expenses";
  import ExpenseCard from "./ExpenseCard.svelte";
  import { ProgressCircular, MaterialApp } from 'svelte-materialify';

  let expenses: ExpensesResponse;
  onMount(async () => {
    expenses = await api<ExpensesResponse>("http://localhost:1323/expenses?limit=10");
  })

  const deleteExpene = (expense: Expense) => {
    const index = expenses.expenses.indexOf(expense);
    expenses.expenses = [...expenses.expenses.slice(0, index), ...expenses.expenses.slice(index + 1)];
  }
</script>

<MaterialApp>
  {#if expenses}
    {#each expenses.expenses as expense }
      <div class="d-flex flex-column">
        <ExpenseCard expense={expense} on:remove={_e => deleteExpene(expense)}/>
      </div>
    {/each}
  {:else}
    <ProgressCircular
      size={50} 
      indeterminate color="primary" 
    />
  {/if}
</MaterialApp>
