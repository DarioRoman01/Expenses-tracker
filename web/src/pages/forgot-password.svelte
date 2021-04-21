<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { forgotPassword } from "../services/users";
  import {TextField, Button, MaterialApp, Alert } from 'svelte-materialify';

  let email: string;
  let error: Error;
  let send: boolean;

  const handleClick = () => {
    const result = forgotPassword(email);
    result.then(() => send = true).catch((err: Error) => error = err)
  }
</script>

<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      {#if !send}   
      <h3 class="text-h3 mb-6">Change password</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={email}>
          email
        </TextField>
      </div>
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

      {:else}
        <h2 class="text-h2">We send you an email</h2>
        <h4 class="text-h4">please check your email</h4>
      {/if}
    </div>
  </div>
</MaterialApp>

