<script lang="ts">
  import { onMount } from "svelte";
  import Navbar from "../components/Navbar2.svelte";
  import api from "../services/users";
  import type {User} from "../services/users"
  import Expenses from "../components/Expenses.svelte"
  
  let user: User;
  let isLoggedIn: boolean;

  onMount(async () => {
    isLoggedIn = false
    user = await api<User>("http://localhost:1323/me");
    if (user) {
      isLoggedIn = true;
    } 
  })

</script>

<Navbar isLoggedIn={isLoggedIn} />
{#if user}
  <div class="home-container">
    <h1>Welcome {user.username}</h1>
    <Expenses />
  </div>  
{:else}
  <p>loading...</p>
{/if}

<style>
  .home-container h1 {
    text-align: left;
    padding-left: 10px;
    color: #003049;
  }
</style>