<script lang="ts">
  import {
    MaterialApp,
    TextField,
    Menu,
    Button,
    List,
    ListItem,
    Icon,
  } from "svelte-materialify";
  import { createEventDispatcher, onMount } from "svelte";
  import { updateExpense } from "../services/expenses";
  import type { Expense, Category } from "../services/expenses";
  import { mdiCloseBox } from "@mdi/js";
  import api from "../services/users";

  const dispatch = createEventDispatcher();
  export let expense: Expense;
  let categorys: Category[];

  const handleUpdates = () => {
    const updatedExpense = updateExpense(
      expense.id,
      expense.description,
      expense.amount,
      expense.category
    );
    updatedExpense
      .then(() => dispatch("updates"))
      .catch((err: Error) => console.log(err.message));
  };

  const handleClose = () => {
    dispatch("close");
  };

  onMount(async () => {
    categorys = await api<Category[]>("http://localhost:1323/categorys");
  });
</script>

<MaterialApp>
  <div class="pt-4 pb-4 pl-4 pr-4">
    <div class="d-flex justify-center mt-4">
      <div class="d-flex flex-column">
        <div class="d-flex justify-space-between">
          <h3 class="text-h3 mb-6">Edit expense</h3>
          <Button class="primary-color" on:click={handleClose}>
            <Icon path={mdiCloseBox} />
          </Button>
        </div>
        <div style="width: 700px;" class="mb-4 mt-3">
          <TextField value={expense.amount.toString()} />
        </div>
        <div style="width: 700px;">
          <TextField bind:value={expense.description} />
        </div>
        <Menu>
          <div slot="activator" class="mt-4">
            <Button>{expense.category}</Button>
          </div>
          <List>
            {#if categorys}
              {#each categorys as category}
                <ListItem on:click={() => (expense.category = category.name)}>
                  {category.name}
                </ListItem>
              {/each}
            {:else}
              <p>loading</p>
            {/if}
          </List>
        </Menu>
        <div style="align-self: center;" class="mt-4">
          <Button size="large" class="primary-color" on:click={handleUpdates}>
            Submit
          </Button>
        </div>
      </div>
    </div>
  </div>
</MaterialApp>
