<script lang="ts">
  import type {Expense} from "../services/users"
  export let expense: Expense;
</script>

<div class="expense">
  <div class="expense-title-container">
    Amount:
    <h3 class="expense-title">
      {expense.amount}
    </h3>
  </div>
  <div class="expense-desc">
    Description: 
    <p>{expense.description}</p>
  </div>
  <div class="expense-catgory">
    Category: 
    <p>{expense.category}</p>
  </div>
  <div class="created">
    <legend>{expense.createdAt}</legend>
  </div>
</div>

<style>
  .expense {
    width: 600px;
    height: 200px;
    border: #002f4928 1px groove;
    margin-top: 10px;
  }
  .expense-title {
    border: #c1121f 1px solid;
    background-color: #c1121f;
    color: #ffffff;
    max-width: 100px;
    text-align: center;
    margin-left: 10px;
    height: 20px;
  }

  .expense-title-container {
    display: flex;
    align-items: center;
    font-weight: bold;
    margin-left: 10px;
  }

  .expense-desc {
    display: flex;
    align-items: center;
    font-weight: bold;
    margin-left: 10px;
  }

  .expense-desc p {
    margin-left: 10px;
    font-weight: initial;
  }

  .expense-catgory {
    display: flex;
    align-items: center;
    font-weight: bold;
    margin-left: 10px;
  }

  .expense-catgory p {
    margin-left: 10px;
    font-weight: initial;
  } 

  .created {
    margin-left: 10px;
  }
</style>