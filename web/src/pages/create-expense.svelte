<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { redirect } from "@roxi/routify";
  import api from "../services/users";
  import type { Category } from "../services/expenses";
  import { addExpense } from "../services/expenses";
  import { onMount } from "svelte";
  import {
    TextField, 
    Button, 
    MaterialApp, 
    Alert,
    Menu,
    List, 
    ListItem 
  } from 'svelte-materialify';

  let categorys: Category[];
  let amount: string;
  let description: string;
  let categoryo: string;
  let error: Error;

  const handleClick = () => {
    const expense = addExpense(description, parseInt(amount), categoryo);
    expense.then(() => $redirect("./home")).catch((err: Error) => error = err)
  }

  onMount(async () => {
    categorys = await api<Category[]>("http://localhost:1323/categorys");
  });
</script>


<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      <h3 class="text-h3 mb-6">Login</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={amount}>
          username
        </TextField>
      </div>
      <div style="width: 700px;">
        <TextField bind:value={description}>
          description
        </TextField>
      </div>
      <Menu>
        <div slot="activator">
          <Button>categorys</Button>
        </div>
        <List>
          {#if categorys}
            {#each categorys as category }
              <ListItem bind:value={categoryo}>
                {category.name}
              </ListItem>
            {/each}
          {:else}
            <p>loading</p>
          {/if}
        </List>
      </Menu>
      <div style="align-self: center;" class="mt-4">
        <Button size="large" class="primary-color" on:click={handleClick}>
          Submit
        </Button>
      </div>
      {#if error}
        <Alert class="error-color">
          {error.message}
        </Alert>
      {/if}
    </div>
  </div>
</MaterialApp>