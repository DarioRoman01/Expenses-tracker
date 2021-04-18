<script lang="ts">
  import { onMount } from "svelte";
  import api from "../services/users"
  import type {ExpensesResponse} from "../services/users";
  import ExpenseCard from "./ExpenseCard.svelte"

  let expenses: ExpensesResponse;
  onMount(async () => {
    expenses = await api<ExpensesResponse>("http://localhost:1323/expenses?limit=10");
  })

</script>

<div class="expenses">
{#if expenses}
  {#each expenses.expenses as expense }
    <div class="expense-container">
      <ExpenseCard expense={expense}/>
    </div>
  {/each}
{:else}
  <p>loading...</p>
{/if}
</div>

<style>
  .expenses {
    display: grid;
    grid-template-rows: auto;
    grid-column: auto;
  }

  .expense-container {
    display: flex;
    justify-content: center;
    
  }
</style>
