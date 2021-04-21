<script lang="ts">
  import { goto, url } from '@roxi/routify';
  import { AppBar, Button, Icon, Menu, ListItem, MaterialApp } from 'svelte-materialify';
  export let isLoggedIn: boolean;

  async function logout() {
    const res = await fetch("http://localhost:1323/logout", {
      method: "post",
      credentials: "include"
    });
    if (res.ok) {
      $goto("./index");
    }
  }
</script>

<MaterialApp>
  <AppBar>
    <span slot="title">
      <a href={isLoggedIn ? $url("./home") : $url("./index")} 
        class="text-decoration-none"
      >
        Expenses Tracker
      </a>
    </span>
    <div style="flex-grow:1" />
    {#if isLoggedIn}
      <Button class="primary-color" on:click={$goto("./create-expense")}>
        Add
      </Button>
      <Button class="ml-3 primary-color" on:click={logout}>
        logout
      </Button>
    {:else}
      <Button class="primary-color" on:click={$goto("./login")}>
        Login
      </Button>
      <Button class="ml-4 primary-color" on:click={$goto("./register")}>
        Register
      </Button>
    {/if}
  </AppBar>
</MaterialApp>

