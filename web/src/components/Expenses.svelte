<script lang="ts">
  import { onMount } from "svelte";
  import api from "../services/users";
  import type {ExpensesResponse, Expense} from "../services/expenses";
  import Modal from "./Modal.svelte"
  import ExpenseCard from "./ExpenseCard.svelte";
  import { ProgressCircular, MaterialApp, Overlay } from 'svelte-materialify';

  let active = false;
  let expenses: ExpensesResponse;
  onMount(async () => {
    expenses = await api<ExpensesResponse>("http://localhost:1323/expenses?limit=10");
  })

  const deleteExpene = (expense: Expense) => {
    const index = expenses.expenses.indexOf(expense);
    expenses.expenses = [...expenses.expenses.slice(0, index), ...expenses.expenses.slice(index + 1)];
  }

  const handleUpdate = () => {
    active = false
    const res = api<ExpensesResponse>("http://localhost:1323/expenses?limit=10");
    res.then(newExpenses => expenses = newExpenses)
  }

</script>

<MaterialApp>
  {#if expenses}
    {#each expenses.expenses as expense }
      <div class="d-flex flex-column">
        <ExpenseCard 
            expense={expense} 
            on:remove={_e => deleteExpene(expense)}
            on:update={_e => active = true}
          />
          <Overlay {active}>
            <Modal expense={expense} on:close={_e => handleUpdate()}/>
          </Overlay>
      </div>
    {/each}
  {:else}
    <ProgressCircular
      size={50} 
      indeterminate color="primary" 
    />
  {/if}
</MaterialApp>
