<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import {TextField, Button, MaterialApp, Alert } from 'svelte-materialify';
  import { redirect, url } from "@roxi/routify";
  import { register } from "../services/users";

  let username: string;
  let email: string;
  let password: string;
  let error: Error;

  const handleRegister = () => {
    const user = register({username, email, password});
    user.then(() => $redirect("./home")).catch((err: Error) => error = err)
  }

</script>

<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      <h3 class="text-h3 mb-6">Login</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={username}>
          username
        </TextField>
      </div>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={email}>
          email
        </TextField>
      </div>
      <div style="width: 700px;">
        <TextField type="password" bind:value={password}>
          password
        </TextField>
      </div>
      <div style="align-self: center;" class="mt-4">
        <Button size="large" class="primary-color" on:click={handleRegister}>
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
