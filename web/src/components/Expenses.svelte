<script lang="ts">
  import { onMount } from "svelte";
  import api from "../services/users";
  import type {ExpensesResponse} from "../services/expenses";
  import ExpenseCard from "./ExpenseCard.svelte";
  import { ProgressCircular, MaterialApp } from 'svelte-materialify';

  let expenses: ExpensesResponse;
  onMount(async () => {
    expenses = await api<ExpensesResponse>("http://localhost:1323/expenses?limit=10");
  })

</script>

<MaterialApp>
  {#if expenses}
    {#each expenses.expenses as expense }
      <div class="d-flex flex-column">
        <ExpenseCard expense={expense}/>
      </div>
    {/each}
  {:else}
    <ProgressCircular
      size={50} 
      indeterminate color="primary" 
    />
  {/if}
</MaterialApp>
