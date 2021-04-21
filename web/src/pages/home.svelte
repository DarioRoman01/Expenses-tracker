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
  });

</script>

<Navbar isLoggedIn={isLoggedIn} />
{#if user}
  <div class="mt-4">
    <h2 class="text-h2">Welcome {user.username}</h2>
    <div class="mt-4">
      <Expenses />
    </div>
  </div>  
{:else}
  <p>loading...</p>
{/if}