<script lang="ts">
  import type {Expense} from "../services/users"
  export let expense: Expense;
</script>

<div class="expense">
  <h3>{expense.amount}</h3>
  <p>{expense.description}</p>
  <p>{expense.category}</p>
  <legend>{expense.createdAt}</legend>
</div>

<style>
  .expense {
    width: 600px;
    height: 200px;
    border: #002f4928 1px groove;
    margin-top: 10px;
  }
</style>