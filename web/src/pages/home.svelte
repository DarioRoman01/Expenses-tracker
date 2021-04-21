<script lang="ts">
  import { onMount } from "svelte";
  import Navbar from "../components/Navbar2.svelte";
  import api from "../services/users";
  import Modal from "../components/Modal.svelte"
  import type {User} from "../services/users"
  import Expenses from "../components/Expenses.svelte"
  import { MaterialApp, ProgressCircular, Icon } from "svelte-materialify";
  import { mdiCurrencyUsd } from "@mdi/js"

  let user: User;
  let amount: number;
  let isLoggedIn: boolean;

  onMount(async () => {
    isLoggedIn = false
    user = await api<User>("http://localhost:1323/me");
    if (user) {
      isLoggedIn = true;
    };
    amount = await api("http://localhost:1323/expenses/avarage")
  });

</script>

<MaterialApp>
  <Navbar isLoggedIn={isLoggedIn} />
  {#if user}
    <div class="mt-4">
      <div class="d-flex justify-space-between align-center">
        <h2 class="text-h2">Welcome {user.username}</h2>
        {#if amount}
          <h4 class="text-h4 mr-3 d-flex align-center">
            <Icon size={30} path={mdiCurrencyUsd} />{amount}
          </h4>
        {:else}
          <ProgressCircular indeterminate color="primary" />
        {/if}
      </div>
      <div class="mt-4">
        <Expenses />
      </div>
    </div>  
  {:else}
    <p>loading...</p>
  {/if}
</MaterialApp>