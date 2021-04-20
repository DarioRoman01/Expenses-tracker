<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { goto } from "@roxi/routify";
  import api from "../services/users";
  import type { Category } from "../services/users";
  import { onMount } from "svelte";

  let categorys: Category[];
  let amount: number;
  let description: string;
  let category: string;
  let error: Error;

  const create = async () => {
    try {
      const res = await fetch("http://localhost:1323/expenses", {
        method: "POST",
        body: JSON.stringify({
          amount: amount,
          description: description,
          category: category,
        }),
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (res.ok) {
        $goto("./home");
      } else {
        error = new Error("something wrong happend");
      }
    } catch (err) {
      console.log(err);
      error = new Error("somthing wrong happend :(");
    }
  };

  onMount(async () => {
    categorys = await api<Category[]>("http://localhost:1323/categorys");
  });
</script>

<Navbar isLoggedIn={true} />
<main>
  <form action="#" on:submit|preventDefault={create}>
    <h1>Add Expense</h1>
    <div class="formcontainer">
      <hr />
      <div class="container">
        <label for="uname"><strong>Amount</strong></label>
        <input type="text" placeholder="amount" bind:value={amount} />
        <label for="uname"><strong>Description</strong></label>
        <input type="text" placeholder="description" bind:value={description} />
        <label for="psw"><strong>Category</strong></label>
        {#if categorys}
          <select name="dropdown" bind:value={category}>
            {#each categorys as category}
              <option value={category.name} selected>{category.name}</option>
            {/each}
          </select>
        {:else}
          <p>loading...</p>
        {/if}
        {#if error}
          <div class="error-container">
            <p style="color: #003049;">{error.message}</p>
          </div>
        {/if}
      </div>
      <button type="submit">Add to Expenses</button>
    </div>
  </form>
</main>
